package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// phone 不序列化
	Phone string `json:"-"`
	// email 值为空时不序列化
	Email string `json:"email,omitempty"`
	// 不加json 标签，默认平铺
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func main() {
	user := User{
		Name:  "John",
		Age:   30,
		Phone: "123456789",
		Email: "",
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}
	fmt.Print(user)
	//  打印序列化的结果
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("序列化的结果：%s", string(jsonData))
}
