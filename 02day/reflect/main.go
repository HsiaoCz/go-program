package main

import (
	"fmt"
	"reflect"
)

type myInt int

// 反射
// reflect.TypeOf,获取对象的数据类型
// 由于go语言里有可以自定义类型
// 所以反射分为，type 和 kind
// kind主要用来描述对象的底层类型
// 引用类型，像指针切片 type 一般为空
func main() {
	var i myInt = 12
	s := new(string)
	*s = "Hello"
	x := 12
	ReflectType(i)
	ReflectType(s)
	fmt.Println(reflect.TypeOf(x))
	ReflectValue(s)
	ReflectValue(i)
}

// 反射函数
func ReflectType(a any) {
	t := reflect.TypeOf(a)
	fmt.Printf("Type:%v kind:%v", t.Name(), t.Kind())
}

// ReflectValueOf() 返回ReflectValue类型
// 包含一些值信息 原始值信息 指针返回的是内存地址
func ReflectValue(a any) {
	t := reflect.ValueOf(a)
	fmt.Println(t)
}

// 在反射中修改值
// 使用Elem()来获取指针对应的值
func ReflectSetValue(a any) {
	v := reflect.ValueOf(a)
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(14)
	}
}
