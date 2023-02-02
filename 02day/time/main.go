package main

import (
	"fmt"
	"time"
)

// 获取时间
// 时间戳，从1970 1.1 0:00:00到现在经过的毫秒数
// 计时器 tick
func Now() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	ticker := time.Tick(1 * time.Second)
	// tick 本质是一个通道
	for i := range ticker {
		fmt.Println(i)
	}
}

// 时间格式化操作
// 时间格式化操作，可以只格式化时分秒，也可以只格式化日期
// time.Prase()解析时间
func main() {
	Now()
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}
