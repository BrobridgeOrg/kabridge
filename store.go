package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type EventStore interface {
	RegisterTopic(subject string) error
	FindStream(streamName string) (Stream, error)
	EnsureConsumer(sessionID int32, streamName string, subject []string, seq int64) error
	WriteEvent(msg *Message) error
	Fetch(sessionID int32, count int32) ([]*Message, error)
}

type Stream interface {
	Name() string
	Contains(subject string) bool
}

// MemoryStream is a simple in-memory implementation of a stream
type MemoryStream struct {
	StreamName string
	Subjects   []string
	Messages   []*Message
	LastSeq    uint64
}

func NewMemoryStream() *MemoryStream {
	return &MemoryStream{
		Subjects: make([]string, 0),
		Messages: make([]*Message, 0),
		LastSeq:  0,
	}
}

func (s *MemoryStream) Name() string {
	return s.StreamName
}

func (s *MemoryStream) Contains(subject string) bool {
	for _, s := range s.Subjects {
		if s == subject {
			return true
		}
	}

	return false
}

// MemoryEventStoreImpl is a simple in-memory implementation of EventStore
type MemoryEventStoreImpl struct {
	streams   map[string]*MemoryStream
	consumers map[int32]*Consumer
}

func NewMemoryEventStore() *MemoryEventStoreImpl {
	return &MemoryEventStoreImpl{
		streams:   make(map[string]*MemoryStream),
		consumers: make(map[int32]*Consumer),
	}
}

func (store *MemoryEventStoreImpl) RegisterTopic(topic string) error {

	// Parsing the subject
	parts := strings.Split(topic, ".")
	if len(parts) < 2 {
		return nil
	}

	subject := strings.Join(parts[1:], ".")

	stream, ok := store.streams[parts[0]]
	if !ok {
		stream = NewMemoryStream()
		stream.StreamName = parts[0]
		stream.Subjects = append(stream.Subjects, subject)
		store.streams[stream.StreamName] = stream
	}

	return nil
}

func (store *MemoryEventStoreImpl) FindStream(streamName string) (Stream, error) {

	stream, ok := store.streams[streamName]
	if !ok {
		return nil, nil
	}

	return Stream(stream), nil
}

func (store *MemoryEventStoreImpl) EnsureConsumer(sessionID int32, streamName string, subject []string, seq int64) error {

	_, ok := store.consumers[sessionID]
	if !ok {
		c := NewConsumer()
		c.Name = fmt.Sprintf("kabridge_%d", sessionID)
		c.Stream = streamName
		c.Seq = seq
		store.consumers[sessionID] = c
	}

	return nil

}

func (store *MemoryEventStoreImpl) WriteEvent(msg *Message) error {

	// Parsing the subject
	parts := strings.Split(msg.Topic, ".")
	if len(parts) < 2 {
		return nil
	}

	subject := strings.Join(parts[1:], ".")

	stream, ok := store.streams[parts[0]]
	if !ok {
		stream = NewMemoryStream()
		stream.StreamName = parts[0]
		stream.Subjects = append(stream.Subjects, subject)
		store.streams[stream.StreamName] = stream
	}

	stream.LastSeq++
	msg.Seq = stream.LastSeq
	stream.Messages = append(stream.Messages, msg)

	log.Printf("Publish Event to Stream: %s, Seq: %d, Topic: %s", stream.StreamName, stream.LastSeq, msg.Topic)

	return nil
}

func (store *MemoryEventStoreImpl) Fetch(sessionID int32, count int32) ([]*Message, error) {

	c, ok := store.consumers[sessionID]
	if !ok {
		return []*Message{}, errors.New("Consumer not found")
	}

	s, ok := store.streams[c.Stream]
	if !ok {
		return []*Message{}, nil
	}
	/*
		if int64(s.LastSeq) <= c.Seq {
			return []*Message{}, nil
		}
	*/
	//queue := s.Messages[c.Seq:]
	queue := s.Messages[0:]
	msgs := make([]*Message, 0, count)
	total := int32(0)
	for _, msg := range queue {
		msgs = append(msgs, msg)

		total++
		if total == count {
			break
		}
	}

	c.Seq += int64(total)

	return msgs, nil
}
