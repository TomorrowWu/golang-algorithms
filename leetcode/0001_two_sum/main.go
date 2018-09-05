package main

import "fmt"

//twoSum 暴力破解
func twoSum(nums []int, target int) []int {
	for i,num1 := range nums{
		for j:=i+1;j<=len(nums)-1;j++{
			if num1+nums[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}

//twoSum2 两遍哈希表
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for i,num :=range nums{
		m[num] = i
	}

	for i,num :=range nums{
		num2 := target-num
		//遍历nums第二遍,num和num2可能相等
		if j,ok := m[num2];ok && num2!=num{
			return []int{i,j}
		}
	}
	return nil
}

//twoSum3 一遍哈希表
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for j,num2 :=range nums{
		num1 := target-num2
		if i,ok := m[num1];ok{
			return []int{i,j}
		}else {
			m[num2] = j
		}
	}
	return nil
}

func main() {
	nums := []int{3,2,4}
	target := 6
	result1 := twoSum(nums, target)
	result2 := twoSum2(nums, target)
	result3 := twoSum3(nums, target)
	fmt.Println("result1=",result1)
	fmt.Println("result2=",result2)
	fmt.Println("result3=",result3)
}
