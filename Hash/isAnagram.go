package Hash

func isAnagram(s string, t string) bool {
	m := make(map[string]int)
	// 循环将字符串放入map
	for _, val := range s {
		// 如果存在,个数加 1
		if _, exist := m[string(val)]; exist {
			m[string(val)]++
		} else { // 如果是新加入的,置为 1
			m[string(val)] = 1
		}
	}
	//fmt.Println(m)
	// 遍历第二个字符串
	for _, val := range t {
		// 字母如果存在, -1
		if _, exist := m[string(val)]; exist {
			// 将 -1 后的 i 赋值回去
			m[string(val)]--
			// 如果字数减到 0 了,删除这个 key
			if m[string(val)] == 0 {
				delete(m, string(val))
			}
		} else {
			// 第二个字符串中的字符如果不在第一个里面, 直接返回 false
			return false
		}
	}
	//fmt.Println(m)
	// 如果 map 为空,说明是异位词
	return len(m) == 0
}

func isAnagram2(s string, t string) bool {
	// 初始化一个长度 26 的数组代表字母
	array := [26]int{}

	// value是字母的ascII码值 - 'a' 就是每个字母的索引
	// a : 97 - 97 = 0 索引 再++,即为索引为0的的元素加一
	// b : 98 - 97 = 1 索引
	for _, val := range s {
		array[val-'a']++
	}
	for _, val := range t {
		array[val-'a']--
	}
	return array == [26]int{}
}
