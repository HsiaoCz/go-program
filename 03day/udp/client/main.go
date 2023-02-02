package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//upd client

func main() {
	//1.建立连接
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(172, 28, 97, 151),
		Port: 9090,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//关闭连接
	defer conn.Close()
	var reply [1024]byte
	//回复数据
	//这里使用bufio从控制台读取消息
	//将消息回复
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入内容:")
		msg, _ := reader.ReadString('\n')
		conn.Write([]byte(msg))
		//回复数据
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("收到回复信息:", string(reply[:n]))
	}

}
