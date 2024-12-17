package main

import (
	"fmt"
	"time"
)

func producer(channel chan int, id int) {
	for i := 0; i < 5; i++ {
		value := i * id
		fmt.Printf("Producer %d: %d\n", id, value)
		channel <- value
		fmt.Println("Producer: Waiting")
		time.Sleep(1 * time.Second)
	}
	close(channel)
}
func consumer(channel chan int, id int) {
	for value := range channel {
		fmt.Printf("Consumer %d: %d\n", id, value)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel1 := make(chan int, 3)
	go producer(channel1, 1)
	go consumer(channel1, 1)
	time.Sleep(10 * time.Second)
}
