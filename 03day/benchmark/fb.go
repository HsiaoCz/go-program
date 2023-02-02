package benchmark

// 求斐波那契数列数列
func Fb(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return Fb(num-1) + Fb((num - 2))
}
