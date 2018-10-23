package src

// 计数排序，a是数组，n是数组大小。假设数组中存储的都是非负整数。
// 时间复杂度O(n),空间复杂度O(n)
func countingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 查找数组中数据的范围
	max := a[0]
	for _, val := range a {
		if max < val {
			max = val
		}
	}

	// 申请一个计数数组c，下标大小[0,max]
	c := make([]int, max+1)

	// 计算每个元素的个数，放入c中
	for _, val := range a {
		c[val]++
	}

	// 依次累加
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组r，存储排序之后的结果
	r := make([]int, n)
	// 计算排序的关键步骤了，有点难理解
	for i := n - 1; i >= 0; i-- {
		//从后往前遍历数组a,根据元素大小,从数组c中查自己排序后的索引
		//将元素写入数组r
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	// 将结果拷贝会a数组
	for i := range a {
		a[i] = r[i]
	}
}
