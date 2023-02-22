package _076_minimum_window_substring

func minWindow(s string, t string) string {
	cnt := map[byte]int{}
	for i := range t {
		cnt[t[i]]++
	}

	cntAns := map[byte]int{}
	l, r := -1, -1
	left := 0
	for right := range s {
		cntAns[s[right]]++
		for mapContain(cntAns, cnt) {
			if l == -1 || right-left < r-l {
				l, r = left, right
			}
			cntAns[s[left]]--
			left++
		}
	}
	if l == -1 {
		return ""
	}
	return s[l : r+1]
}

func mapContain(a, b map[byte]int) bool {
	if len(b) > len(a) {
		return false
	}
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}
