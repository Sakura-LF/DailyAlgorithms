package allSort

import (
	"fmt"
	"testing"
)

// 测试选择排序
func TestSelectionSort(t *testing.T) {
	nums := []int{5, 8, 2, 4, 9, 10, 7}
	SelectionSort(nums)
	fmt.Println(nums)
}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	nums := []int{5, 8, 100, 2, 4, 9, 10, 7}
	BubbleSort(nums)
	fmt.Println(nums)
}

// 测试插入排序
func TestInsertionSort(t *testing.T) {
	nums := []int{5, 8, 100, 2, 4, 9, 10, 7}
	InsertionSort1(nums)
	fmt.Println(nums)
}

// 测试归并排序
