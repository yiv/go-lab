package main

import (
	"context"
	"fmt"
	. "github.com/Shopify/sarama"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess ConsumerGroupSession, claim ConsumerGroupClaim) error {
	fmt.Println("start consumer")
	defer fmt.Println("consumer exit")
	
	for msg := range claim.Messages() {
		fmt.Println("edwin #66")
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	// Init config, specify appropriate version
	config := NewConfig()
	config.Version = V2_3_0_0
	config.Consumer.Return.Errors = true
	
	// Start with a client
	client, err := NewClient([]string{"10.72.17.30:19092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()
	
	// Start a new consumer group
	group, err := NewConsumerGroupFromClient("group1", client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()
	
	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()
	
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		fmt.Println("group re-balancing")
		topics := []string{"test"}
		handler := exampleConsumerGroupHandler{}
		
		if err = group.Consume(ctx, topics, handler); err != nil {
			panic(err)
		}
	}
	
}