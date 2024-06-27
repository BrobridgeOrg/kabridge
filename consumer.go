package main

type Consumer struct {
	Stream   string
	Name     string
	Seq      int64
	Subjects []string
}

func NewConsumer() *Consumer {
	return &Consumer{
		Subjects: make([]string, 0),
	}
}
