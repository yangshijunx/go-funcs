package main

import (
	"fmt"
	"time"
)

//在 Go 语言中，关于时间处理的知识点主要包括以下几点：
//time 包：Go 标准库提供了 time 包来处理时间和日期相关的操作。
//time.Time 结构体：表示一个时间点，精确到纳秒。常用方法包括 Now() 获取当前时间，Format() 将时间格式化为字符串，Parse() 解析时间字符串等。
//时间格式化：使用 time.Time.Format(layout string) 方法将时间格式化为字符串。格式化布局字符串通常使用 2006-01-02 15:04:05 这样的固定格式。
//时间解析：使用 time.Parse(layout string, value string) 方法将字符串解析为 time.Time 对象。
//时间间隔：time.Duration 表示两个时间点之间的间隔，单位可以是纳秒、微秒、毫秒、秒等。常用方法包括 time.Second、time.Millisecond 等。
//定时器：time.Timer 和 time.After 可以用来创建定时器，执行延迟任务或周期性任务。
//计时器：time.Ticker 用于定期发送时间信号，常用于周期性任务。
//时间计算：可以使用 time.Time.Add(d time.Duration) 方法在时间点上增加或减少时间间隔。
//时间比较：使用 time.Time.Before(t time.Time)、time.Time.After(t time.Time) 和 time.Time.Equal(t time.Time) 方法进行时间比较。

// 打印当前时间
func printNow() {
	// 打印当前时间
	fmt.Printf("当前的时间是：%d\n", time.Now().Unix())
	fmt.Printf("时间格式转换为XX年MM月dd日 HH时mm分ss秒：%s\n", time.Now().Format("2006年01月02 15时04分05秒"))
}

func trickTime() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker")
		}
	}
}

func duration() {
	// 1. 创建一个时间对象
	now := time.Now()
	// 2. 延时一秒
	time.Sleep(time.Second)
	// 3. 获取当前时间对象
	now2 := time.Now()
	// 4. 获取时间间隔
	fmt.Printf("时间间隔：%d\n", now2.Sub(now))
}

// after 用法
func after() {
	<-time.After(time.Second)
	fmt.Println("1秒之后")
	afterChan := time.After(3 * time.Second)
	select {
	case <-afterChan:
		println("3秒之后")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
}
func main() {
	printNow()
	duration()
	after()
	trickTime()
}
