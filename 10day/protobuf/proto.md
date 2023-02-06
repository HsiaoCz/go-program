ProtoBuf
是一种数据交换协议
protobuf 全称：协议缓冲区，是一种与语言无关的，与平台无关的可扩展机制，用于序列化结构数据
json/xml 都是基于文本格式，protobuf 是二进制格式

编写 proto 代码

```protobuf
syntax="proto3";// 声明使用的版本号

package hello;// 定义包名

option go_package="./;hello";// 定义go的包名，用于生成.pd.go文件
// 定义消息体
message Say{
    int64 id =1;
    string hello=2;
    repeated string word=3
}
```

生成 go 代码

```go
protoc --go_out=. hello.proto
```

关键字：
syntax：必须，定义在第一行，不写默认使用 proto2
package；定义 proto 文件的包名
option go_package:定义生成的 pb.go 的包名
message:用于定义消息体

repeated :用于声明数组

message:消息，用来在 protobuf 中指定我们要定义的数据结构
