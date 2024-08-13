package Hash

func IsHappy(n int) bool {
	m := make(map[int]bool)

	// 无限循环,直到 sum 重复
	for n != 1 && !m[n] {
		n = getSum(n)
		m[n] = true
	}

	return n == 1
}

// 计算 sum 的值
func getSum(n int) int {
	sum := 0
	// 每次循环计算一位的平方和
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
