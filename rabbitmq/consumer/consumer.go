package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// 封装消息处理逻辑
func processMessages(messages <-chan amqp.Delivery, consumerID int) {
	for msg := range messages {
		fmt.Printf("消费者 %d 收到消息: %s\n", consumerID, msg.Body)
		err := msg.Ack(false) // 手动确认消息
		if err != nil {
			log.Printf("消费者 %d 确认消息失败: %s", consumerID, err)
		}
		time.Sleep(time.Second * 1) // 模拟消息处理耗时
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
		false,  // auto-ack 或者称之为 no-ack 就是不需要确认
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}
	for i := 0; i < 2; i++ {
		go processMessages(messages, i)
	}
	// 主进程阻塞，防止程序退出
	log.Println("等待消息中...")
	select {} // 无限阻塞
	//for msg := range messages {
	//	fmt.Printf("第一个Received a message: %s\n", msg.Body)
	//	err := msg.Ack(false)
	//	// err := msg.Ack(true) // 批量ack确认
	//	time.Sleep(time.Second * 1)
	//	if err != nil {
	//		failOnError(err, "Failed to ack")
	//	}
	//}
}
