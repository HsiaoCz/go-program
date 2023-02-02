package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// http 服务端
// 请求处理函数
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// post请求的处理函数
func HelloPo(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(reader))
}

// 拿到请求参数
func HelloGet(w http.ResponseWriter, r *http.Request) {

}

func HelloBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
func main() {
	//注册请求路由和处理函数
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/hello/po", HelloPo)
	http.HandleFunc("/hello/:name", HelloGet)
	http.HandleFunc("/hello/book", HelloBook)
	//监听端口
	http.ListenAndServe("172.28.106.164:9090", nil)
}
