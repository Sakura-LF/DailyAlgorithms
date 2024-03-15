package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 8, 2, 4, 9, 10, 7}
	SelectionSort(nums)
	//fmt.Println(nums)
	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})
	//fmt.Println(nums)
}
