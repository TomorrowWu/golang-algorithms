package main

import (
	"testing"
)

func Test_twoSum(t *testing.T)  {
	nums := []int{3,2,4}
	target := 6
	result1 := twoSum(nums, target)
	t.Log("result1=",result1)
}

func Test_twoSum2(t *testing.T)  {
	nums := []int{3,2,4}
	target := 6
	result1 := twoSum2(nums, target)
	t.Log("result2=",result1)
}

func Test_twoSum3(t *testing.T)  {
	nums := []int{3,2,4}
	target := 6
	result1 := twoSum3(nums, target)
	t.Log("result3=",result1)
}
