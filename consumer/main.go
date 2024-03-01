package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "test-kafka-topic" 
	time.Sleep(time.Minute)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "my-go-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	c.SubscribeTopics([]string{topic}, nil)

	// Handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		case ev := <-c.Events():
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Message: %s (topic %s)\n", string(e.Value), *e.TopicPartition.Topic)
			}
		}
	}

	c.Close()
}

//go get -u github.com/confluentinc/confluent-kafka-go/kafka