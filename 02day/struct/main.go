package main

import "fmt"

// 结构体是值类型
type person struct {
	name   string
	gender string
}

// 函数传参 传递得值拷贝
func f(p person) {
	p.gender = "男"
}

// 使用指针可以修改值
func ff(p *person) {
	// 根据内存地址找到那个原变量，修改得是原变量的值
	// (*p).gender="男"
	// 也可以使用这种
	p.gender = "男"
}

func main() {
	p := person{
		name:   "张三",
		gender: "n",
	}
	f(p)
	fmt.Println(p)
	ff(&p)
	fmt.Println(p.name, p.gender)
}
