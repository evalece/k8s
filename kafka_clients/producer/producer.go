package main

//go run producer.go

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Set up a Kafka writer (acts as a producer)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"}, // later drive by env var
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	})

	// Send 10 messages
	for i := 0; i < 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("traceID%d", i)),
			Value: []byte(fmt.Sprintf("Hello Kafka from Go #%d", i)),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf(" failed to write message: %v", err)
		} else {
			log.Printf(" sent: %s", msg.Value)
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Clean up
	if err := writer.Close(); err != nil {
		log.Fatalf(" failed to close writer: %v", err)
	}
}
