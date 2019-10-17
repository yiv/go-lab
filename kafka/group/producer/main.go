package main

import (
	"fmt"
	. "github.com/Shopify/sarama"
	"log"
	"time"
)

func main() {
	config := NewConfig()
	config.Producer.Partitioner = NewManualPartitioner
	producer, err := NewAsyncProducer([]string{"10.72.17.30:19092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = producer.Close() }()
	
	var enqueued, errors int
	var partition int32 = 0
	for {
		time.Sleep(time.Millisecond * 3000)
		select {
		case producer.Input() <- &ProducerMessage{
			Topic:     "EventUserAddBookComment",
			Key:       nil,
			Value:     StringEncoder(fmt.Sprintf("partition %v testing %v",partition, time.Now().Unix())),
			Partition: partition,
		}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		}
	}
}
