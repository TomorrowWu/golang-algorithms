package main

import "fmt"

func SelectSort(arr *[5]int) {
	//(*arr)[1] = 600 等价于 arr[1] = 600,编译器会加*
	//(*arr)[1] = 600

	//1,先完成将第一个大值和arr[0] ==>先易后难

	//1  假设 arr[0]是最大值

	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j

		//2,遍历后面 1---[len(arr)-1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}

		//交换
		if maxIndex != j {
			//Go语言底层再使用中间变量,进行交换
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("第 %d 次 arr=%v \n", j, arr)
	}

}

func main() {
	//定义一个数组,从大到小
	arr := [5]int{10, 34, 19, 100, 80}
	//SelectSort(&arr) //从大到小
	SelectSort2(&arr) //从小到大
}

func SelectSort2(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		minIndex := j
		min := arr[j]

		for i := j + 1; i < len(arr); i++ {
			if arr[i] < min {
				minIndex = i
				min = arr[i]
			}
		}
		if minIndex != j {
			//交换
			arr[minIndex], arr[j] = arr[j], arr[minIndex]
		}
		fmt.Printf("第%d次,arr=%v \n", j, arr)
	}
}
