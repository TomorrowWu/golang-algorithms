package selectsort

// SelectSort 选择排序.
// 时间复杂度O(n^2).
// 不稳定排序.(比如[5,5,2])
func SelectSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}

	for i := 0; i < arrLen; i++ {
		minIndex := i
		for j := i + 1; j < arrLen; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			//交换
			a[i], a[minIndex] = a[minIndex], a[i]
		}
	}
}
