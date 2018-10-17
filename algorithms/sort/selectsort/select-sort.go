package selectsort

// SelectSort 选择排序
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
