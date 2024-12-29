package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var ctx = context.Background()

func CloseRedis(rdb *redis.Client) {
	err := rdb.Close()
	if err != nil {
		panic(err)
	} else {
		println("redis close")
	}
}

func TestRedis(rdb *redis.Client) {
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		println("redis connect success")
	}
}

func raceRedis(rdb *redis.Client) {
	key := "秒杀-iphone"
	//	清除key对应的值
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	} else {
		println("清除key对应的值")
	}
	wg := sync.WaitGroup{}
	const P = 100
	for i := 0; i < P; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				value := fmt.Sprintf("%d抢到了", j)
				cmd := rdb.SetNX(ctx, key, value, 0)
				if cmd.Err() != nil {
					panic(cmd.Err())
				} else {
					if cmd.Val() {
						fmt.Printf("%d抢到了\n", j)
					}
				}
			}
		}()
	}
	wg.Wait()
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	TestRedis(rdb)
	raceRedis(rdb)
	defer func() {
		CloseRedis(rdb)
	}()
}
