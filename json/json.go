package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "John",
		Age:  30,
	}
	fmt.Print(user)
	//  打印序列化的结果
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("序列化的结果：%s", string(jsonData))
}
