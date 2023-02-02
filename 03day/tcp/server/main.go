package main

import (
	"fmt"
	"log"
	"net"
)

//一个tcp服务端可以同时连接多个客户端
//tcp 连接流程
//1.监听端口
//2.接收客户端的连接请求
//3.创建goroutine 处理连接

func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "172.28.101.50:9090")
	if err != nil {
		log.Fatalln(err)
	}
	//等待别人连接
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	//3、与客户端建立通信
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(tmp[:n]))
	}
}
