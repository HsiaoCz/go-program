gRPC 是一种现代化开源的高性能 RPC 框架
使用 HTTP/2 作为传输协议

在 gRPC 里，客户端可以像调用本地方法一样直接调用其他机器上的服务端应用程序的方法，帮助你更容易创建分布式应用程序和服务。与许多 RPC 系统一样，gRPC 是基于定义一个服务，指定一个可以远程调用的带有参数和返回类型的的方法。在服务端程序中实现这个接口并且运行 gRPC 服务处理客户端调用。在客户端，有一个 stub 提供和服务端相同的方法。

使用 grpc 需要先安装几个工具:

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

grpc 的开发方式
一共分为三步：
**1.编写.proto 文件定义服务**

```protobuf
syntax="proto3";

package hello;

option go_package="./;hello";

service  HelloService{
    rpc SayHello(HelloRequest) returns (HelloResponse);
}

message HelloRequest{
   string greeting=1;
}

message HelloResponse{
   string reply=1;
}


```

在 grpc 中可以定义四种类型服务的方法 1.普通 rpc,客户端向服务器发送一个请求，然后得到一个响应，就像普通函数调用一样

```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse);
```

2.服务器流式 rpc，其中客户端向服务器发送请求，获得一个流来读取一系列消息，客户端从返回的流中读取，直到没有更多的消息。grpc 保证在单个 rpc 调用中的消息是有序的

客户端发出一个RPC请求，服务端与客户端之间建立一个单向的流，服务端可以向流中写入多个响应消息，最后主动关闭流；而客户端需要监听这个流，不断获取响应直到流关闭。应用场景举例：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。

```protobuf
rpc LostOfReplies(HelloRequest) returns (stream HelloResponse);
```

3.客户端流式 rpc，其中客户端写入一系列消息将其发送到服务器，同样使用提供的流。一旦客户端完成了消息的写入，它就等待服务器读取消息并返回响应，同样，grpc 保证在单个 rpc 调用中对消息进行排序

客户端传入多个请求对象，服务端返回一个响应结果。典型的应用场景举例：物联网终端向服务器上报数据、大数据流式计算等。

```protobuf
rpc LostOfGreeetings(stream HelloRequest) returns (HelloResponse)
```

4.双向流式 rpc,其中双方使用读写流发送一系列消息，两个流独立运行

双向流式RPC即客户端和服务端均为流式的RPC，能发送多个请求对象也能接收到多个响应对象。典型应用示例：聊天应用等。

```protobuf
rpc BindiHello(stream HelloRequest)returns (stream HelloRequest);
```

**2.生成指定语言的代码**
通常在客户端调用 api，在服务端实现相应的 api

在服务器端，服务器实现服务声明的方法，并运行一个 gRPC 服务器来处理客户端发来的调用请求。gRPC 底层会对传入的请求进行解码，执行被调用的服务方法，并对服务响应进行编码。
在客户端，客户端有一个称为存根（stub）的本地对象，它实现了与服务相同的方法。然后，客户端可以在本地对象上调用这些方法，将调用的参数包装在适当的 protocol buffers 消息类型中—— gRPC 在向服务器发送请求并返回服务器的 protocol buffers 响应之后进行处理。

**3.编写业务逻辑代码**

编写完proto文件
在项目的根目录执行以下代码:
```go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto //这里是proto文件所在的位置
```
--go_out=[这里是文件生成的位置]

