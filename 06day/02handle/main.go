package main

import (
	"net/http"
)

// 多个Handler
// 不指定Server struct里面的Handler字段值
// 可以使用http.Handle将某个Handler附加到DefaultServeMux
// http包有一个Handle函数
// ServerMux struct 也有一个Handle方法
// 如果你调用http.Handle，实际上调用的是DefaultServeMux上的Handle方法
// DefaultServeMux就是ServerMux的指针变量
// http.Handle
// func Handle(pattern string,handler Handler)
// type Handler interface{
//   ServeHTTP(ResponseWriter,*Request)
// }

// 使用handle函数可以添加多个Handler,每个Handler对应不同的路由
type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// 再添加一个Handler
type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About"))
}

func main() {
	// 使用handle函数注册请求的handler
	h := helloHandler{}
	a := aboutHandler{}
	http.Handle("/hello", &h)
	http.Handle("/about", &a)
	http.HandleFunc("/welcome", welcome)
	http.Handle("/user", http.HandlerFunc(user))
	server := http.Server{
		Addr:    "172.26.215.133:9090",
		Handler: nil,
	}
	server.ListenAndServe()
}

// http.HandleFunc
// Handler 函数就是那些行为与Handler类似的函数
// Handler 函数的签名与ServeHTTP方法的签名一样，接收两个参数：
// 第一个http.ResponseWriter
// 一个指向http.Request的指针

// HandleFunc 这个函数 可以将某个具有适当签名的函数f,设配成一个Handler，而这个Handler具有方法f
// handleFunc 在内部调用HandlerFunc这个函数类型函数，将函数适配成一个handler
// mux.Handle(pattern,HandlerFunc(handler)) 这里是HandleFunc做的调用，本质上还是调用Handler
// HandlerFunc 是一个函数类型，相当于一个适配器
// HandlerFunc 为什么能将一个函数适配成一个handler呢？
// HandlerFunc 是一个类型，它实现了Handler接口，它本身就是一个Handler
// 它就一个具有wr签名的函数做了类型转换 转换成了handler

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User is comming"))
}
