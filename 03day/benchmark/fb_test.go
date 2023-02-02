package benchmark

import "testing"

func BenchmarkFb1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(1)
	}
}

// benchmark 至少要运行1秒钟
// 性能基准测试
// 使用命令 go test -bench=Fb
// benchmark 用来测试程序的性能

func BenchmarkFb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(2)
	}
}

func BenchmarkFb10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(10)
	}
}

func BenchmarkFb20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(20)
	}
}
