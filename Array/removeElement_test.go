package main

import "testing"

func Test_removeElement3(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{
			nums: []int{3, 43, 23, 5, 67, 89, 5, 22, 45},
			val:  5,
		},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before:", tt.args.nums, " delete:", tt.args.val)
			if got := removeElement3(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement3() = %v, want %v", got, tt.want)
			}
			t.Log("After:", tt.args.nums, " left:", tt.want)

		})
	}
}
