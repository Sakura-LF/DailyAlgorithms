package allSort

// InsertionSort 插入排序
// 从数组索引1开始,循环判断左边大于当前索引
// 时间复杂度: O(n^2)
func InsertionSort(s []int) {
	length := len(s)

	for i := 1; i < length; i++ {
		// j 当前索引的左边的元素
		// j + 1 当前索引
		for j := i - 1; j >= 0 && s[j] > s[j+1]; j-- {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}

// InsertionSort1 /* 插入排序（标志优化）*/
// 时间复杂度: O(n^2)
// 如果插入
func InsertionSort1(s []int) {
	length := len(s)
	for i := 1; i < length; i++ {
		bash := s[i]
		// j 当前索引的左边的元素'
		j := i - 1
		for ; j >= 0 && s[j] > bash; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = bash
	}
}
