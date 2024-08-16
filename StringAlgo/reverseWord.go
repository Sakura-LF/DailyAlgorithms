package StringAlgo

import (
	"strings"
)

func reverseWords(s string) string {
	// 将s转换为字符串
	ByteSlice := []byte(s)

	// 一:首先使用双指针去除多余空格
	slow := 0
	for i := 0; i < len(ByteSlice); i++ {
		// 快指针要移动到不等于 空格的元素上
		if ByteSlice[i] != ' ' {
			// 如果slow不指向起始位置,添加空格
			if slow != 0 {
				ByteSlice[slow] = ' '
				slow++
			}
			// 虽然是两层for循环,但是时间复杂度是O(n)
			// 在遇到下一个空格
			for i < len(ByteSlice) && ByteSlice[i] != ' ' { // 复制逻辑
				ByteSlice[slow] = ByteSlice[i]
				slow++
				i++
			}
		}
	}
	//去除空格之后,切片截取到slow指针指向的位置
	ByteSlice = ByteSlice[0:slow]

	// 翻转整个字符串
	reverse(ByteSlice)

	// 翻转每个单词
	last := 0
	for i := 0; i <= len(ByteSlice); i++ {
		// 必须先判断 i == len(ByteSlice)
		if i == len(ByteSlice) || ByteSlice[i] == ' ' {
			reverse(ByteSlice[last:i])
			last = i + 1
		}
	}

	return string(ByteSlice)
}

// 反转字符串
func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseWords2(s string) string {
	// Fields 去除多余空格
	fields := strings.Fields(s)

	// 双指针法
	// 交换切片中的每个元素
	left := 0
	right := len(fields) - 1
	for left < right {
		fields[left], fields[right] = fields[right], fields[left]
		left++
		right--
	}
	// 将切片转换为字符串,并且使用空格分割
	return strings.Join(fields, " ")
}
