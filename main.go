package main

import (
	"fmt"
	"log"
)

func main() {
	broker := NewBroker()
	address := "localhost:9092"

	go func() {
		if err := broker.Start(address); err != nil {
			log.Fatalf("failed to start broker: %v\n", err)
		}
	}()

	fmt.Printf("Kafka broker running at %s\n", address)
	/*
		// 模擬發佈訊息
		go func() {
			for {
				broker.Publish("example-topic", []byte("Hello Kafka"))
				time.Sleep(2 * time.Second)
			}
		}()
	*/
	// 保持主程式運行
	select {}
}
