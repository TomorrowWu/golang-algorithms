package main

import "fmt"

//func InsertSort(a *[]int) {
//	//完成第一次,给第二个元素找到合适的位置并插入
//
//	arr := *a
//
//	for j := 1; j < len(arr); j++ {
//		insertVal := arr[j]
//		insertIndex := j - 1
//
//		//从大到小
//		for insertIndex >= 0 && arr[insertIndex] < insertVal {
//			arr[insertIndex+1] = arr[insertIndex] //数据后移
//			insertIndex--
//		}
//
//		//插入
//		if insertIndex+1 != j {
//			arr[insertIndex+1] = insertVal
//		}
//
//		fmt.Printf("第%d次插入后arr=%v \n", j, arr)
//	}
//
//}

//对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func InsertSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			} else {
				break
			}
		}
	}
}

func main() {
	arr := []int{23, 0, 12, 56, 34}
	fmt.Printf("原始数组arr=%v \n", arr)
	InsertSort(arr)

	fmt.Printf("排序后数组arr=%v \n", arr)
}
