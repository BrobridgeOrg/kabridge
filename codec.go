package main

import "github.com/segmentio/kafka-go/protocol"

type Request struct {
	ApiVersion    int16
	CorrelationID int32
	ClientID      string
	Msg           protocol.Message
	Err           error
}
