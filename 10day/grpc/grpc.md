gRPC是一种现代化开源的高性能RPC框架
使用HTTP/2作为传输协议

在gRPC里，客户端可以像调用本地方法一样直接调用其他机器上的服务端应用程序的方法，帮助你更容易创建分布式应用程序和服务。与许多RPC系统一样，gRPC是基于定义一个服务，指定一个可以远程调用的带有参数和返回类型的的方法。在服务端程序中实现这个接口并且运行gRPC服务处理客户端调用。在客户端，有一个stub提供和服务端相同的方法。

使用grpc需要先安装几个工具:
```go
// protoc的go插件 这个插件会生成一个后缀为.pb.go文件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
// 安装gRPC插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
// 这个插件会生成一个后缀为_grpc.pb.go文件
// 其中包含：
// 一种接口类型，供客户端调用的服务方法
// 服务端要实现的接口类型
```

grpc的开发方式
一共分为三步：
1.编写.proto文件定义服务

在grpc中可以定义四种类型服务的方法
1.普通rpc,客户端向服务器发送一个请求，然后得到一个响应，就像普通函数调用一样
```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse)
```

2.服务器流式rpc，其中客户端向服务器发送请求，获得一个流来读取一系列消息，客户端从返回的流中读取，直到没有更多的消息