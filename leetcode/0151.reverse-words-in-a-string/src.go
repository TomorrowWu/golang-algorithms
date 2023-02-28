package _151_reverse_words_in_a_string

// 原地处理
func reverseWords(s string) string {
	l1 := len(s)
	s1 := []byte(s)

	// 多个空格，只留1个空格
	left := 0
	for j := left; j < l1; {
		for s1[j] == ' ' && j+1 < l1 && s1[j+1] == ' ' {
			j++
		}
		s1[left] = s1[j]
		left++
		j++
	}
	s1 = s1[:left]

	if len(s1) >= 2 && s1[0] == ' ' {
		s1 = s1[1:]
	}
	if len(s1) >= 1 && s1[len(s1)-1] == ' ' {
		s1 = s1[:len(s1)-1]
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
