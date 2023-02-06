rpc:远程过程调用

区别于本地过程调用

```go
package main

import "fmt"

func main(){
// 调用本地函数
a:=10
b:=20
ret:=add(a,b)
fmt.Println(ret)
}

func add(s,a int)int{
    return a+b
}

```

将这个程序编译之后运行，会输出结果 30
我们在 main 函数里调用 add 函数，有这样几个步骤 1.将变量 a 和 b 的值分别压入堆栈上 2.执行 add 函数，从堆栈中获取到 a 和 b 的值，并将它们分配给 x 和 y 3.计算 x+y 的值并且将其值保存到堆栈中 4.退出 add 函数并将 x+y 的值赋值给 ret

rpc 调用：
本地过程调用发生在同一个进程中，定义函数的代码和调用函数的代码是共享一个内存空间的，所以调用能够正常执行

但是我们无法直接调用在另一个服务器上的 add 函数，它们在两个程序中，内存空间物理隔离
rpc：远程过程调用就是为了解决这类问题
要实现 rpc 需要实现以下三个问题: 1.如何确定要执行的函数?
在本地调用的过程中，函数主体是通过函数指针函数指定，调用 add 函数，编译器通过函数指针函数自动确定 add 函数在内存中的位置。
在 rpc 中没办法通过函数指针函数确定函数在内存中的位置，所以调用双方需要维护一个{function--->ID}的映射表，用来确定调用正确的函数

2.如何表达参数？
本地过程调用的参数传递是通过内存堆栈完成的，远程调用传递参数和返回参数需要在传递期间序列化并转换成字节流

3.如何实现网络传输
函数调用方和被调用方通常是通过网络连接的

go rpc 由 net/rpc 包提供
它使用的是 go 特有的编码方式 gob,所以 go 的客户端和服务端都必须使用 go 语言编写

对于服务端,net/rpc 要求用一个导出的结构体来表示 rpc 服务，这个结构体中含有特定要求的方法

```go
type T struct{}

func (t *T)MethodName(argType T1,replyType *T2)error
```

这里要求结构体，方法都是可导出的，也就是首字母要求大写
方法的第二个参数是指针，方法的返回值是 error

服务端通过 rpc.Dial(对 tcp 服务)连接服务端，然后使用 call 调用 rpc 服务中的方法:
`rpc.Call("T.MethodName",argType T1,replyType \*T2)

net/rpc 允许使用 json 格式

```go
// 服务段的编码
rpc.ServeCodec(SomeServerCodec(conn)) // SomeServerCodec 是个编码器

// 客户端的解码
conn, _ := net.Dial("tcp", "localhost:1234")
client := rpc.NewClientWithCodec(SomeClientCodec(conn)) // SomeClientCodec 是个解码器
```

net/rpc/jsonrpc 就是这样的一种实现，它使用 json 而不是 Gob 编码，可以用来跨语言 Rpc
net/rpc/jsonrpc 提供了很 net/rpc 大致上相同的 api

```go
//使用net/rpc/jsonrpc和之前的基本上是一样的
// 服务端的main需要改变一下
go jsonrpc.ServeConn(conn)

// 实际上jsonrpc.ServerConn的实现是rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
// 在调用的时候，将DialHelloService中连接服务的代码改一改就可以使用了
c,err:=jsonrpc.Dial(network,address)

// 或者这样
conn, _ := net.Dial("tcp", "localhost:1234")
client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
```
