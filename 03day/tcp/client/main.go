package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//tcp 客户端

func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "172.28.101.50:9090")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	//2.发送数据
	// s := "Hello Server"
	// _, err = conn.Write([]byte(s))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//可以简单聊天的tcp
	var s string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入聊天信息:")
		s, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if s == "exit" {
			break
		}
		_, err = conn.Write([]byte(s))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
