package ti_huan_kong_ge_lcof

// 遍历添加
func replaceSpace(s string) string {
	ans := make([]byte, 0, len(s))
	for i, _ := range s {
		if s[i] == ' ' {
			ans = append(ans, []byte("%20")...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

// 原地扩容切片+双指针修改
func replaceSpace2(s string) string {
	data := []byte(s)
	l1 := len(data)
	spaceCnt := 0
	for _, v := range data {
		if v == ' ' {
			spaceCnt++
		}
	}

	// 扩容切片
	tmp := make([]byte, 2*spaceCnt)
	data = append(data, tmp...)

	// i为要移动的元素下标，j为空位置
	for i, j := l1-1, len(data)-1; i >= 0; {
		if data[i] == ' ' {
			data[j] = '0'
			data[j-1] = '2'
			data[j-2] = '%'
			i--
			j -= 3
		} else {
			data[j] = data[i]
			i--
			j--
		}
	}
	return string(data)
}
