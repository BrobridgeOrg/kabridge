package main

import (
	"time"

	"github.com/segmentio/kafka-go/protocol"
)

type Message struct {
	Seq       uint64
	Time      time.Time
	Stream    string
	Partition int32
	Topic     string
	Headers   []protocol.Header
	Data      []byte
}
