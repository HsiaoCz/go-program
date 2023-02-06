package main

import "errors"

// 抽象的计算函数类型

// Operation 是计算的抽象
type Operation func(Number1, Number2 float64) float64

// Add 是加法的operation实现
func Add(Number1, Number2 float64) float64 {
	return Number1 + Number2
}

// Sub 是Operation的减法实现
func Sub(Number1, Number2 float64) float64 {
	return Number1 - Number2
}

// Mul 是乘法的Operation 实现
func Mul(Number1, Number2 float64) float64 {
	return Number1 * Number2
}

// div是除法的Operation实现
func Div(Number1, Number2 float64) float64 {
	return Number1 / Number2
}

// 工厂
// Operators 注册所有支持的运算
var Operatiors = map[string]Operation{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

// CreateOperation 通过string表示operator获取适合的Operation函数
func CreateOperation(operator string) (Operation, error) {
	var oper Operation
	if oper, ok := Operatiors[operator]; ok {
		return oper, nil
	}
	return oper, errors.New("illegal Operator")
}
