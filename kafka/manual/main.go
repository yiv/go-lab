package main

import (
	"fmt"
	. "github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errc := make(chan error)
	// Trap SIGINT to trigger a shutdown.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("edwin #0")

	config := NewConfig()

	// First, we tell the producer that we are going to partition ourselves.
	config.Producer.Partitioner = NewManualPartitioner

	producer, err := NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close producer:", err)
		}
	}()

	// Now, we set the Partition field of the ProducerMessage struct.
	msg := &ProducerMessage{Topic: "test", Partition: 6, Value: StringEncoder("test")}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to produce message to kafka cluster.")
	}

	if partition != 6 {
		log.Fatal("Message should have been produced to partition 6!")
	}

	log.Printf("Produced message to partition %d with offset %d", partition, offset)
}
