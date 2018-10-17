package src

func maxSubArray(nums []int) int {
	size := len(nums)

	sum := nums[0]
	res := sum

	for i := 1; i < size; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
