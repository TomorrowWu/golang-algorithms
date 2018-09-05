package main

import "testing"

func Test_inserSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	insertSort(arr)
	t.Log("arr=", arr)
}

func Test_quickSort(t *testing.T) {
	//arr := []int{5, 4, 3, 2, 1}
	//arr := []int{5,4,5,2}
	arr := []int{1, 3, 2}
	quickSort(arr)
	t.Log("arr=", arr)
}

func Test_findMedianSortedArrays(t *testing.T) {
	median1 := findMedianSortedArrays([]int{1, 3}, []int{2})
	median2 := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	t.Logf("median1=%f\n", median1)
	t.Logf("median2=%f\n", median2)
}
