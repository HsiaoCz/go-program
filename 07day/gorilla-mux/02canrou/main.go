package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 带参数的路由 可以是普通的路由，也可以是正则匹配的路由

func main() {
	r := mux.NewRouter()
	// 普通参数路由
	r.HandleFunc("/articles/{title}", TitleHandler)
	// 正则路由参数
	r.HandleFunc("/articles/{title:[a-z]+}", TitleHandler)
	// 一个受限的路由
	r.HandleFunc("/hello", HelloHandler).Host("www.baidu.com").Methods("GET").Schemes("http")

	// mux可以设置子路由 也就是路由组
	s := r.Host("http://127.0.0.1:9090/books/").Subrouter()
	s.HandleFunc("/info", BookInfo)
	// 当多个路由都注册进路由器，并且都匹配了
	// 会使用第一个路由

	// 给路由添加前缀
	// book := r.PathPrefix("/books").HandlerFunc(BookInfo)

	// 分组路由
	books := r.PathPrefix("/books").Subrouter()
	books.HandleFunc("/info", BookInfo)
	books.HandleFunc("/author", BookAuth)

	// 使用中间件
	// mux中间件的使用
	// 中间件的定义
	// type middlewareFunc func(http.Handler)http.Handler
	// 使用中间件
	r.Use(LoggerMid)
	http.ListenAndServe(":9000", r)

}

func TitleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //mux.Vars(r) 获取路由参数
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, vars["title"])
}

// mux可以限制路由
// r := mux.NewRouter()
// 只匹配 www.example.com
// r.Host("www.example.com")
//  动态匹配子路由
// r.Host("{subdomain:[a-z]+}.example.com")

// r := mux.NewRouter()

// r.PathPrefix("/products/")    //前缀匹配
// r.Methods("GET", "POST")      //请求方法匹配
// r.Schemes("https")            //schemes
// r.Headers("X-Requested-With", "XMLHttpRequest")  //header 匹配
// r.Queries("key", "value")  //query的值匹配

// 用户自定义方法 匹配
// r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
//     return r.ProtoMajor == 0
// })

func HelloHandler(w http.ResponseWriter, r *http.Request) {}
func BookInfo(w http.ResponseWriter, r *http.Request)     {}
func BookAuth(w http.ResponseWriter, r *http.Request)     {}

// 中间件有点特殊
func LoggerMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Fprintln(w, r.URL)
		next.ServeHTTP(w, r)
	})
}
