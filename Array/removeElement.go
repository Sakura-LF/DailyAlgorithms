package main

// RemoveElement 移除元素
func removeElement(nums []int, val int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		// 如果遇到相等元素
		if nums[i] == val {
			// 从这个元素的后一位开始
			// 将后面的所有元素向前移动
			for j := i + 1; j < length; j++ {
				nums[j-1] = nums[j]
			}
			// 原本是 i 的元素也会先前
			i--
			// 长度 -1
			length--
		}
	}
	return length
}

func removeElement2(nums []int, val int) int {
	length := len(nums)
	// 慢指针
	slowIndex := 0
	for fastIndex := 0; fastIndex < length; fastIndex++ {
		// 当快指针指向的元素 == val
		// 跳过本次循环
		if nums[fastIndex] == val {
			continue
		}
		// 否则交换元素
		nums[slowIndex] = nums[fastIndex]
		slowIndex++
	}
	return slowIndex
}

func removeElement3(nums []int, val int) int {
	// left指向更新后的元素
	left := 0
	// right指向要删除的元素
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
