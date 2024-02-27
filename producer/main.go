package main

import (
	"fmt"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "my-kafka-topic" // Replace if needed

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	// Send a message
	message := "Hello from Go Producer!"
	deliveryChan := make(chan kafka.Event)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	// Report success or failure of message delivery
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to %v\n", m.TopicPartition)
	}

	close(deliveryChan)
}
