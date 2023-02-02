package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// action 就是Go模板中嵌入的命令，位于两组花括号之间{{xxx}}
// . 就是一个action，而且是最重要的一个。它代表了传入的数据
// action可以分为五类
// 条件类 迭代/遍历类 设置类 包含类 定义类

// 条件类Action
// {{ if arg }}
// some content
// {{ end }}
// 这种条件action还可以加个{{ else }}
func main() {
	http.HandleFunc("/time", IfExample)
	http.HandleFunc("/range", RangeExample)
	http.HandleFunc("/bao", BaoExample)
	http.ListenAndServe(":9090", nil)
}

func IfExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

// 迭代/遍历Action
// {{ range array }}
// Dot is set to the element {{.}}
// 这时候这个.表示的是每次遍历的数据
// {{end}}
// 这类action可以用来遍历数组/slice/map或者channel等数据结构
// 这种action 如果遍历的集合是空的，还可以在中间加一个else
func RangeExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	daysOfWeek := []string{"mon", "Tue", "wed", "thu", "fri", "sat", "sun"}
	t.Execute(w, daysOfWeek)
}

// 设置action
// {{ with arg }}
// . 设置值
// {{ end }}
// 它允许在指定的范围内，让"."来表示其他指定的值
// 简单说 就是把点变成我们自己指定的值了
// 如果with后面为空 可以使用else来指定表示回落
// {{ else }}

func WithExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello")
}

// 包含类的Action
// {{template "name"}}
// 它允许你在模板中包含其他的模板
func BaoExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "t1.html")
	t.Execute(w, "Hello")
}

// 定义的ction define action
