package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 上传文件
// multipart/form-data最常见的应用场景就是上传文件

func main() {
	listenAddr := flag.String("listenAddr", ":9090", "set http server listenAddr")
	flag.Parse()
	http.HandleFunc("/", RootHandle)
	http.HandleFunc("/hello", HelloHandle)
	// 使用FormFile的handler
	http.HandleFunc("/hello/hello", HelloHandWithFormFile)
	server := http.Server{
		Addr:    *listenAddr,
		Handler: nil,
	}
	server.ListenAndServe()
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	// 调用parseMultipartForm()方法，加载数据
	r.ParseMultipartForm(1024)
	// 访问请求的mutlipartForm里面有一个字段file 值是一个map
	// 里面的key对应的是html文件里的文件上传块的name
	// 这里允许上传多个文件，所有这里的map值为[]string
	// 这里使用[]string的第一个值表示读取第一个文件
	// 这里返回的是一个fileHeader的指针
	fileHeader := r.MultipartForm.File["uploaded"][0]
	// 使用fileHeader的Open方法 可以得到文件
	file, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
	}
	// 如果没有错误 使用io.ReadAll函数读取文件 读取成一个[]byte
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// 最后将读取的文件写入responsewriter 这个将文件内容做了类型转换
	fmt.Fprintln(w, string(data))
}

func RootHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "upload.html")
}

// 上传文件还有一个简单的方法:FormFile
// 无需使用ParseMultipartForm方法
// 返回指定key对应的第一个value
// 同时返回file和fileHeader，以及错误信息
// 如果只上传一个文件 那么这种方式会更快一些

func HelloHandWithFormFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("upload")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(data))
}

// 需要注意的点是，不是所有的POST请求都来自FOrm
// 客户端框架(例如Angular等)会以不同的方式对POST请求编码
// ParseForm这种方法 无法处理application/json这种格式

// 读取form的方法：
// MutlipartReader()
// func (r *Requset)MutlipartReader()(*multipart.Reader,error)
// 如果是multipart/form-data或mutlipart混合POST请求
// 那么会返回MultipartReader 返回一个MIME multipart reader
// 否则返回nil 和一个错误
// 可以使用该函数代替ParseMultipartForm来把请求的body作为stream进行处理
// 之前的方法 都是把表单作为一个对象，来一次性处理的
// ParseMultipartForm 这个方法 可以逐个检查来自表单的值，然后一次处理一个
