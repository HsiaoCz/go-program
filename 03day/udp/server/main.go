package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// udp 是一种无连接的传输方式,不可靠，但是实时性比较好
// udp 一般用在直播等方面
func main() {
	//1.连接
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(172, 28, 97, 151),
		Port: 9090,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//不需要建立连接，直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(data[:])
		reply := strings.ToUpper(string(data[:n]))
		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}
