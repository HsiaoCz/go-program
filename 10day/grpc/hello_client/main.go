package main

import (
	"context"
	"flag"
	"go-program/10day/grpc/hello_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// hello_client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", ":3001", "the addr to connect")
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()

	//连接到server端,此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	//执行rpc调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("cound not greed:%v", err)
	}
	log.Printf("Greeting:%s", r.GetReply())
}
