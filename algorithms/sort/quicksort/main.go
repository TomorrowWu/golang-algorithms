package main

import "fmt"

func QuickSort(left, right int, arr []int) {
	l := left
	r := right

	//数组的中位数
	pivot := arr[(left+right)/2]

	//比pivot小的放左边,比pivot大的放右边
	for l < r {
		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l >= r {
			break
		}

		//交换
		arr[l], arr[r] = arr[r], arr[l]

		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}

	//如果l==r,再移动下
	if l == r {
		l++
		r--
	}

	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}

	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

func main() {
	//arr := []int{100, 99, 98, 0, -98, -99, -100}
	arr := []int{-100, -99, -98, 0, 98, 99, 100, 0, 1, -100, 98}

	fmt.Println("原始数组arr=", arr)

	//QuickSort(0, len(arr)-1, arr)
	qsort(arr)
	fmt.Println("快速排序后arr=", arr)
}
