package string

import "log"

/**
 * WuMing
 *2017/9/15 上午1:40
 *
 */

/**
 吴名
 2017/9/15 上午2:00
朴素的模式匹配算法
*/

/* 返回子串T在主串S中第pos个字符之后的索引。
若不存在，则函数返回值为0。 */
func index(s, t string) int {
	ls := len(s)
	lt := len(t)

	//i用于主串s中当前位置下标,若pos不为1,则从pos位置开始匹配
	i := 0
	//j用于子串t中的当前位置下标
	j := 0
	count := 0
	for i <= ls-1 && j <= lt-1 {
		count++
		//若两字母相等
		if s[i] == t[j] {
			i++
			j++
		} else {
			//指针后退重新开始分配

			//i退回上次匹配首位的下一位
			i = i - j + 1
			//j退回到子串t的首位
			j = 0
		}
	}
	log.Printf("循环次数:%v", count)
	if j == lt {
		return i - lt
	} else {
		return 0
	}
}
