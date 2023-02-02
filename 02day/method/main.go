package main

import "fmt"

// 方法 方法是作用于特定类型的函数

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

// 方法是作用于特定类型的函数
// 接收者表示的调用该方法的具体类型变量，多用类型首字母的小写
// 方法的值接受者和指针接收者的区别
// 根本区别在于指针接收者的方法和值接收者的方法的方法名是不一样的

type Man struct{}

func (m Man) Say() {
	fmt.Println("Hello")
}

type Woman struct{}

func (m *Woman) Say() {
	fmt.Println("Hello")
}

// 在这Man的Say()方法 方法名为Man.Say()
// 在这Woman的Say()方法 方法名为(*Woman).Say()
// 当然在这 使用*Man()也能调用Man.Say()方法 使用Woman也能调用(*Woman).Say()方法
// 这是因为编译器进行了隐式转换，
func main() {
	d := newDog("旺财")
	d.wang()
}
