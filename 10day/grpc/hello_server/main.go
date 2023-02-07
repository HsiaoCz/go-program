package main

import (
	"context"
	"go-program/10day/grpc/hello_server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// hello server

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{Reply: "Hello" + in.Name}, nil
}

func main() {
	//监听本地端口
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()                  //创建grpc服务器
	pb.RegisterGreeterServer(s, &server{}) //在gRpc服务端注册服务

	//启动服务
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
