package allSort

import "testing"

func TestBubbleSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.s)
		})
	}
}
