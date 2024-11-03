package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := math.MinInt32 // 初始化最大和为最小值
	currentSum := 0         // 当前子数组的和

	for _, num := range nums {
		// 如果当前子数组的和加上当前元素仍然小于当前元素本身，
		// 则丢弃当前子数组，从当前元素重新开始一个新的子数组
		if currentSum+num > num {
			currentSum += num
		} else {
			currentSum = num
		}

		// 更新最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
