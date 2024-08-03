package allSort

// SelectionSort 选择排序
// 思路: 每一轮从剩余未排序的部分中找到最小的元素，并将其与已排序部分的末尾元素交换位置
// 初始状态下,所有元素未排序
// 每一次循环比出一位最小的数
// 时间复杂度: O(n^2)
func SelectionSort(s []int) {
	length := len(s)
	// 不需要比最后一个数
	for i := 0; i < length-1; i++ {
		// 使用 minIndex 记录未排序区间的最小元素
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		// 将为排序元素和选出来的最小元素交换
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
