package utest

// 用于测试的函数
// go 的测试依赖go test命令
// go test 测试以*_test.go结尾的文件
// 测试有单元测试，基准测试，示例函数
// 单测主要测程序能不能正常执行
// benchmark 主要测性能
// example 提供一些示例代码
// 使用go test -v 查看详细的信息
func Hello(a, b int) int {
	return a + b
}
