package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	producer, err := sarama.NewAsyncProducer([]string{"192.168.1.205:19092"}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("edwin #1")

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("edwin #2")

	fmt.Println("edwin #3")

	var enqueued, errors int

	go func() {
		for {
			select {
			case producer.Input() <- &sarama.ProducerMessage{Topic: "test", Partition: 0, Key: nil, Value: sarama.StringEncoder("testing 123")}:
				enqueued++
				fmt.Println("edwin #4")
				time.Sleep(1 * time.Second)
			case err := <-producer.Errors():
				log.Println("Failed to produce message", err)
				errors++
				fmt.Println("edwin #5")
			}
		}
	}()

	log.Printf("Enqueued: %d; errors: %d; err :%v\n", enqueued, errors, <-errc)
}
