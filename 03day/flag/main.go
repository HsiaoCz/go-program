package main

import (
	"flag"
	"fmt"
	"os"
)

// flag包实现了命令行参数的解析
// os.args 简单的获取命令行参数
func main() {
	orgsGet()
	flagGet()
}

// os.args 是一个字符串切片
func orgsGet() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args [%d]=%v\n", index, arg)
		}
	}
}

func flagGet() {
	name := flag.String("name", "zhangsan", "请输入姓名:")
	flag.Parse()
	fmt.Printf("%v\n", name)
	var s string
	flag.StringVar(&s, "name", "张三", "请输入名字:")
	flag.Parse()
	fmt.Println(s)
}
