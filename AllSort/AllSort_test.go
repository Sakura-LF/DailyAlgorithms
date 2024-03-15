package AllSort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 8, 2, 4, 9, 10, 7}
	SelectionSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 8, 100, 2, 4, 9, 10, 7}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 8, 100, 2, 4, 9, 10, 7}
	InsertionSort(nums)
	fmt.Println(nums)
}
