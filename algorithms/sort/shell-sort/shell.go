package shell

func shellSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for key := n / 2; key > 0; key = key / 2 {
		for i := key; i < n; i++ {
			for j := i; j >= key && a[j] < a[j-key]; j -= key {
				a[j], a[j-key] = a[j-key], a[j]
			}
		}
	}
}
