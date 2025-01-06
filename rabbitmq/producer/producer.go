package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	// 创建通道
	ch, err := conn.Channel()
	// 负载均衡
	err = ch.Qos(
		2,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		failOnError(err, "Failed to open a channel")
		panic(err)
	}
	defer func() {
		if err := ch.Close(); err != nil {
			panic(err)
		}
	}()
	// 声明队列如果不存在就创建
	q, err := ch.QueueDeclare(
		"test",
		false, // 是否持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "声明队列失败")
		panic(err)
	}
	for i := 0; i < 100; i++ {
		body := "test-mq"
		body = body + "-" + strconv.Itoa(i)
		fmt.Printf("push消息：%s\n", body)
		err := ch.Publish(
			"",     // 交换器名称
			q.Name, // 队列名称
			false,  // 延迟消息
			false,  // 持久化消息
			amqp.Publishing{
				DeliveryMode: amqp.Transient, // 瞬态消息
				ContentType:  "text/plain",
				Body:         []byte(body),
			},
		)
		if err != nil {
			failOnError(err, "push消息失败")
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}
}
