package main

import "fmt"

// 变量的声明方式
// 变量有多种声明方式
func main() {
	// 第一种声明方式
	// 通用的变量声明方式
	var s string = "Hello"
	fmt.Println(s)

	// 第二种声明方式
	// 类型推导的方式
	var m = "hi"
	fmt.Println(m)

	// 多变量声明的方式
	var i, j, k int
	i, j, k = 10, 12, 13
	fmt.Println(i, j, k)

	// 短变量声明的方式
	// 短变量声明方式声明的变量会自动进行类型推导
	// 只能在函数内使用
	a := 12
	fmt.Println(a)

	// 多变量的短声明
	b, c, l := 13, 14, 15
	fmt.Println(b, c, l)
	// 这里要注意的是 多变量的短声明可以用于声明不同类型的变量
	// 短变量是声明并赋值
	// 可以使用 _ 用来占位
	x, v := 12, "Hello"
	fmt.Println(x, v)
}
