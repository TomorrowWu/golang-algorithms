package src

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	repeat := make(map[int]bool, len(nums))

	res := [][]int{}
	for i, v := range nums {
		if _, ok := repeat[v]; ok {
			continue
		}

		l := make([]int, len(nums)-1)
		copy(l[:i], nums[:i])
		copy(l[i:], nums[i+1:])
		for _, x := range permuteUnique(l) {
			res = append(res, append([]int{nums[i]}, x...))
		}
		repeat[v] = true
	}

	return res
}
