package src

import "bytes"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var res bytes.Buffer
Here:
	for i := 0; ; i++ {

		//检查当前字母是否相等
		var cur byte
		for _, str := range strs {

			//检查是否越界
			if len(str) <= i {
				break Here
			}

			if cur == 0 {
				cur = str[i]
				continue
			}

			if str[i] != cur {
				//字母相等，
				break Here
			}
		}
		res.WriteByte(cur)
	}

	return res.String()
}
