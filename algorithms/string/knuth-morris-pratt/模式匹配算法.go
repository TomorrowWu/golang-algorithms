package knuth_morris_pratt

import "fmt"

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
//优化过后的next 数组求法
func genNextArr(patternStr string) []int {
	length := len(patternStr)
	arr := make([]int, length)
	arr[0] = -1
	prefixI, suffixI := 0, 1
	for suffixI < length-1 {
		fmt.Printf("pre:%v,suf:%v,%v\n", prefixI, suffixI, arr[suffixI])
		//patternStr[prefixI]表示前缀,patternStr[suffixI]表示后缀
		if prefixI == -1 || (patternStr[prefixI] == patternStr[suffixI]) {
			prefixI++
			suffixI++
			if patternStr[prefixI] == patternStr[suffixI] {
				//从next数组取值的时机是失配时,那么如果这里相等,那么失配时移动到prefixI,还是失配,还是需要在从next数组取值,所以这里把这一步直接去掉
				//相当于递归查找失配元素的下标
				arr[suffixI] = arr[prefixI]
			} else {
				//next数组中存的都是失配时前缀失配元素的下标
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
