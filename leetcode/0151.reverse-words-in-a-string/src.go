package _151_reverse_words_in_a_string

// 原地处理
func reverseWords(s string) string {
	l1 := len(s)
	s1 := []byte(s)
	// 移除开头和结尾的空格
	left, right := 0, l1-1
	for left <= right && s[left] == ' ' {
		left++
	}
	for left <= right && s[right] == ' ' {
		right--
	}
	if left > right {
		return ""
	}
	s1 = s1[left : right+1]

	// 移除中间的空格
	// TODO 吴名 2023/2/28 21:31
	l3 := len(s1)
	for i, j := 0, 0; j < l3; {

	}

	// 反转所有字母
	reverseStr(s1)

	// 将单词反转
	l2 := len(s1)
	for i, j := 0, 0; j < l2; {
		if j == l2-1 {
			reverseStr(s1[i:])
			break
		} else if s1[j+1] == ' ' {
			reverseStr(s1[i : j+1])
			i = j + 2
			j = j + 2
		} else {
			j++
		}
	}
	return string(s1)
}

func reverseStr(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
