package Hash

func intersection(nums1 []int, nums2 []int) []int {
	// 使用 map 模拟 set
	set := make(map[int]struct{})
	// 结果集
	var res []int
	// 遍历第一个切片
	for _, val := range nums1 {
		set[val] = struct{}{}
	}
	// 遍历第二个切片
	for _, val := range nums2 {
		// 如果元素在 set 中
		if _, exist := set[val]; exist {
			// 追加到结果集
			res = append(res, val)
			// 删除元素,防止重复 [1, 2, 2] [2, 2]
			delete(set, val)
		}
	}
	return res
}
