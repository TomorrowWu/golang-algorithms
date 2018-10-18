package src

func quickSort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}

	pivot := a[0]
	head := 0
	tail := size - 1
	for i := 1; i <= tail; {
		if a[i] > pivot {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[head], a[i] = a[i], a[head]
			head++
			i++
		}
	}
	quickSort(a[:head])
	quickSort(a[head+1:])
}

func findKthLargest(nums []int, k int) int {
	size := len(nums)
	if size < k {
		return -1
	}

	pivot := nums[0]
	head := 0
	tail := size - 1
	for i := 1; i <= tail; {
		if nums[i] < pivot {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		}
	}

	if head+1 == k {
		return nums[head]
	}

	if head+1 < k {
		return findKthLargest(nums[head+1:], k-head-1)
	} else {
		return findKthLargest(nums[:head], k)
	}
}
