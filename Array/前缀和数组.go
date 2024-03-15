package main

import (
	"fmt"
)

// 需求:
// 假设有一个数组arr,
// 用户总是频繁的查询arr中范围的累加和
// 你如何组织数据，
// 能让这种查询变得每次都特别快捷？

// 解决方法:
// 新构建一个数组包含,数组范围的和
//     原数组:   s[2, 5, 3, -1, 6, 2]
// 前缀和数组: sum[2, 7 , 10, 9, 15, 16]
// 1. 求数组 0-5的和: sum[5]
// 2. 求数组 2-4的和: sum[4] - sum[1]

func main() {
	s := []int{1, 5, 6, 9, 4, 8}
	array := CreateSumArray(s)
	fmt.Println("原数组:", s)
	fmt.Println("前缀和数组:", array)

	fmt.Println("0-4的和:", GetSum(array, 0, 4))
	fmt.Println("2-4的和:", GetSum(array, 2, 4))
}

// CreateSumArray 构建前缀和数组
func CreateSumArray(s []int) []int {
	sum := make([]int, len(s))
	// 两种方式计算前缀和
	// 第一种
	sum[0] = s[0]
	for i := 1; i < len(s); i++ {
		//num += s[i]
		sum[i] = sum[i-1] + s[i]
	}

	// 第二种
	//num := 0
	//for i := 0; i < ; i++ {
	//	num += s[i]
	//	sum[i] = num
	//}
	return sum
}

func GetSum(s []int, l int, r int) int {
	if l == 0 {
		return s[r]
	} else {
		return s[r] - s[l-1]
	}
}
