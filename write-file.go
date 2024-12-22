package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func writeWithoutBuffer(filename string) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	// 使用 defer 关闭文件并检查关闭时的错误
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file %s: %v\n", filename, err)
		}
	}()

	// 写入10000行数据
	for i := 0; i < 10000; i++ {
		_, err := file.WriteString(fmt.Sprintf("This is line %d\n", i))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeWithBuffer(filename string) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file %s: %v\n", filename, err)
		}
	}()

	// 使用bufio.Writer来写入文件
	writer := bufio.NewWriter(file)
	// 写入10000行数据
	for i := 0; i < 10000; i++ {
		_, err := writer.WriteString(fmt.Sprintf("This is line %d\n", i))
		if err != nil {
			return err
		}
	}
	// 必须调用 Flush 来确保缓冲区的内容被写入文件
	return writer.Flush()
}

func main() {
	// 记录不使用buffer的写入时间
	start := time.Now()
	err := writeWithoutBuffer("no_buffer.txt")
	if err != nil {
		fmt.Println("Error writing without buffer:", err)
		return
	}
	fmt.Printf("Time taken to write without buffer: %v\n", time.Since(start))

	// 记录使用buffer的写入时间
	start = time.Now()
	err = writeWithBuffer("with_buffer.txt")
	if err != nil {
		fmt.Println("Error writing with buffer:", err)
		return
	}
	fmt.Printf("Time taken to write with buffer: %v\n", time.Since(start))
}
