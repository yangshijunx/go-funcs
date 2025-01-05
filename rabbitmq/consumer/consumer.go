package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError(err, "消费者连接rabbitmq失败")
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %s", err)
		}
	}()
	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "消费者创建channel失败")
		panic(err)
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close channel: %s", err)
		}
	}()
	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
	}
	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}
	for msg := range messages {
		fmt.Printf("Received a message: %s\n", msg.Body)
	}
}
