package main

import (
	"fmt"
	"os"
	"time"
)

// 检查文件是否存在
func checkFileIsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func main() {
	readmePath := "./README.md"
	exist, err := checkFileIsExist(readmePath)
	if exist {
		fmt.Printf("文件 '%s' 存在。\n", readmePath)
		// 读取文件并打印
		file, err := os.OpenFile(readmePath, os.O_RDWR|os.O_APPEND, 066)
		if err != nil {
			fmt.Printf("打开文件 '%s' 失败: %v\n", readmePath, err)
			return
		}
		content, err := os.ReadFile(readmePath)
		if err != nil {
			fmt.Printf("读取文件 '%s' 失败: %v\n", readmePath, err)
			return
		}
		fmt.Printf("文件内容为:%s\n", string(content))
		// 写入文件
		_, err = file.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Printf("写入文件 '%s' 失败: %v\n", readmePath, err)
			return
		}
		fmt.Println("写入成功")
		defer func() {
			fmt.Println("文件关闭")
			err := file.Close()
			if err != nil {
				fmt.Printf("文件关闭失败: %v\n", err)
				return
			}
		}()
	} else {
		fmt.Printf("文件 '%s' 不存在。\n", readmePath)
		//	创建文件
		file, err := os.Create(readmePath)
		if err != nil {
			fmt.Printf("创建文件 '%s' 失败: %v\n", readmePath, err)
			return
		}
		defer func() {
			fmt.Println("文件关闭")
			err := file.Close()
			if err != nil {
				fmt.Printf("文件关闭失败: %v\n", err)
				return
			}
		}()
	}
	if err != nil {
		fmt.Printf("检查文件时出错: %v\n", err)
		return
	}
}
