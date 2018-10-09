package tagvalidator

import "strings"

func isValid(code string) bool {
	// isAvailbleTagName 返回 true， 当 code[i:j] 符合 TAG_NAME 的命名规范时
	isAvailbleTagName := func(i, j int) bool {
		if i > j || //code[i:] 中没有 ">"
			i == j || // TAG_NAME 为 ""
			i+9 < j { // TAG_NAME 长度超过 9
			return false
		}
		for k := i; k < j; k++ {
			if code[k] < 'A' || code[k] > 'Z' {
				// TAG_NAME 中存在非大写字母
				return false
			}
		}
		return true
	}

	//store tag(start,end)
	stack := make([]string, 0, 16)
	for i := 0; i < len(code); {
		// 所有的内容都应该在 tag 中
		// 所以，除非 i == 0
		// stack 中都应该至少有一个 TAG_NAME
		if i > 0 && len(stack) == 0 {
			return false
		}

		switch {
		case strings.HasPrefix(code[i:], "<![CDATA["):
			i += 9
			//j is index of "]]>" in code
			j := i + strings.Index(code[i:], "]]>")
			//not found "]]>"
			if i > j {
				return false
			}
			i = j + 3
		case strings.HasPrefix(code[i:], "</"):
			//end tag
			i += 2
			//j is index of ">" in code
			j := i + strings.Index(code[i:], ">")
			//code[i,j] is tagName
			if !isAvailbleTagName(i, j) {
				return false
			}
			tagName := code[i:j]
			i = j + 1
			n := len(stack)
			//there is no start tag || tagName not match
			if n == 0 || stack[n-1] != tagName {
				return false
			}
			//start tag , end tag match,remove the tags
			stack = stack[:n-1]
		case code[i] == '<':
			//start tag
			i++
			j := i + strings.Index(code[i:], ">")
			if !isAvailbleTagName(i, j) {
				return false
			}
			tagName := code[i:j]
			i = j + 1
			stack = append(stack, tagName)
		default:
			i++
		}
	}
	// len(stack) != 0 说明， TAG 没有封闭
	return len(stack) == 0
}
