package allSort

import "testing"

func TestSelectionSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.s)
		})
	}
}
