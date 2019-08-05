package main

func numPairsDivisibleBy60(time []int) int {
	l := len(time)
	if l <= 1 {
		return 0
	}
	count := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j <= l-1; j++ {
			if (time[i]+time[j])%60 == 0 {
				count += 1
			}
		}
	}
	return count
}

func numPairsDivisibleBy602(time []int) int {
	m := make([]int, 60)
	cnt := 0
	// 通过取余数
	for _, n := range time {
		if n%60 == 0 {
			// 已经是60的倍数
			cnt += m[0]
		} else {
			// (time[i] + time[j]) % 60 == 0
			cnt += m[60-n%60]
		}
		// 把当前值计入次数
		m[n%60]++
	}
	return cnt
}
