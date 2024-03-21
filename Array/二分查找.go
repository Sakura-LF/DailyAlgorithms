package main

import (
	"fmt"
)

func main() {
	s := []int{1, 3, 4, 6, 7, 13, 45, 67, 78, 90}
	search := BinarySearch(s, 7)
	if search != -1 {
		fmt.Println("找到了下标为:", search)
	} else {
		fmt.Println("没找到")
	}

	left := GELeft(s, 20)
	fmt.Println(left)

	// 二分查找找到 >= target 的最左
	slice1 := []int{2, 3, 3, 3, 7, 8, 9, 14}
	geLeft := GELeft(slice1, 5)
	fmt.Println(slice1, "中：")
	fmt.Println(">= ", 5, "最左为：", slice1[geLeft])

	// 二分查找找到 <= target 的最右
	slice2 := []int{1, 2, 3, 3, 3, 4, 7, 8, 9, 14}
	fmt.Println(slice2, "中：")
	right := LTRight(slice2, 5)
	fmt.Println("<= ", 5, "最右为：", slice2[right])
}

// BinarySearch 二分查找
func BinarySearch(s []int, target int) int {
	left := 0
	right := len(s) - 1

	// 会出现left = right的情况
	for left <= right {
		middle := left + (right-left)>>1
		if s[middle] == target {
			return middle
		} else if s[middle] > target {
			right = middle - 1
		} else if s[middle] < target {
			left = middle + 1
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
