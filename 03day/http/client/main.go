package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// http 客户端
// 客户端向服务端发送请求 也可以使用浏览器和postman发送请求
func HelloGet() {
	resp, err := http.Get("http://172.28.106.164:9090/hello")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// 返回的是字节数据的切片
	reply, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(reply))
	//获取请求头的信息
	for k, v := range resp.Header {
		fmt.Printf("%s,%v\n", k, v)
	}
	//获取请求的传输协议
	fmt.Println(resp.Proto)
	//获取请求的状态
	fmt.Println(resp.Status)
}

// 客户端发送POST请求
func HelloPost() {
	reader := strings.NewReader("Hello,hi")
	http.Post("http://172.28.106.164:9090/hello/po", "text/plain", reader)
}

// 复杂的请求
func complexRequest() {
	reader := strings.NewReader("Hello server")
	if req, err := http.NewRequest("POST", "http://172.28.106.164:9090/", reader); err != nil {
		log.Fatal(err)
	} else {
		// 自定义请求头
		req.Header.Add("User-Agent", "中国")
		req.Header.Add("MyHeaderKey", "myHeaderValue")
		// 自定义cookie
		req.AddCookie(&http.Cookie{
			Name:    "auth",
			Value:   "passwrod",
			Domain:  "localhsot",
			Path:    "/",
			Expires: time.Now().Add(time.Duration(time.Hour * 24 * 365)),
		})
		client := http.Client{
			Timeout: 100 * time.Millisecond,
		}

		// 提交请求 Do
		if resp, err := client.Do(req); err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
		}
	}

}

func main() {
	HelloGet()
	HelloPost()
	complexRequest()
}
