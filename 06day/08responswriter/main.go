package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

// 从服务器向客户端返回响应需要使用ResponseWriter
// ResponseWriter是一个接口,handler用它来返回响应
// 真正支撑ResponseWriter的幕后Struct是一个非导出的http.response
// 为什么Handler的ServeHTTP(w ResponseWriter,r *Request),只有一个指针类型，而w是按值传递的吗？
// 这个ResponseWriter代表一个指向http.response的指针
// 所有可以看作是指针传递
// http.response的指针接收者实现了ResponseWriter

// 如何写入到ResponseWriter
// Writer方法接收一个byte切片作为参数，然后把它写入到Http响应的body里面
// 如果在Write方法被调用时，header里面没有设定content type，那么数据的前512字节就会被用来检测content type
func main() {
	listenAddr := flag.String("listenAddr", ":8000", "set http server listenAddr")
	flag.Parse()

	//handlerFunc
	http.HandleFunc("/hello", WriteExpample)
	http.HandleFunc("/write-header", WriteHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json-example", jsonExample)
	server := http.Server{
		Addr:    *listenAddr,
		Handler: nil,
	}
	server.ListenAndServe()
}

func WriteExpample(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>uploadFile</title>
	</head>
	<body>
		<form action="http://172.28.30.66:9090/hello?hello=world&thread=123" method="post",enctype="multipart/form-data">
			<input type="text" name="first_name" />
			<input type="text" name="last_name" />
			<input type="file" name="uploaded" />
			<input type="submit"/>
		</form>
	</body>
	</html>`
	w.Write([]byte(str))
}

// WriteHeader方法
// 这个方法接收一个整数类型(HTTP 状态码)并把它作为HTTP响应的状态码返回
// 如果这个方法没有显示的调用，那么在第一次调用Write方法前，会隐式的调用WriteHeader(http.StatusOK)
// 所有 WriterHeader主要用来发送错误类的HTTP状态码
// 调用完WriteHeader方法之后，仍然可以写入到ResponseWriter，但无法再修改Header了

func WriteHeaderExample(w http.ResponseWriter, r *http.Request) {
	//这里调用WriteHeader(501)
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service, try next door")
}

// Header 方法返回headers的map，可以进行修改
// 修改后的headers将会体现在返回给客服端的HTTP响应里

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

// json 方法
func jsonExample(w http.ResponseWriter, r *http.Request) {
    // 这里设置header的Content-Type的类型为application/json
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "zhangsan",
		Threads: []string{"first", "second", "third"},
	}
	// 这里使用json.Marshal()方法将结构体序列化成json格式的字节切片
	// 将json格式的字节切片写入到响应体中
	json, _ := json.Marshal(post)
	w.Write(json)
}

type Post struct {
	User    string
	Threads []string
}
