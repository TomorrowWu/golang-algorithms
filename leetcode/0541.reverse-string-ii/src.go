package _541_reverse_string_ii

func reverseStr(s string, k int) string {
	data := []byte(s)
	n := len(data)
	for i := 0; i < n; i += 2 * k {
		// 剩余字符数
		left := n - i
		if left < k {
			reverseString(data[i:])
		} else {
			reverseString(data[i : i+k])
		}
	}
	return string(data)
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
