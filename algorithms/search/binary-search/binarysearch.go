package binarysearch

func binarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	for low, high := 0, n-1; low <= high; {
		mid := low + (high-low)>>1
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// varient 1: Finds the element who is first value that equals the given value
func binarySearchFirstEqual(a []int, value int) int {
	n := len(a)
	for low, high := 0, n-1; low <= high; {
		mid := low + (high-low)>>1
		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || a[mid-1] != value {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// varient 2: Finds the element who is last value that equals the given value
func binarySearchLastEqual(a []int, value int) int {
	n := len(a)
	for low, high := 0, n-1; low <= high; {
		mid := low + (high-low)>>1
		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == n-1 || a[mid+1] != value {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// varient 3: Finds the element who is first value that large or equals the given value
func binarySearchFirstLargeEqual(a []int, value int) int {
	n := len(a)
	for low, high := 0, n-1; low <= high; {
		mid := low + (high-low)>>1
		if a[mid] >= value {
			if mid == 0 || a[mid-1] < value {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// varient 4: Finds the element who is last value that small or equals the given value
func binarySearchLastSmallEqual(a []int, value int) int {
	n := len(a)
	for low, high := 0, n-1; low <= high; {
		mid := low + (high-low)>>1
		if a[mid] > value {
			high = mid - 1
		} else {
			if mid == n-1 || a[mid+1] > value {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
