package src

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := [][]int{}
	for i := range nums {
		l := make([]int, len(nums)-1)
		copy(l[:i], nums[:i])
		copy(l[i:], nums[i+1:])
		for _, x := range permute(l) {
			//以 nums[i] 为首的排列
			res = append(res, append([]int{nums[i]}, x...))
		}
	}
	return res
}
