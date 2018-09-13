package main

import "strings"

//本人自己写的算法(执行用时：2428 ms)
func findSubstring(s string, words []string) []int {
	result := make([]int, 0)

	wordsLen := 0
	heads := make(map[uint8]bool, len(words))
	tails := make(map[uint8]bool, len(words))
	for _, word := range words {
		if len(word) > 0 {
			wordsLen += len(word)
			heads[word[0]] = true
			tails[word[len(word)-1]] = true
		}
	}

	strLen := len(s)
	if strLen < wordsLen || wordsLen == 0 {
		return result
	}

	for i, a := range s {
		//剩下的子串长度小于words长度
		if strLen-i < wordsLen {
			break
		}

		//字符和head相等
		isHead, _ := heads[uint8(a)]
		isTail, _ := tails[s[i+wordsLen-1]]
		//头和尾符合words要求
		if isHead && isTail {
			newStr := s[i : i+wordsLen]
			//比较子串是否等于words的组合
			newWords := append([]string{}, words[0:]...)
			for j := 0; j < len(words); j++ {
				for k, word := range newWords {
					//每次去掉一个word,words中每一个都得有
					if strings.HasPrefix(newStr, word) {
						newStr = strings.TrimPrefix(newStr, word)
						newWords = append(newWords[:k], newWords[k+1:]...)
						break
					}
				}
			}

			if len(newStr) == 0 {
				result = append(result, i)
			}
		}
	}

	return result
}
