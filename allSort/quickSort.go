package allSort

// 快速排序的哨兵划分
func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		// 从右向左,寻找小于基准数的元素, 不小于继续往前走
		for i < j && nums[j] >= nums[left] {
			j--
		}
		// 从左往右,寻找大于基准数的元素,不大于继续往后走
		for i < j && nums[i] <= nums[left] {
			i++
		}
		// 左边大于基准数和右边小于基准数的进行交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 交换基准数到两个子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	//返回基准数索引
	return i
}

// 快速排序流程
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 哨兵划分
	i := partition(nums, left, right)
	// 递归遍历左子数组和右子数组
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
