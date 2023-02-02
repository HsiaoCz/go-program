package main

import "net/http"

// go语言内置了5个handler
// http.NotFoundHandler
// func NOtFoundHandler() Handler
// 返回一个handler,它给每个请求的响应都是404 page not found

//2、http.RedirectHandler
// func RedirectHandler(url string,code int)Handler
// 返回一个handler 它将每个请求使用给定的状态码跳转到指定的URL
// 参数：
// url:要跳转的URL
// code : 跳转的状态码(3xx),常见的:StatusMovedPermanently
// StatusFound 或者 StatusSeeOther等

// 3、http.StripPrefix
// func StripPrefix 返回一个Handler,它从请求的URL中去掉指定的前缀，然后调用另一个handler
// 如果请求的URL与提供的前缀不符，那么404
// 有点像中间件
// prefix,URL将要被移除的字符串前缀
// h,是一个handler，在移除字符串前缀之后，这个handler将会接收到请求
// 它实际上是修饰了另外一个handler

// http.TimeoutHandler
// func TimeoutHandler(h Handler,dt time.Duration,msg string)Handler
// 返回一个handler，它用来在指定时间内运行传入的h
// 也相当于一个修饰器
// h 将要被修饰的handler
// dt 第一个handeler允许的处理时间
// msg 如果超时，那么就把msg返回给请求，表示响应时间过长

// http.FileServer
// func FileServer(root FIleSystem)Handler
// 返回一个handler，使用基于root的文件系统来响应请求
// type FileSystem interface{ Open(name string)(File,error)}
// 使用时需要用到操作系统的文件系统，所以还需要委托给:
// type Dir string
// func (d Dir)Open(name string)(File,error)
func main() {
	// http.HandleFunc("/", viweFile)
	http.ListenAndServe(":9090", http.FileServer(http.Dir("wwwroot/static")))
}

// func viweFile(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "wwwroot/static"+r.URL.Path)
// }

// 使用http.FileServer也可以实现
