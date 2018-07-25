package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("edwin #0 ")

	consumer, err := sarama.NewConsumer([]string{"192.168.1.205:9092"}, nil)
	if err != nil {
		panic(err)
	}

	{
		topics, err := consumer.Topics()
		if err != nil {
			panic(err)
		}
		fmt.Println("edwin #3 ", topics)
		parts, err := consumer.Partitions(topics[0])
		if err != nil {
			panic(err)
		}
		fmt.Println("edwin #3 ", parts)
	}

	fmt.Println("edwin #1 ")

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	fmt.Println("edwin #2 ")

	partitionConsumer, err := consumer.ConsumePartition("EventCreatePattiRobot", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	fmt.Println("edwin #3 ")

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	fmt.Println("edwin #4 ")

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
			fmt.Println("edwin #5 ")
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
