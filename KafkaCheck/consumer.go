package main 

import (
	
	"log"
	"fmt"
	"gopkg.in/Shopify/sarama.v1"

)


func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

func main() {

    fmt.Println("Creating Consumer...")

	config := sarama.NewConfig()
	brokers := []string{"localhost:9092"}
    topics := "test-go"

    master, err := sarama.NewConsumer(brokers, config)
    check(err)

    consumer, err := master.ConsumePartition(topics, 0, sarama.OffsetOldest)
    check(err)


    for m := range consumer.Messages() {
		fmt.Println(string(m.Value))
	}
}