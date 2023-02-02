package main

import "fmt"

// 先写个经典的hello World
// 使用go run hello.go 简单运行程序
// 使用go build hello.go 将编译程序
// 使用 ./hello 运行
// mian() 作为程序的入口 只能在main包里
// mian 包定义一个独立的可执行程序，不是库
// 使用import 来导入其他的包
// { 必须和func在同一行
func main() {
	fmt.Println("Hello World")
}
