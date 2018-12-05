package src

//int [26] 用于字母 ‘a’ - ‘z’或 ‘A’ - ‘Z’
//int [128] 用于ASCII码
//int [256] 用于扩展ASCII码
func lengthOfLongestSubstring(s string) int {
	maxStrLen := 0
	charData := make(map[int32]int, 128) //减少内存分配
	p := 0                               //重复字符的上次出现的位置,即为子串起点
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
