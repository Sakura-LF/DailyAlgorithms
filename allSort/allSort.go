package allSort

// SelectionSort 选择排序
// 思路: 每一轮从剩余未排序的部分中找到最小的元素，并将其与已排序部分的末尾元素交换位置
// 每一次循环比出一位最小的数
// 时间复杂度: O(n^2)
func SelectionSort(s []int) {
	length := len(s)
	// 不需要比最后一个数
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}

// BubbleSort 冒泡排序
// 每一趟遍历都会把当前未排序部分的最大值“冒泡”到序列的末尾
// 后续逐渐缩小循环边界
// 时间复杂度: O(n^2)
func BubbleSort(s []int) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// InsertionSort 插入排序
// 从数组索引1开始,循环判断左边大于当前索引
// 时间复杂度: O(n^2)
func InsertionSort(s []int) {
	length := len(s)

	for i := 1; i < length; i++ {
		// j 当前索引的左边的元素
		for j := i - 1; j >= 0 && s[j] > s[j+1]; j-- {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}
