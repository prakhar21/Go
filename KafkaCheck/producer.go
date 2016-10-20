package main 

import (
	
	"log"
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
	"bufio"
	"os"

)


func check(e error) {

	if e != nil {

		log.Fatal(e)
	
	}
}

func main() {

    fmt.Println("Reading file...")
    infile, err := os.Open("data.txt")
    check(err)
    
    scanner := bufio.NewScanner(infile)

    config := sarama.NewConfig()
    brokers := []string{"localhost:9092"}
    topics := "test-go"
    
    message := &sarama.ProducerMessage{Topic: topics, Partition: 0}
    producer, err := sarama.NewSyncProducer(brokers, config)
    check(err)

    for scanner.Scan() {
        message.Value = sarama.StringEncoder(scanner.Text())
    	_, _, err := producer.SendMessage(message)
    	check(err)
	
	}

	fmt.Println("All the messages have been sent...")

}