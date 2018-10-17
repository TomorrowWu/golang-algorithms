package insertionsort

// InsertSort 对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func InsertSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > value {
				array[j+1] = array[j] //数据移动
			} else {
				break
			}
		}
		array[j+1] = value //插入数据
	}
}
