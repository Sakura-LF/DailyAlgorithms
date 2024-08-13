package Hash

import (
	"fmt"
	"testing"
)

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
