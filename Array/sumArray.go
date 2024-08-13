package main

// 需求:
// 假设有一个数组arr,
// 用户总是频繁的查询arr中范围的累加和
// 你如何组织数据，
// 能让这种查询变得每次都特别快捷？

// 解决方法:
// 新构建一个数组包含,数组范围的和
//
//	原数组:   s[2, 5, 3, -1, 6, 2]
//
// 前缀和数组: sum[2, 7, 10, 9, 15, 17]
// 1. 求数组 0-5的和: sum[5]
// 2. 求数组 2-4的和: sum[4] - sum[1]
//func main() {
//	nums := []int{2, 5, 3, -1, 6, 2}
//	array := CreateSumArray(nums)
//	fmt.Println(array)
//}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//
//	scanner.Scan()
//	size, _ := strconv.Atoi(scanner.Text())
//
//	nums := make([]int, size)
//	for i := 0; i < size; i++ {
//		scanner.Scan()
//		nums[i], _ = strconv.Atoi(scanner.Text())
//	}
//	//fmt.Println(nums)
//	// 构建前缀和
//	sumArray := CreateSumArray(nums)
//
//	for scanner.Scan() {
//		var left int
//		var right int
//		line := scanner.Text()
//		_, err := fmt.Sscanf(line, "%d %d", &left, &right)
//		if err != nil {
//			log.Println(err)
//			continue
//		}
//		if left < 0 || right >= size {
//			continue
//		}
//
//		fmt.Println(GetSum(sumArray, left, right))
//	}
//}

// CreateSumArray 构建前缀和数组
func CreateSumArray(s []int) []int {
	nums := make([]int, len(s))

	//for i := 0; i < len(s); i++ {
	//	num += s[i]
	//	nums[i] = num
	//}
	nums[0] = s[0]
	for i := 1; i < len(nums); i++ {
		// 前缀数组的前一个元素 + 数组的这个元素
		nums[i] = nums[i-1] + s[i]
	}
	return nums
}

func GetSum(s []int, l int, r int) int {
	if l == 0 {
		return s[r]
	} else {
		return s[r] - s[l-1]
	}
}
