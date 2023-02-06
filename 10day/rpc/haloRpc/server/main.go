package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type HelloServer struct {
}

func (s *HelloServer) Hello(request string, response *string) error {
	// 返回值是通过修改response的值来实现
	*response = "hello," + request
	return nil
}

func main() {
	// 将给客户端访问的名字和helloServer实例注册
	rpc.RegisterName("HelloServer", new(HelloServer))

	// 通过http服务
	rpc.HandleHTTP()
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 使用tcp 连接
	// listener, err := net.Listen("tcp", "3001")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// conn, err := listener.Accept() //acept()接受请求 不过这里只能接受一个请求
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rpc.ServeConn(conn)

	// 接受多个请求
	// for{
	// conn, err := listener.Accept() //acept()接受请求 不过这里只能接受一个请求
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rpc.ServeConn(conn)
	// }
}
