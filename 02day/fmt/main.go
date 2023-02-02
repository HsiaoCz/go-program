package main

import "fmt"

func main() {
	// print打印是不换行的
	fmt.Print("Hello")
	// println打印换行
	fmt.Println("哈哈")
	// %T 查看类型
	// %d 十进制数
	// %o 八进制数
	// %x 十六进制数
	// %b 二进制数
	// %v 默认输出
	// %+v 结构体会加点东西
	// %#v 值得go语言显示
	// %% 就是%
	// %f 浮点数表示
	// %p 指针
	// %e 科学计数法表示
	// %t 布尔值
	// %q unicode表示
	// %b 在浮点数得时候 是浮点数没有小数部分得二进制表示
	// %q 在字符串中 将字符串表示成双引号得

	num := 90
	fmt.Printf("%d%%\n", num)
	s := "Hello"
	fmt.Printf("%q\n", s)
	// %7s 右对齐
	// %-7s 左对齐
	fmt.Printf("%7s\n", s)
	fmt.Printf("%-7s\n", s)
	// 获取用户输入
	var m string
	fmt.Scanln(&m)
	fmt.Println(m)
	// Scanf 格式化输入
	var (
		name  string
		age   int
		class string
	)
	// Scanf()按照空格分割
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Printf("%s:%d:%s\n", name, age, class)
	// Spintf() 拼接字符串
	a := fmt.Sprintf("%s:%s", "zhangsan", "hello")
	fmt.Println(a)
}
