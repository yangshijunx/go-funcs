package main

import (
	"context"
	"fmt"
	"sync"
	"time"

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
		// 清除Redis
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

	// 获取list1
	if list1, err := rdb.LRange(ctx, "list1", 0, -1).Result(); err != nil { // 获取list1
		fmt.Println("获取list1失败")
	} else {
		fmt.Println("获取list1成功:", list1)
	}

	// Lpush 和 Rpush 的区别
	// Lpush 是从左边插入，Rpush 是从右边插入
	// 为什么左插入顺序是对的，而右插入顺序是反的？
	if err := rdb.LPush(ctx, "list2", 1, 2, 3).Err(); err != nil {
		fmt.Println("LPush失败")
	} else {
		fmt.Println("LPush成功")
	}

	// 获取list2
	if list2, err := rdb.LRange(ctx, "list2", 0, -1).Result(); err != nil {
		fmt.Println("获取list2失败")
	} else {
		fmt.Println("获取list2成功:", list2)
	}

	// 创建一个 WaitGroup 用来等待 goroutine 完成
	var wg sync.WaitGroup
	pubsub := rdb.Subscribe(ctx, "mychannel")
	wg.Add(1) // 在等待组中增加一个任务

	// 启动 goroutine 接收消息
	go func() {
		defer func() {
			fmt.Println("关闭订阅")
		}()
		for {
			// 等待消息
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				fmt.Println("接收消息失败:", err)
				break
			}
			fmt.Println("接收到消息:", msg.Channel, msg.Payload)
			wg.Done()
		}
	}()

	// 等待一些时间，确保订阅已经开始
	time.Sleep(500 * time.Millisecond)

	// 发布消息
	err = rdb.Publish(ctx, "mychannel", "hello").Err()
	if err != nil {
		fmt.Println("发布消息失败:", err)
	} else {
		fmt.Println("发布消息成功")
	}

	// 等待消息接收完成
	wg.Wait()

	// 清理资源
	defer func() {
		pubsub.Close()
		if err := rdb.Close(); err != nil {
			panic(err)
		}
	}()
}
