package main

// BinarySearch 二分查找
func BinarySearch(s []int, target int) int {
	left := 0
	right := len(s) - 1

	for left <= right {
		// 中间
		middle := left + (right-left)>>1
		if s[middle] > target {
			right = middle - 1 // target 在左区间
		}
		if s[middle] < target {
			left = middle + 1 // target 在右区间
		}
		if s[middle] == target {
			return middle
		}
	}
	return -1
}

// GELeft 需求: 二分查找到 >= target 的最左位置
// [2, 3, 3, 3, 7, 8, 9, 14] -> 找大于等于5的最左位置 -> 下标: 4 -> 7
func GELeft(s []int, target int) int {
	result := -1
	left := 0
	right := len(s) - 1

	for left <= right {
		middle := left + (right-left)>>1
		// 中间值大于目标值,后面的都砍掉
		if s[middle] >= target {
			result = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return result
}

// LTRight 需求: 二分查找找到 <= target 的最右
// [1, 2, 3, 3, 3, 4, 7, 8, 9, 14] -> 找小于等于5的最右位置 -> 下标: 3 -> 3
func LTRight(s []int, target int) int {
	length := len(s)

	left := 0
	right := length - 1
	result := 0

	for left <= right {
		middle := left + (right-left)>>1
		if s[middle] <= target {
			result = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return result
}
