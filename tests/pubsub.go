package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const topic = "my-stream.my-topic"
const partition = 0

func consume() error {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		//		GroupID:   "my-group",
		Topic:     topic,
		Partition: partition,
	})
	r.SetOffset(0)

	fmt.Println(r.Offset())

	go func() {

		log.Println("fetching messages...")

		ctx := context.Background()
		for {

			//m, err := r.FetchMessage(ctx)
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Fatal("failed to fetch message:", err)
				break
			}

			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			/*
				if err := r.CommitMessages(ctx, m); err != nil {
					log.Fatal("failed to commit messages:", err)
				}
			*/
		}
	}()

	log.Println("waiting for messages...")

	return nil
}

func produce(messages ...kafka.Message) error {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.Hash{},
	})

	w.AllowAutoTopicCreation = true

	ctx := context.Background()
	err := w.WriteMessages(ctx, messages...)
	if err != nil {
		return err
	}

	log.Println("sent")

	return nil
}

func main() {

	log.Println("starting producer...")
	err := produce(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("starting consumer...")
	err = consume()
	if err != nil {
		log.Fatal("failed to start consumer:", err)
	}

	select {}
}
