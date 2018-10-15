package src

func positiveSolution(n int) []int {
	result := make([]int, n)
	//next表示是否需要跳过这位数字
	next := false
	//i就是待标记的数字
	i := 1
	for i <= n {
		for p := 0; p < n; p++ {
			if result[p] == 0 {
				if !next {
					result[p] = i
					i++
				}
				next = !next
			}
		}
	}

	return result
}
