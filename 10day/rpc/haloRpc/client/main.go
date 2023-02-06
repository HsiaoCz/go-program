package main

import (
	"log"
	"net/rpc"
)

func main() {
	//建立连接 这里使用tcp连接
	// client, err := rpc.Dial("tcp", ":3001")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 建立连接 使用http连接
	client, err := rpc.DialHTTP("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}
	var response string
	// 使用client.Call()方法进行调用的时候 传入的服务名称应该是结构体名称.方法名
	err = client.Call("HelloServer.Hello", "bob", &response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
