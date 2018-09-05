package main

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
//
//请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
//
//你可以假设 nums1 和 nums2 均不为空。

//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//中位数是 (2 + 3)/2 = 2.5

//插入排序
func insertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//快速排序
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	arrLen := len(arr)
	head, tail := 0, arrLen-1
	for i := 1; i <= tail; {
		if arr[i] > pivot {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			//head即pivot的索引,始终i-head=1
			head++
			i++
		}
	}
	quickSort(arr[:head])
	quickSort(arr[head+1:])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	memian := 0.0
	//合并数组
	arr := make([]int, 0, len(nums1)+len(nums2))
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)

	//排序(部分有序时,用什么排序最快)
	//insertSort(arr)
	quickSort(arr)
	arrLen := len(arr)
	index := arrLen / 2
	if arrLen%2 == 0 {
		//偶数长度
		memian = float64(arr[index-1]+arr[index]) / 2
	} else {
		memian = float64(arr[index])
	}
	return memian
}

func main() {

}
