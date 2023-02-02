package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// 中间件：请求先经过中间件的处理，再到请求的Handler,Handler处理之后
// 再经中间件，中间件将响应返回

// 设计中间件
// 中间件里面有一个handler
// 中间件本身也是handler
// 中间件最终会调用next这个handler
// 但是在调用之前，我们可以先做一些别的处理
// 然后调用m.Next.ServeHTTP(w,r)
// 调用之后，也可以做一些其他的处理
type MiddlewareBook struct {
	next http.Handler
}

// 创建中间件逻辑
func (m MiddlewareBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.next == nil {
		m.next = http.DefaultServeMux
	}
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	m.next.ServeHTTP(w, r)
}
func main() {
	http.HandleFunc("/books", BookHandler)
	// 使用中间件
	http.ListenAndServe(":9090", new(MiddlewareBook))
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	b := Book{
		ID:      rand.Intn(10000),
		Name:    "hello",
		Country: "USA",
	}
	// NewEncoder()的参数是转换到哪
	enc := json.NewEncoder(w)
	enc.Encode(b)
}

type Book struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
