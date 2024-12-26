package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// 全局上下文，用于管理超时
var ctx = context.Background()

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378", // Redis 服务地址
		Password: "",               // 如果没有密码，设为空字符串
		DB:       0,                // 使用默认数据库
		PoolSize: 10,               // 连接池大小
	})

	// 测试是否连接成功
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Printf("连接 Redis 失败: %v\n", err)
	} else {
		fmt.Println("连接 Redis 成功")
		//	清除Redis
		rdb.FlushDB(ctx)
	}

	err = rdb.Set(ctx, "test", "测试内容", 0).Err()
	if err != nil {
		fmt.Printf("设置键值对失败: %v\n", err)
	}
	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Printf("获取键值对失败: %v\n", err)
	} else {
		fmt.Println("获取到的键值对:", val)
	}

	if err := rdb.RPush(ctx, "list1", 1, 2, 3).Err(); err != nil {
		fmt.Println("RPush失败")
	} else {
		fmt.Println("RPush成功")
	}

	defer func() {
		if err := rdb.Close(); err != nil {
			panic(err)
		}
	}()
}
