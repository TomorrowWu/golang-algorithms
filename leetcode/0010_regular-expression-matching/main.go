package main

//题目描述: https://leetcode-cn.com/problems/regular-expression-matching/description/

//+++++++++++++++++++++++++++++++正确的算法+++++++++++++++++++++++++++++++++

// 程序中存在以下假设
// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

//动态规划
func isMatch3(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true

	for i := 2; i < len(dp); i += 2 {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	for i := 1; i < len(dp); i++ {
		if i < len(p) && p[i] == '*' {
			continue
		}

		for j := 1; j < len(dp[0]); j++ {

			if p[i-1] == '*' {
				if p[i-2] == '.' {
					dp[i][j] = dp[i-2][j-1] || dp[i][j-1] || dp[i-2][j]
				} else {
					dp[i][j] = (dp[i-2][j]) || (p[i-2] == s[j-1] && (dp[i-2][j-1] || dp[i][j-1]))
				}
			} else if p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
			}
		}
	}

	return dp[len(p)][len(s)]
}

//动态规划--递归版本(本人自研)
func isMatch4(s string, p string) bool {
	lenP := len(p)
	lenS := len(s)

	switch lenP {
	case 0:
		if lenS == 0 {
			return true
		} else {
			return false
		}
	case 1:
		if lenS != 1 {
			return false
		}
		if p[0] == '.' || s[0] == p[0] {
			return true
		} else {
			return false
		}
	case 2:
		switch p[1] {
		case '*':
			if p[0] == '.' {
				return true
			} else {
				for _, char := range s {
					if char != int32(p[0]) {
						return false
					}
				}
				return true
			}
		case '.':
			if lenS != 2 {
				return false
			}
			return isMatch4(s[:1], p[:1])
		default:
			if lenS != 2 {
				return false
			}
			if s[1] != p[1] {
				return false
			}
			return isMatch4(s[:1], p[:1])
		}
	default:
		if lenS == 0 {
			//""   "a*c*"
			if p[lenP-1] == '*' {
				return isMatch4(s, p[:lenP-2])
			} else {
				return false
			}
		}
		if p[lenP-1] == '*' {
			if isMatch4(s, p[:lenP-2]) || ((s[lenS-1] == p[lenP-2] || p[lenP-2] == '.') && isMatch4(s[:lenS-1], p)) {
				return true
			} else {
				return false
			}
		} else if p[lenP-1] == '.' {
			return isMatch4(s[:lenS-1], p[:lenP-1])
		} else {
			if s[lenS-1] == p[lenP-1] && isMatch4(s[:lenS-1], p[:lenP-1]) {
				return true
			} else {
				return false
			}
		}
	}
}

func main() {

}
