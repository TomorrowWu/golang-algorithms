[0591_标签验证器](https://leetcode-cn.com/problems/tag-validator/description/)

## 题目描述

给定一个表示代码片段的字符串，你需要实现一个验证器来解析这段代码，并返回它是否合法。合法的代码片段需要遵守以下的所有规则：

1. 代码必须被合法的闭合标签包围。否则，代码是无效的。
2. 闭合标签（不一定合法）要严格符合格式：<TAG_NAME>TAG_CONTENT</TAG_NAME>。其中，<TAG_NAME>是起始标签，</TAG_NAME>是结束标签。起始和结束标签中的 TAG_NAME 应当相同。当且仅当 TAG_NAME 和 TAG_CONTENT 都是合法的，闭合标签才是合法的。
3. 合法的 TAG_NAME 仅含有大写字母，长度在范围 [1,9] 之间。否则，该 TAG_NAME 是不合法的。
4. 合法的 TAG_CONTENT 可以包含其他合法的闭合标签，cdata （请参考规则7）和任意字符（注意参考规则1）除了不匹配的<、不匹配的起始和结束标签、不匹配的或带有不合法 TAG_NAME 的闭合标签。否则，TAG_CONTENT 是不合法的。
5. 一个起始标签，如果没有具有相同 TAG_NAME 的结束标签与之匹配，是不合法的。反之亦然。不过，你也需要考虑标签嵌套的问题。
6. 一个<，如果你找不到一个后续的>与之匹配，是不合法的。并且当你找到一个<或</时，所有直到下一个>的前的字符，都应当被解析为 TAG_NAME（不一定合法）。
7. cdata 有如下格式：<![CDATA[CDATA_CONTENT]]>。CDATA_CONTENT 的范围被定义成 <![CDATA[ 和后续的第一个 ]]>之间的字符。
8. CDATA_CONTENT 可以包含任意字符。cdata 的功能是阻止验证器解析CDATA_CONTENT，所以即使其中有一些字符可以被解析为标签（无论合法还是不合法），也应该将它们视为常规字符。

合法代码的例子::
```
输入: "<DIV>This is the first line <![CDATA[<div>]]></DIV>"

输出: True

解释:

代码被包含在了闭合的标签内： <DIV> 和 </DIV> 。

TAG_NAME 是合法的，TAG_CONTENT 包含了一些字符和 cdata 。

即使 CDATA_CONTENT 含有不匹配的起始标签和不合法的 TAG_NAME，它应该被视为普通的文本，而不是标签。

所以 TAG_CONTENT 是合法的，因此代码是合法的。最终返回True。


输入: "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"

输出: True

解释:

我们首先将代码分割为： start_tag|tag_content|end_tag 。

start_tag -> "<DIV>"

end_tag -> "</DIV>"

tag_content 也可被分割为： text1|cdata|text2 。

text1 -> ">>  ![cdata[]] "

cdata -> "<![CDATA[<div>]>]]>" ，其中 CDATA_CONTENT 为 "<div>]>"

text2 -> "]]>>]"


start_tag 不是 "<DIV>>>" 的原因参照规则 6 。
cdata 不是 "<![CDATA[<div>]>]]>]]>" 的原因参照规则 7 。
```

不合法代码的例子::
```
输入: "<A>  <B> </A>   </B>"
输出: False
解释: 不合法。如果 "<A>" 是闭合的，那么 "<B>" 一定是不匹配的，反之亦然。

输入: "<DIV>  div tag is not closed  <DIV>"
输出: False

输入: "<DIV>  unmatched <  </DIV>"
输出: False

输入: "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"
输出: False

输入: "<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"
输出: False

输入: "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"
输出: False
```

注意:

1. 为简明起见，你可以假设输入的代码（包括提到的任意字符）只包含数字, 字母, '<','>','/','!','[',']'和' '。

## 代码实现

```Golang
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
```

## GitHub
- [源码在这里](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0591.tag-validator/main.go)

## 公众号「tomorrowwu」，编程·职场·思维

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
