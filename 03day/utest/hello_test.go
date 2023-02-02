package utest

import (
	"testing"
)

func TestHello(t *testing.T) {
	get := Hello(12, 23)
	want := 35
	if get != want {
		t.Errorf("get:%d want:%d\n", get, want)
	}
}

// 测试组
// 使用一个结构体存储想要输入的数据和期望值
// 测试组的好处是 可以直接添加测试组
func TestHelloGroup(t *testing.T) {
	type helloGroup struct {
		firstNum int
		lastNum  int
		want     int
	}
	// 测试组
	var testHelloGroup = []helloGroup{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 4, 8},
		{12, 13, 25},
	}

	for _, th := range testHelloGroup {
		get := Hello(th.firstNum, th.lastNum)
		if get != th.want {
			t.Errorf("get:%d,want:%d\n", get, th.want)
		}
	}
}

// 子测试
// 测试覆盖率 go test -cover
func TestHelloSon(t *testing.T) {
	type helloGroup struct {
		firstNum int
		lastNum  int
		want     int
	}
	testHelloSon := map[string]helloGroup{
		"case1": {1, 2, 3},
		"case2": {2, 3, 5},
		"case3": {3, 4, 7},
		"case4": {4, 4, 8},
	}
	//运行子测试
	for name, th := range testHelloSon {
		t.Run(name, func(t *testing.T) {
			get := Hello(th.firstNum, th.lastNum)
			if get != th.want {
				t.Errorf("get:%d,want:%d\n", get, th.want)
			}
		})
	}
}
