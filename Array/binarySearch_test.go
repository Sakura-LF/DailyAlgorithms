package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		s      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "BinarySearch",
			args: args{
				s:      []int{3, 43, 23, 5, 67, 89, 5, 22, 45},
				target: 7,
			},
			want: 4,
		},
		{
			name: "BinarySearch",
			args: args{
				s:      []int{1, 3, 4, 6, 7, 13, 45, 67, 78, 90},
				target: 20,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGELeft(t *testing.T) {
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
