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
	for i < lS && j < lP {
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

/**
 吴名
 2017/10/19 下午6:07
生成next数组
 */
func genNextArr(patternStr string) []int {
	length := len(patternStr)
	arr := make([]int, length)
	arr[0] = -1
	prefixI, suffixI := -1, 0
	for suffixI < length-1 {
		if prefixI == -1 || (patternStr[prefixI] == patternStr[suffixI]) {
			prefixI++
			suffixI++
			if patternStr[prefixI] == patternStr[suffixI] {
				arr[suffixI] = arr[prefixI]
			} else {
				arr[suffixI] = prefixI
			}

		} else {
			prefixI = arr[prefixI]
		}
	}
	return arr
}

/**
 吴名
 2017/10/19 下午8:29
KMP匹配
 */
func KMPSearch(source string, pattern string) int {

	ls := len(source)
	lp := len(pattern)

	si := 0
	pi := 0
	nextPatternIndexArray := genNextArr(pattern)
	for si < ls && pi < lp {
		if pi == -1 || (source[si] == pattern[pi]) {
			si++
			pi++
		} else {
			pi = nextPatternIndexArray[pi]
		}
	}
	if pi == lp {
		return si - pi
	} else {
		return -1
	}
}
