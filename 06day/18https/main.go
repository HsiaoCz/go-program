package main

import (
	"fmt"
	"net/http"
)

// http传输是明文的，非常不安全
// https 安全的传输协议 TLS
// 之前使用的是http.ListenAndServe()函数，只能使用http
// 使用https,需要使用TLS
// 这里需要生成两个tls文件
func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServeTLS(":8000", "", "", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
