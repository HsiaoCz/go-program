package main

import (
	"flag"
	"fmt"
	"net/http"
)

//http 请求 HTTP Request 和 HTTP Response 请求和响应
// 它们具有相同的结构
// 请求(响应)行
// 0个或多个Header
// 空行
// 可选的消息体

// http 请求的例子
// GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
// Host:www.w3.org
// User-Agent:Mozilla/5.0
// 空行

// go 语言中的 net/http 包提供了用于表示HTTP消息的结构
// 表示请求的 Request
// Request 是一个struct，代表了客户端发送的HTTP请求消息
// 重要的字段：
// URL
// Header
// Body
// Form、POSTForm、MultipartForm
// 通过Request的方法访问请求中的Cookie、URL、User Agent等信息
// Request既可以代表发送到服务器的请求，又可以代表客户端发出的请求

// 请求的URL
// Request的URL字段就代表了请求行(请求信息第一行)里面的部分内容
// URL字段指向url.URL类型的一个指针，url.URL是一个struct:
// type URL struct{
//   Scheme string
//   Opaque string
//   User *UserInfo
//   Host string
//   Path string
//   RawQuery string
//   Fragment string
// }
//  URL的通用格式：scheme://[userinfo@]host/path[?query][#fragment]
// url query RawQuery 会提供实际查询的字符串
// 例如 http://www.example.com/post?id=123&thread_id=456
// 这个URL的 rawQuery的值就是id=123&thread_id=456
// 还有一个简便的方法可以得到key-value对:通过Request的Form字段
// URL Fragment 如果从浏览器发出的请求，那么你无法提取出Fragment字段的值
// 浏览器在发送请求 时会把fragment部分去掉
// 但不是所有的请求都是从浏览器发出的

// 请求和响应(request/response)的headers是通过Header类型来描述的,它是一个map,
// 用来表示HTTP Header里的Key-Value对
// Header map的key是string类型,value是[]string
// 设置key的时候会创建一个空的[]string作为value,value里面第一个元素就是新的header的值
// 为指定的key添加一个新的header值，执行append操作即可
// 想要获得header 使用r.Header 返回map r.Header["Accept-Encoding"]
// 返回:[gzip,deflate]([]string类型)
// r.Header.Get("Accept-Encoding")
// 返回：gzip,deflate(string类型)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "set http server addr")
	flag.Parse()
	http.HandleFunc("/header", ReturnHeader)
	http.HandleFunc("/posthan",PostHandle)
	server := http.Server{
		Addr:    *listenAddr,
		Handler: nil,
	}
	server.ListenAndServe()
}

func ReturnHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Header["Accept-Encoding"])
	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
}

// Request Body 请求和响应的消息体都是body字段来表示的
// Body是一个io.ReadCloser接口
// 这个接口包含了两个接口
// 一个是Reader接口 一个是Closer接口
// Reader接口定义了一个Open方法 这个方法有一个[]byte参数，返回byte的数量和错误
// Closer接口定义了一个Close方法
// 这个方法没有参数返回的是一个可选的错误信息
// 想要读取请求body的内容，可以调用Body的Read方法

func PostHandle(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
