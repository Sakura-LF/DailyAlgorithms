package main

func GenerateMatrix(n int) [][]int {
	// 定义 x 和 y 轴
	startx, starty := 0, 0

	// 矩阵循环的次数
	// 奇数,eg,3  3 / 2 = 1 中间单独赋值
	// 偶数,eg,4  4 / 2 = 2
	loop := n / 2

	// 每次一边遍历的长度
	offset := 1

	// 要填入矩阵的数
	num := 1

	// 二维数组
	res := make([][]int, n)
	for i := range res { // 初始化数组
		res[i] = make([]int, n)
	}

	for loop > 0 {
		x := startx
		y := starty

		for ; y < n-offset; y++ {
			res[x][y] = num
			num++
		}
		//fmt.Println(res)

		for ; x < n-offset; x++ {
			res[x][y] = num
			num++
		}
		//fmt.Println(res)

		for ; y > starty; y-- {
			res[x][y] = num
			num++
		}

		for ; x > startx; x-- {
			res[x][y] = num
			num++
		}

		startx++
		starty++
		offset++

		// 循环次数 -1
		loop--
	}
	// 如果是奇数,单独赋值最中间的值
	if n%2 == 1 {
		res[n/2][n/2] = num
	}

	return res
}
