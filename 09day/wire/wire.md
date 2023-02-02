# wire

wire 谷歌开源的依赖注入工具，是一个代码生成器，并不是框架
我们只需要在一个特殊的go文件中告诉wire类型之间的依赖关系，它会自动帮我们生成代码，帮助我们创建指定类型的对象，并组装它的依赖

安装工具:

```go
go get github.com/google/wire/cmd/wire
```

```go

func InitMission(name string) Mission {
 wire.Build(NewMonster, NewPlayer, NewMission)
 return Mission{}
}
```

实际上 wire 的要求很简单
需要一个初始化函数，我们需要创建一个Mission对象
函数的返回值就是我们创建的对象类型,wire只需要知道类型,return后返回什么不重要
在函数中，我们调用wire.build()将创建mission所依赖的类型的构造器传进去，
例如 需要调用NewMission()创建Mission类型，NewMission()接受两个参数一个Monster类型，一个Player类型,Monster类型对象需要调用NewMonster()创建,Player类型对象调用NewPlayer()创建，所以NewMonster()和NewPlayer()我们也需要传递给wire

wire的基础概念:
Prodiver(构造器)
Injector(注入器)
Provider实际上就是创建函数,我们上面InitMission就是Injector
每个注入器实际就是一个对象的创建和初始化函数

参数:编写InitMission()函数带有一个string类型的参数
并且在生成的initMission()函数中，这个参数传递给了NewPlayer().
NewPlayer()需要一个string类型的参数

如果NewMonster()也需要一个string类型的参数
生成的代码中NewMonster()也会传递这个参数
