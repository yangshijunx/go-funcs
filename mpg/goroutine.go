package main

import (
	"fmt"
	"sync"
)

// 打印hello

func printHello() {
	println("hello")
}

func printHelloAsync(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello async")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 启动 Goroutine 通过 go 关键字。它会将指定的函数执行放到后台进行并发处理。
		wg.Add(1)
		go printHello()
		// 这样做可以确保每个 Goroutine 有足够的时间来执行 printHello 函数并打印 "hello"。
		//time.Sleep(time.Second * 1)
		go printHelloAsync(&wg)
	}
	wg.Wait()
}
