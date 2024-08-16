package StringAlgo

func reverseString(s []byte) {
	// 双指针
	// 分别指向 字符串的开头和结尾
	left := 0
	right := len(s) - 1

	// left = right 的时候说明是奇数,并且已经反转完了
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
