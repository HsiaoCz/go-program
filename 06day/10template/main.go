package main

import (
	"html/template"
	"net/http"
)

// Web模板就是预先设计好的HTML页面，它可以被模板引擎反复的使用，来产生HTML页面
// go标准库提供了text/template，html/template两个模板库
// 大多数的Go的web框架都使用这些库作为默认的模板引擎
// 所谓的模板引擎，就是可以合并模板与上下文数据，产生最终的HTML

// 两种理想的模板引擎
// 1、无逻辑模板引擎
// 通过占位符、动态数据替换到模板中
// 不做任何逻辑处理，只做字符串的替换
// 处理完全由handler来完成
// 目标是展示层和逻辑完全分离

// 2、逻辑嵌入模板引擎
// 编程语言被嵌入到模板中
// 在运行时由模板引擎来执行，也包含替换功能
// 功能强大
// 逻辑代码遍布handler和模板，难以维护

// go的模板引擎
// 主要使用text/template，html相关部分使用html/template，是一个混合体
// 模板可以完全无逻辑，但是又具有足够的嵌入特性
// 和大多数模板引擎一样，Go Web的模板位于无逻辑和嵌入逻辑之间的某个地方

// go 模板引擎的工作原理
// 在web应用中，通常是由handler来触发模板引擎
// handler 调用模板引擎，并将使用的模板传递给引擎，通常是一组模板文件和动态数据
// 模板引擎生成HTML，并将其写入到ResponseWriter
// ResponseWriter再将它加入到HTTP响应中，返回给客户端

// 模板
// 模板必须是可读的文本格式，扩展名任意。对于Web应用通常就是HTML
// 不过里面会内置一些命令(叫作action)
// text/template是通用模板引擎，html/template是HTML模板引擎
// action 是位于两层花括号之间:{{.}}
// 这里的.就是一个action
// 它可以命令模板引擎将其替换成一个值

// 如何使用模板引擎
// 解析模板源(可以是字符串或模板文件)，从而创建一个解析好的模板的struct
// 执行解析好的模板，并传入ResponseWriter和数据
// 这会触发模板引擎组合解析好的模板和数据，来产生最终的HTML，传递给ResponseWriter
func main() {
	http.HandleFunc("/template", templateExample)
	server := http.Server{
		Addr:    ":9090",
		Handler: nil,
	}
	server.ListenAndServe()
}

// 使用模板
func templateExample(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("template.html")
	name := "zhangsan"
	// 执行模板
	t.Execute(w, name)
}
