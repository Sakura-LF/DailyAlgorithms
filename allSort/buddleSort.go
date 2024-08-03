package allSort

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
