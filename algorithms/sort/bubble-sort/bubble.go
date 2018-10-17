package bubble

func bubbleSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}

	for i := 0; i < arrLen; i++ {
		//提前退出标志位(没有数据交换,说明已经全部有序)
		flag := false
		for j := 0; j < arrLen-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
