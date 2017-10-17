package 字符串模式匹配算法

/**
 * WuMing 
 *2017/10/16 下午7:54
 *
 */

/**
 吴名
 2017/10/16 下午7:54
暴力匹配
 */
func violentMatch(s, p string) int {
	lS := len(s)
	lP := len(p)

	i, j := 0, 0
	for ; i < lS && j < lP; {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == lP {
		return i - j
	} else {
		return -1
	}
}
