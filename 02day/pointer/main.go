package main

import "fmt"

// 指针 指向变量的值
// 指针的值是一个变量的地址，一个指针指示的是值所保存的位置
func main() {
	a := 10
	b := &a
	// 指针就是 &取地址
	// * 根据地址取值
	// 这里b保存的是a 的地址
	// 改变*b 的值 a的值也会改变
	fmt.Println(*b)
	fmt.Println("&b:", b)
	*b = 12
	fmt.Println(a)
}
