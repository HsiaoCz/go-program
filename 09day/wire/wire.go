package main

// func InitMission(name string) Mission {
// 	wire.Build(NewMonster, NewPlayer, NewMission)
// 	return Mission{}
// }

// 实际上 wire 的要求很简单
// 需要一个初始化函数，我们需要创建一个Mission对象
// 函数的返回值就是我们创建的对象类型,wire只需要知道类型,return后返回什么不重要
// 在函数中，我们调用wire.build()将创建mission所依赖的类型的构造器传进去，
// 例如 需要调用NewMission()创建Mission类型，NewMission()接受两个参数一个Monster类型，一个Player类型,Monster类型对象需要调用NewMonster()创建,Player类型对象调用NewPlayer()创建，所以NewMonster()和NewPlayer()我们也需要传递给wire
