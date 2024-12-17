package main

import (
	"fmt"
)

// Student 结构体定义
type Student struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	// 获取数据库实例
	db := NewDB()

	// 自动迁移表结构
	db.AutoMigrate(&Student{})

	// 创建数据
	student := Student{Name: "张三", Age: 25}
	db.Create(&student)

	// 查询数据
	var result Student
	db.First(&result, student.ID)
	fmt.Printf("查询到的学生: %+v\n", result)

	// 更新数据
	db.Model(&result).Update("Age", 26)

	// 删除数据
	//db.Delete(&result)
}
