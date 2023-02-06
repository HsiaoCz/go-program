package calc

import "net/rpc"

// 定义服务端和客户端通用的rpc接口

// serviceName 计算器服务名
const ServiceName = "CalcService"

// ServiceInterface 计算器服务接口
type ServiceInterface interface {
	// CalcTwoNumber 对两个数进行运算
	CalcTwoNumber(request Calc, reply *float64) error
	// GetOperators 获取所有支持的运算
	GetOperators(request struct{}, reply *[]string) error
}

func RegisterCalcService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

type Calc struct {
	Number1  float64
	Number2  float64
	Operator string
}
