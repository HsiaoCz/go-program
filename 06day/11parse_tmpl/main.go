package main

import (
	"html/template"
	"net/http"
)

// 解析模板
// ParseFiles
// ParseGlob
// Parse
// 解析模板文件，并创建一个解析好的模板struct，后续可以被执行
// ParseFiles函数是Template struct上ParseFiles方法的简便调用
// 调用ParseFiles后，会创建一个新的模板，模板的名字是文件名

func main() {
	// 整个函数的调用 做了这么几步
	// 先把模板文件里的内容读成字符串
	// 然后使用模板文件名，不包含路径创建一个新的模板
	// 然后调用模板里边的一个Parse()来解析模板文件内容
	// t,err:=template.ParseFiles("tmpl.html")
	//  if err != nil {
	// 	log.Fatal(err)
	//  }
	//  t.Execute()

	// 还可以使用这种方式
	// 使用template.New()创建一个模板，这个模板的名字就是传进去的文件名 然后使用模板上的解析文件方法
	// t := template.New("tmpl.html")
	// t, _ = t.ParseFiles("tmpl.html")
	// t.Execute()
	// New 函数 ParseFiles的参数数量可变，但只返回一个模板
	// 当解析多个文件时，第一个文件作为返回的模板(名，内容)，其余的作为map，供后续执行使用

	// ParseGlob()使用模式匹配来解析特点的文件
	// t,_:=template.ParseGlob("*.html")

	// Parse() 这个函数可以解析字符串模板，其他方式最终都会调用Parse
	http.HandleFunc("/template", ExecuteTemplate)
	server := http.Server{
		Addr:    ":9090",
		Handler: nil,
	}
	server.ListenAndServe()
}

// lookup方法
// 通过模板名来寻找模板，如果没找到就返回nil
// Must 函数 可以包裹一个函数，返回一个模板的指针和一个错误
// 如果错误不为nil，那么就panic
// 解析模板的时候可能会发生错误
// 使用must函数来包裹模板函数 减少错误处理

// 执行模板 Execute()
// 参数是ResponseWriter、数据
// 单模板很适用，当有多个模板的时候，只使用第一个模板
// ExecuteTemplate() 参数是：ResponseWriter、模板名、数据
// 适用于模板集

func ExecuteTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html")
	t.Execute(w, "Hello")
	tm, _ := template.ParseFiles("t1.html", "t2.html")
	tm.ExecuteTemplate(w, "t2.html", "Hello")
}
