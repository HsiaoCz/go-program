package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 使用gorilla/mux
// go自带的路由库比较简单，http.ServerMux，本质上是一个map[string]Handler
// 是请求的URL路径和对应的handler的映射
// 不支持参数设定，例如/usr/:id这种泛型匹配
// 对RESTful不友好，无法限制访问的方法
// 不支持正则表达式
// github.com/gorilla/mux

// 使用mux创建普通路由
func main() {
	r := mux.NewRouter()
	//普通路由
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", ProductsHandler)
	http.ListenAndServe(":9090", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello")
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello Products")
}
