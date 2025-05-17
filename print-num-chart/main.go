package main

import (
	"fmt"
)

func main() {
	numChan := make(chan struct{})
	letterChan := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Print(i)
			letterChan <- struct{}{}
		}
	}()

	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			<-letterChan
			fmt.Print(string(i))
			if i == 'Z' {
				done <- struct{}{}
			}
			numChan <- struct{}{}
		}
	}()

	// 开始执行，先打印数字
	numChan <- struct{}{}
	<-done
}
