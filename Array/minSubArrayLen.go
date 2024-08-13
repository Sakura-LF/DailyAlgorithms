package main

import "math"

// 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	// 子数组的长度
	subArrayLen := 0
	// 子数组的和
	sum := 0
	// 最后返回的长度
	result := math.MaxInt32
	// 数组起始位置
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		// 滑动窗口核心
		for sum >= target {
			// 减少子数组的大小
			subArrayLen = j - i + 1
			// 确保是最小的子数组
			if subArrayLen < result {
				result = subArrayLen
			}
			// 减去前一个数
			sum -= nums[i]
			i++
		}
	}
	if result == math.MaxInt32 {
		return 0
	} else {
		return result
	}
}
