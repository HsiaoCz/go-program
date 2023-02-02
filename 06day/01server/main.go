package main

import "net/http"

// go语言是如何处理http请求的
// 首先有一个handler handler用来接收http请求
// 每进来一个请求，go语言都会给它创建一个goroutine,对请求的处理 就是由goroutine来完成
// 当我们使用ListenAndServe注册服务的时候，最后一个参数就指定使用的handler
// 当我们传入的参数为nil的时候，默认使用的是http.DefaultServeMux
func main() {
	http.HandleFunc("/", handlerBasic)
	//启动服务
	//启动服务时 使用nil，代表使用http 默认的多路复用器
	//defaultServerMutx
	// http.ListenAndServe第一个参数 网络地址 默认是80端口
	// 第二个参数指定handler
	http.ListenAndServe("172.26.215.133:9090", nil)

	//使用第二种方式注册服务
	startServer()
}

func handlerBasic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// 创建web server 还可以使用http.Server这是一个struct
// 其中有一些字段
// Addr 字段表示网络地址，默认使用80端口
// Handler 字段设置为nil 使用默认的多路复用器 DefaultServeMux
// ListenAndServe() 函数
// 我们调用http.ListenAndServe() 函数也是调用的http.Server的ListenAndServe()函数

func startServer() {
	server := http.Server{
		Addr:    "172.26.215.133:9090",
		Handler: nil,
	}
	server.ListenAndServe()
}

//handler 是一个接口
// handler定义了一个方法 ServeHTTP()
// ServeHTTP()有两个参数
// 第一个参数:HTTPResponseWriter 指的是响应
// 第二个参数 指向Request这个struct的指针
// 任何东西 只要含有这个方法 就实现了Handler这个接口

// type Handler interface{
// 	ServerHTTP(ResponseWriter,*Request)
// }

// handler为nil的时候 默认使用DefaultServeMux
// 它也是一个handler
// 请求进来的时候，它接收请求，由它将请求分发给某一个Handler
// 所以默认的多路复用器也是Handler

// 自己实现一个Handler

// type myHandler struct{}

// func (m *myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Hello"))
// }

// 我们自己实现的handler有一个问题
// 当一个请求过来，我们直接将其返回了，这样不管访问什么得到的结果都是一样的
// 而实际上，一个多路复用器需要能够允许请求注册进来，然后对请求进行分发
