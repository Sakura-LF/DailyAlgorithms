package main

func sortedSquares(nums []int) []int {
	length := len(nums)
	extraSlice := make([]int, length)
	// 左右指针
	left, right := 0, length-1

	k := length - 1
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			extraSlice[k] = leftSquare
			left++
		}
		if leftSquare < rightSquare {
			extraSlice[k] = rightSquare
			right--
		}
		if leftSquare == rightSquare {
			extraSlice[k] = nums[left] * nums[left]
			left++
		}
		k--
	}
	return extraSlice
}
