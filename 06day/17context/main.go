package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Request Context
// func(*Request)Context()context.Context
// 这个函数返回请求的上下文
// func(*Request)WithContext(ctx context.Context)context.Context
// 这个函数可以基于Context进行"修改",实际上创建一个新的Context
func main() {
	http.HandleFunc("/books", BookHandler)
	// 使用中间件
	http.ListenAndServe(":9090", &TimeoutMiddleware{next: new(MiddlewareBook)})
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
	}
	m.next.ServeHTTP(w, r)
}

type TimeoutMiddleware struct {
	next http.Handler
}

func (t TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.next == nil {
		t.next = http.DefaultServeMux
	}
	ctx := r.Context()
	ctx, cc := context.WithTimeout(ctx, time.Second*3)
	r = r.WithContext(ctx)
	fmt.Println(cc)
	ch := make(chan struct{})

	go func() {
		t.next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
