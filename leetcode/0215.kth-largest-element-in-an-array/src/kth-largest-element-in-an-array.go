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

// findKthLargest 类似快速排序的思路,先进行分区,然后只需要递归一个分支
// 时间复杂度为O(n)
func findKthLargest(nums []int, k int) int {
	size := len(nums)
	if size < k {
		return -1
	}

	pivot := nums[0]
	head := 0
	tail := size - 1
	//把整个数组扫一遍，比pivot大的放左边，比pivot小的放右边
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

	//索引为head的元素即第k大
	if head+1 == k {
		return nums[head]
	}

	if head+1 < k {
		return findKthLargest(nums[head+1:], k-head-1)
	}
	return findKthLargest(nums[:head], k)
}
