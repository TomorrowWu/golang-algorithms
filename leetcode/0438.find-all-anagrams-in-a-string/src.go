package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	l1, l2 := len(s), len(p)
	if l1 < l2 {
		return []int{}
	}

	ans := []int{}
	c1 := [26]int{}
	c2 := [26]int{}
	for i, ch := range p {
		c1[ch-'a']++
		c2[s[i]-'a']++
	}
	if c1 == c2 {
		ans = append(ans, 0)
	}

	for i := 1; i <= l1-l2; i++ {
		c2[s[i-1]-'a']--
		c2[s[i+l2-1]-'a']++
		if c2 == c1 {
			ans = append(ans, i)
		}
	}
	return ans
}
