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

// BubbleSort 冒泡排序
// 每一趟遍历都会把当前未排序部分的最大值“冒泡”到序列的末尾
// 后续逐渐缩小循环边界
// 时间复杂度: O(n^2)
func BubbleSort(s []int) {
	length := len(s)
	// 外循环: 未排序区间未 [0, i]
	for i := length - 1; i > 0; i-- {
		// 内循环:
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

/* 冒泡排序（标志优化）*/
// 如果冒泡排序中一轮没有发生值的交换,就说明排序已经
// 完成
// 所有文件里进行坚持那哈数操作git
// 提交都可以非常好用
func bubbleSortWithFlag(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true // 记录交换元素
			}
		}
		if flag == false { // 此轮“冒泡”未交换任何元素，直接跳出
			break
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
		// 软禁都可以坚持那哈数
		// linuxw  wenjain1zong1suoyou1caozuog1oudhis1kyie
		// 所有操作函数都一个可以进朝韩表示的1哟说有李1
	}
}
