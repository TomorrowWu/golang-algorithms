package main

import "fmt"

//给定一个字符串，找出不含有重复字符的最长子串的长度。
//
//示例：
//
//给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
//
//给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
//
//给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。

//int [26] 用于字母 ‘a’ - ‘z’或 ‘A’ - ‘Z’
//int [128] 用于ASCII码
//int [256] 用于扩展ASCII码
func lengthOfLongestSubstring(s string) int {
	maxStrLen := 0
	charData := make(map[int32]int, 128) //减少内存分配
	p := 0                               //重复字符的上次出现的索引,上一个重复字符和和当前重复字符的之间的子串即有效子串
	for i, char := range s {
		if j, ok := charData[char]; ok {
			//字符重复,从重复字符下一位开始重新算子串
			if j > p {
				p = j
			}
		}
		//防止字符串长度为1时,i-p=0
		if i+1-p > maxStrLen {
			maxStrLen = i + 1 - p
		}
		charData[char] = i + 1
	}
	return maxStrLen
}

func main() {
	//len := lengthOfLongestSubstring(" ")
	//len := lengthOfLongestSubstring("dvdf")
	len := lengthOfLongestSubstring("abcabcbb")
	//len := lengthOfLongestSubstring("abba")
	fmt.Println("len=", len)
}
