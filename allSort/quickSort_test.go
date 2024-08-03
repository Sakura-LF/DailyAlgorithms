package allSort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "QuickSort",
			args: args{
				nums:  []int{2, 4, 1, 0, 3, 5},
				left:  0,
				right: 5,
			},
			// 最后基准数的索引
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Before:", tt.args.nums)
			if got := partition(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("After:", tt.args.nums)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "QuickSort", args: args{
				nums:  []int{2, 4, 1, 0, 3, 5},
				left:  0,
				right: 5,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Before:", tt.args.nums)
			quickSort(tt.args.nums, tt.args.left, tt.args.right)
			fmt.Println("After:", tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("Partition() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
