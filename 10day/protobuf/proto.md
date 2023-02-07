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
    int64 id =1; // 数字表示分配表示号
    string hello=2; // 1-15标识号会占用一个字节，我们可以为那些频繁使用的分配1-15
    repeated string word=3
}
```

生成 go 代码

```go
protoc --go_out=. hello.proto
```

**关键字：**
syntax：必须，定义在第一行，不写默认使用 proto2
package；定义 proto 文件的包名
option go_package:定义生成的 pb.go 的包名
message:用于定义消息体

repeated :用于声明数组

message:消息，用来在 protobuf 中指定我们要定义的数据结构

Reserved :保留标识号，保留一些标识号为以后可能使用到的使用

```protobuf
message Test{
    reserved 2,3,5,7 to 10; // 保留2,3,5,7到10的标识号给以后用
}
```
如果使用了保留的标识号，编译会报错

**protobuf可以编译成各种语言的代码**

语法格式为：
```protobuf
protoc --option_out=[这里是保存的路径] [这里是proto文件名称]
```

**枚举关键字**

当定义一个消息类型的时候，可能想为一个字段指定预定义的值中的某一个值，这时候可以使用枚举

```protobuf
syntax="proto3"; //指定版本信息

enum SexType // 枚举消息类型，使用enum指定
{
    UNKONW=0;//proto3中首个成员必须是0,成员不应该有相同的值
    MALE=1; // 男
    FEMALE=2;//女
}

// 定义一个用户消息
message UserInfo{
    string name=1; //姓名字段
    SexType sex=2; //性别字段，使用SexType枚举类型
}

```

**消息嵌套**
开发go语言时经常嵌套使用结构体
protobuf中同样支持消息嵌套
1.使用其他消息体
```protobuf
message Article{
    string url=1;
    string title=2;
    repeated string tags=3;
}

// 定义ListArticle消息
message ListArticle{
    //引用上面的结构体
    repeated Article article=1;
}
```

还可以在消息里直接嵌套消息
```protobuf
message ListArticle{
    //嵌套定义消息
    message Article{
        string url=1;
        string title=2;
        repeated string tags=3;
    }
    //引用嵌套类型
    repeated Article articles=1;
}
```

import导入其他的proto文件定义的消息

先创建article.proto
```protobuf
syntax = "proto3";

package nesting;

option go_package = "./;article";

message Article {
  string          url   = 1;
  string          title = 2;
  repeated string tags  = 3; // 字符串数组类型
}
```
再创建list_article.proto
```protobuf
syntax = "proto3";
// 导入Article消息定义
import "article.proto";

package nesting;

option go_package = "./;article";

// 定义ListArticle消息
message ListArticle {
  // 使用导入的Result消息
  repeated Article articles = 1;
}
```

**map类型**
go的切片类型对应的是repeated
go的map类型对应的是map
`map<key_type,value_type>map_field=N;`
key_type可以是除了浮点数和字节之外的任意类型
枚举不能作为key_type
value_type可以是除了另一个map外的任意类型
Map不能使用repeated字段

```protobuf
syntax = "proto3";

package map;

option go_package = "./;score";

message Student{
  int64              id    = 1; //id
  string             name  = 2; //学生姓名
  map<string, int32> score = 3;  //学科 分数的map
}
```

protobuf语法的更多内容可以查看[https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh/]
关于one of等特殊情况可以看[https://www.liwenzhou.com/posts/Go/oneof-wrappers-field_mask/]