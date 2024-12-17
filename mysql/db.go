package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// NewDB 返回数据库实例
func NewDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gin-blob?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
