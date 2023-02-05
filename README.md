# go-program

go程序设计语言，一些重点，偏基础

## 1.声明

声明这有个值得注意的点
短变量声明只能在函数内，并且意味着声明并赋值

变量的一些声明方式

## 2、变量的声明方式

这主要注意短变量声明方式的一些特点

## 3、指针

& 取地址 * 根据地址取值
指针的零值是nil
判断p!=nil 如果结果为true说明指针指向的有值

另外指针是可以比较的，当且仅当两个指针指向同一个变量或者两个指针都是nil的时候 指针相等

函数传参传递的是值得拷贝，当传入指针的时候，改变指针的值可以改变变量的值
当我们希望在函数内修改变量的值的时候，可以传递指针

## 4、指针接收者和值接收者的区别

这里只需要记住一个根本区别
那就是值接收者的方法和指针接收者的方法的方法名是不一样的

## 5、时间函数

这里主要注意一下time.timer,time.ticke
时间格式化，2006-01-02 15:04:05

## 6、sync.Once

只需要执行一次的操作 可以使用sync.once
保证只执行一次

// 定义
var once sync.Once
// 使用
once.Do(函数)
// once.Do() 这个函数传进去的只能是没有参数和返回值的函数
// 所有要执行一个有参数或者返回值的函数的时候，可以使用闭包

## 7、sync.Map

go原生的map不是并发安全的，想要安全的并发Map可以使用sync.Map
sync.Map 主要是注意它的设置值和取值都都有自己的方法

## 8、tcp/udp

tcp的粘包问题
tcp 之所以会出现粘包，是因为tcp数据传输是流模式，在保持长连接的时候可以进行多次的收和发
粘包既可能发生在接收端，也可能发生在发送端
Nagle算法造成的粘包，Nagle算法是一种改善网络传输效率的算法，简单来讲就是当我们提交一段数据给TCP发送时，TCP并不立即发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则一次把这两段数据发送出去
接收端接收不及时造成接收端粘包，TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据，当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放几段数据

粘包的解决办法：
出现粘包的关键在于接收方不确定要传输的数据包的大小，因此我们可以对数据包进行封包和拆包操作。

## 9、Makefile

makefile有两个概念：
目标、依赖
目标就是指要干什么，或者说运行make后生成什么，依赖告诉目标如何实现目标
在makefile中，目标和依赖是通过规则来表达的

目标是冒号前面的，冒号下面的是生成命令，整一套称为规则
makefile可以定义多个目标，当没有指定哪个目标时，默认使用第一个目标

和目标在一行 冒号后面的就是依赖
依赖是先决条件，构建的时候先从依赖开始构建

规则，一个规则由目标，依赖和，命令组成
make在运行一个规则时，会先检查依赖中相关文件的时间戳，如果依赖文件的时间戳大于目标文件的时间戳，那说明依赖文件比较新，会先执行依赖

假目标，由于make是看时间戳执行，所以可能出现同名文件执行错误的情况
这时候可以使用假目标

假目标使用:.PHONY来定义

makefile中的变量：
变量的定义，一个变量名，后面跟一个等号，等号后面是变量的期望值
变量引用，使用$(变量名)或者${变量名}这种方式

```makefile
a=main.go
run:
    @go run $(a)
```

## 10、todo+的用法

ctrl+shift+p 搜索打开todo
ctrl+enter 添加代办事项 再按取消
alt + 5 开始代办事项 再按撤销
alt + d 完成事项
alt + c 撤销事项
在todo下右键点击achive归档事件

## 11、参数校验库validator

这个库主要是用来做参数校验的
1.标记符号的说明
`,`:可以把多个验证隔开，但是之间不能有空格
`-`:跳过验证该字段
`|`:使用多个验证标记但是只需要满足其中的一个即可
required:表示字段必须设置，不能为默认值
omitempty:如果字段未设置，忽略它

2.范围比较验证
len:等于参数值
max:最大值，小于等于参数值 validate:"max=20"
min:最小值，大于等于参数值 validate:"min=0"
ne:不等于
oneof:只能是列举中的值的一个，以空格分隔，列举的字符串里有空格，字符串用单引号引起来

3.字符串验证
contains:包含参数子串,validate:"contains=tom"
excludes:不包含参数子串,validate:"excludes=tom"
startswith:以参数子串为前缀
endwith:以参数子串为后缀

4.字段验证
eqcsfield:跨不同结构体字段验证，它的意思是这个结构体的这个字段值要等于另一个结构体的字段值 validate:"eqcsfield=Struct2.Field2"
neqcsfield:跨不同结构体字段的值不相等
eqfield:同一结构体字段验证相等，最常见的用法是输入两次密码验证

```go
type User struct{
    Name string `validate:"lte=4"`
    Age int `validate:"min=20"`
    Password string `validate:"min=10"`
    Password2 string `validate:"eqfield=Password"
}
```

nefield:同一结构体字段验证不相等
gtefield:大于等于同一结构体字段
ltefiled:小于等于同一结构体字段

5.网络验证
ip:字段值是否包含有效的ip地址
ipv4:
ipv6:
uri:
url:

```go
validate.New() //可以注册一个验证器
validate.Var() //字符串
validate.Struct() //验证结构体
```

## 12、Mysql

### 12.1、数据库语言

DDL:数据定义语言
DML:数据操作语言
DQL:数据查询语言
DCL:数据控制语言

### 12.2、操作数据库

```sql
// 创建数据库
create database if not exists student;
// 删除数据库
drop database if exists student;
// 使用数据库
use student;
// 查看数据库
show databses;
```

### 12.3、列的数据类型

> 数值
tinyint  十分小的数据  1个字节
smallint 较小的数据   2个字节
mediuint 中等大小的数据 3个字节
int    4个字节
bigint 8个字节
float  4个字节
double 8个字节
decimal 字符串形式的浮点数
>字符串
char  字符串固定大小  0~255
varchar  可变字符串 0~65535
tinytext  微型文本 2^8-1
text  文本类型 2^16-1
>时间日期
date YYYY-MM-DD  日期格式
time HH:mm:ss 时间格式
datetime  YYYY-MM-DD HH:mm:ss 日期格式
timestamp 时间戳，从1970.1.1到现在的秒数
>null 没有值，未知，不要使用null进行运算，结果还是为null

### 12.4、数据可的字段属性

Unsigned:无符号的整数，声明了该列不能为负数
zerofill:0填充的，不足的位数用0来填充
AUTO_INCREMENT:自增，用来设置主键，必须是整数类型
null/not null:空和非空 非空不赋值会报错，空不赋值就为空
default:设置默认值

```sql
//这几个字段用来表示一个记录存在的意义
id          主键
version     乐观锁
is_delete   伪删除
gmt_create  创建时间
gmt_update  修改时间
```

### 12.5、创建数据库

```sql
create table if not exists `sutdent`(
  `id` INT(4) NOT NULL AUTO_INCREMENT COMMENT '学号',
  `name` VARCHAR(30) NOT NULL DEFAULT '匿名' COMMENT '姓名',
  `sex` VARCHAR(2) NOT NULL DEFAULT '女' COMMENT '性别',
  `birthday` DATETIME NOT NULL COMMENT '出生日期',
  `address` VARCHAR(100) DEFAULT NULL COMMENT '家庭地址',
  `email` VARCHAR(30) DEFAULT NULL COMMENT '邮箱',
  PRIMARY KEY(`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8
```

查看表结构：`desc sudent`,查看表的描述信息

### 12.6、修改和删除数据表字段

修改:

```sql
-- 修改表名
ALTER TABLE `teacher` RENAME AS `teacher1`;
-- 新增表字段
ALTER TABLE `teacher` ADD age INT(3);
-- 修改约束
ALTER TABLE `teacher` MODIFY age varchar(2);
-- 给字段重命名
ALTER TABLE `teacher` change age age1 INT(3);
-- 删除表的字段
ALTER TABLE `teacher` drop age1;
```

删除表:
`drop table if exists teacher;`

mysql 8.0版本的创建用户:

create user 'shaw'@'%' identified by 'hsiaocz123';
GRANT ALL PRIVILEGES ON user.* TO 'shaw'@'%' WITH GRANT OPTION;
flush privileges;
