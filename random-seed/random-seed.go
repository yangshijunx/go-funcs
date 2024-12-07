package main

import (
	"math/rand"
)

func createRandom() {
	rand.Seed(42)
	// 生成一个随机数
	randomNumber := rand.Intn(100)
	println(randomNumber)
}

// 多线程生成随机数
func createRandomMultiThread() {
	source := rand.NewSource(42)
	random := rand.New(source)
	println(random.Intn(100))
}

func main() {
	createRandom()
	createRandomMultiThread()
}
