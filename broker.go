package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"slices"
	"strings"
	"time"

	"github.com/segmentio/kafka-go/protocol"
	"github.com/segmentio/kafka-go/protocol/apiversions"
	apiversionsAPI "github.com/segmentio/kafka-go/protocol/apiversions"
	"github.com/segmentio/kafka-go/protocol/createtopics"
	describeconfigsAPI "github.com/segmentio/kafka-go/protocol/describeconfigs"
	fetchAPI "github.com/segmentio/kafka-go/protocol/fetch"
	listgroupsAPI "github.com/segmentio/kafka-go/protocol/listgroups"
	listoffsetsAPI "github.com/segmentio/kafka-go/protocol/listoffsets"
	metadataAPI "github.com/segmentio/kafka-go/protocol/metadata"
	produceAPI "github.com/segmentio/kafka-go/protocol/produce"
)

type Broker struct {
	incoming      chan *Message
	topics        map[string]struct{}
	store         EventStore
	allowedTopics []string
}

func NewBroker() *Broker {
	return &Broker{
		incoming:      make(chan *Message),
		topics:        make(map[string]struct{}),
		store:         NewMemoryEventStore(),
		allowedTopics: make([]string, 0),
	}
}

func (b *Broker) Publish(topic string, msg *Message) {
	//	b.incoming <- msg
	b.store.WriteEvent(msg)
}

func (b *Broker) Start(address string) error {

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer ln.Close()

	go b.handleIncomingMessages()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v\n", err)
			continue
		}

		log.Println("New connection accepted")

		go b.handleConnection(conn)
	}
}

func (b *Broker) handleIncomingMessages() {
	for msg := range b.incoming {
		fmt.Println("Received message:", msg)
	}
}

func (b *Broker) handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		apiVersion, correlationID, clientID, msg, err := protocol.ReadRequest(conn)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		if err != nil {
			log.Printf("failed to decode request: %v\n", err)
			return
		}

		req := &Request{
			ApiVersion:    apiVersion,
			CorrelationID: correlationID,
			ClientID:      clientID,
			Msg:           msg,
			Err:           err,
		}

		b.handleRequest(conn, req)
	}
}

func (b *Broker) handleRequest(conn net.Conn, req *Request) {

	log.Printf("Request: %s(%d)\n", apiKeys[apiKey(req.Msg.ApiKey())], req.Msg.ApiKey())

	switch req.Msg.ApiKey() {
	case protocol.ApiVersions:
		b.handleApiVersions(conn, req)
	case protocol.Metadata:
		b.handleMetadata(conn, req)
	case protocol.ListOffsets:
		b.handleListOffsets(conn, req)
	case protocol.ListGroups:
		b.handleListGroups(conn, req)
	case protocol.DescribeConfigs:
		b.handleDescribeConfigs(conn, req)
	case protocol.CreateTopics:
		b.handleCreateTopics(conn, req)
	case protocol.Produce:
		b.handleProduce(conn, req)
	case protocol.Fetch:
		b.handleFetch(conn, req)
	default:
		log.Printf("unsupported request key: %d\n", req.Msg.ApiKey())
	}
}

func (b *Broker) handleApiVersions(conn net.Conn, req *Request) {

	res := &apiversionsAPI.Response{
		ErrorCode: 0,
		ApiKeys:   make([]apiversionsAPI.ApiKeyResponse, 0),
	}

	for apiKey, v := range APIVersions {
		res.ApiKeys = append(res.ApiKeys, apiversions.ApiKeyResponse{
			ApiKey:     int16(apiKey),
			MinVersion: v.min,
			MaxVersion: v.max,
		})
	}

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleMetadata(conn net.Conn, req *Request) {

	r := req.Msg.(*metadataAPI.Request)

	res := &metadataAPI.Response{
		ControllerID: 1,
		Brokers: []metadataAPI.ResponseBroker{
			{
				NodeID: 0,
				Host:   "0.0.0.0",
				Port:   9092,
			},
		},
		Topics: make([]metadataAPI.ResponseTopic, 0),
	}

	if r.AllowAutoTopicCreation {
		for _, topic := range r.TopicNames {
			if slices.Contains(b.allowedTopics, topic) {
				continue
			}

			b.allowedTopics = append(b.allowedTopics, topic)
		}
	}

	// Loading current topics
	for _, topic := range b.allowedTopics {
		//log.Printf("Loading topic: %s\n", topic)

		res.Topics = append(res.Topics, metadataAPI.ResponseTopic{
			ErrorCode: 0,
			Name:      topic,
			Partitions: []metadataAPI.ResponsePartition{
				{
					ErrorCode:      0,
					PartitionIndex: 0,
					LeaderID:       0,
					ReplicaNodes:   []int32{0},
					IsrNodes:       []int32{0},
				},
			},
		})
	}

	//log.Printf("Metadata request: %v\n", r)
	//log.Printf("Metadata response: %v\n", res)

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleListOffsets(conn net.Conn, req *Request) {

	r := req.Msg.(*listoffsetsAPI.Request)

	res := &listoffsetsAPI.Response{
		Topics: make([]listoffsetsAPI.ResponseTopic, 0),
	}

	for _, t := range r.Topics {

		topic := listoffsetsAPI.ResponseTopic{
			Topic: t.Topic,
			Partitions: []listoffsetsAPI.ResponsePartition{
				{
					Partition:   0,
					ErrorCode:   0,
					Timestamp:   -1,
					Offset:      0,
					LeaderEpoch: -1,
				},
			},
		}

		parts := strings.Split(t.Topic, ".")

		_, err := b.store.FindStream(parts[0])
		if err != nil {
			topic.Partitions[0].ErrorCode = 1
			res.Topics = append(res.Topics, topic)
			break
		}

		res.Topics = append(res.Topics, topic)
	}

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleListGroups(conn net.Conn, req *Request) {

	res := &listgroupsAPI.Response{
		ErrorCode: 0,
		Groups:    make([]listgroupsAPI.ResponseGroup, 0),
	}

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleDescribeConfigs(conn net.Conn, req *Request) {

	res := &describeconfigsAPI.Response{
		ThrottleTimeMs: 100,
		Resources:      make([]describeconfigsAPI.ResponseResource, 0),
	}

	res.Resources = append(res.Resources, describeconfigsAPI.ResponseResource{
		ErrorCode:     0,
		ConfigEntries: []describeconfigsAPI.ResponseConfigEntry{},
	})

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleCreateTopics(conn net.Conn, req *Request) {

	r := req.Msg.(*createtopics.Request)

	res := &createtopics.Response{
		Topics: make([]createtopics.ResponseTopic, 0),
	}

	for _, t := range r.Topics {
		log.Printf("Created Topic: %s (pattition=%d, relicas=%d)\n", t.Name, t.NumPartitions, t.ReplicationFactor)
		b.topics[t.Name] = struct{}{}
		res.Topics = append(res.Topics, createtopics.ResponseTopic{
			Name:              t.Name,
			ErrorCode:         0,
			NumPartitions:     t.NumPartitions,
			ReplicationFactor: t.ReplicationFactor,
			Configs:           []createtopics.ResponseTopicConfig{},
		})
	}

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleProduce(conn net.Conn, req *Request) {

	r := req.Msg.(*produceAPI.Request)

	log.Printf("Produce request: %v\n", r)

	res := &produceAPI.Response{
		Topics: make([]produceAPI.ResponseTopic, 0),
	}

	for _, t := range r.Topics {

		// Preparing topic response
		resTopic := produceAPI.ResponseTopic{
			Topic:      t.Topic,
			Partitions: make([]produceAPI.ResponsePartition, 0),
		}

		// Produce by partition
		for _, p := range t.Partitions {

			log.Printf("============ RecordSet: v=%d\n", p.RecordSet.Version)

			for {
				rs, err := p.RecordSet.Records.ReadRecord()
				if err == io.EOF {
					break
				}

				data, err := protocol.ReadAll(rs.Value)
				if err != nil {
					fmt.Println("failed to read record value: %v", err)
					break
				}

				msg := &Message{
					Time:      time.Now(),
					Headers:   rs.Headers,
					Partition: p.Partition,
					Topic:     t.Topic,
					Data:      data,
				}

				b.Publish(t.Topic, msg)

				resTopic.Partitions = append(resTopic.Partitions, produceAPI.ResponsePartition{
					Partition:      p.Partition,
					ErrorCode:      0,
					BaseOffset:     0,
					LogAppendTime:  0,
					LogStartOffset: 0,
					RecordErrors:   make([]produceAPI.ResponseError, 0),
					ErrorMessage:   "",
				})
			}
		}

		res.Topics = append(res.Topics, resTopic)
	}

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}

func (b *Broker) handleFetch(conn net.Conn, req *Request) {

	r := req.Msg.(*fetchAPI.Request)

	res := &fetchAPI.Response{
		ErrorCode: 0,
		SessionID: r.SessionID,
		Topics:    make([]fetchAPI.ResponseTopic, 0),
	}

	// Grouping by stream (the first part)
	streams := make(map[string][]*fetchAPI.RequestTopic, 0)
	for _, t := range r.Topics {
		parts := strings.Split(t.Topic, ".")

		subjects, ok := streams[parts[0]]
		if !ok {
			subjects = make([]*fetchAPI.RequestTopic, 0)
		}

		subjects = append(subjects, &t)
		streams[parts[0]] = subjects
	}

	// Ensure consumer for streams
	for stream, topics := range streams {

		seq := int64(0)
		subjects := make([]string, 0)
		for _, t := range topics {

			for _, p := range t.Partitions {
				if seq < p.FetchOffset {
					seq = p.FetchOffset
				}
			}

			subjects = append(subjects, t.Topic)
		}

		err := b.store.EnsureConsumer(r.SessionID, stream, subjects, seq)
		if err != nil {
			log.Printf("failed to ensure consumer: %v\n", err)
			res.ErrorCode = 1
			protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
			return
		}

		// Pull messages from store
		msgs, err := b.store.Fetch(r.SessionID, 100)
		if err != nil {
			log.Printf("failed to fetch messages: %v\n", err)
			protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
			return
		}

		// Preparing responses
		for _, t := range topics {

			log.Printf("Fetching messages for topic: %s\n", t.Topic)

			resTopic := fetchAPI.ResponseTopic{
				Topic:      t.Topic,
				Partitions: make([]fetchAPI.ResponsePartition, 0),
			}

			// Preparing records
			records := make([]protocol.Record, 0)
			for _, m := range msgs {
				if m.Topic != t.Topic {
					continue
				}

				log.Printf("============ Fetched message: %d %s\n", m.Seq, string(m.Data))

				records = append(records, protocol.Record{
					Offset:  int64(m.Seq),
					Time:    m.Time,
					Key:     protocol.NewBytes([]byte{}),
					Value:   protocol.NewBytes(m.Data),
					Headers: m.Headers,
				})
			}

			// Put records into one partition
			p := fetchAPI.ResponsePartition{
				Partition:            0,
				ErrorCode:            0,
				HighWatermark:        100000,
				LogStartOffset:       0,
				PreferredReadReplica: 0,
				AbortedTransactions:  make([]fetchAPI.ResponseTransaction, 0),
				RecordSet: protocol.RecordSet{
					Version: 2,
					//					Attributes: protocol.Attributes(protocol.Gzip),
					Records: protocol.NewRecordReader(records...),
				},
			}

			//TODO: support multiple partitions
			resTopic.Partitions = append(resTopic.Partitions, p)

			res.Topics = append(res.Topics, resTopic)
		}
	}

	log.Printf("Fetch response: %v\n", res)

	protocol.WriteResponse(conn, req.ApiVersion, req.CorrelationID, res)
}
