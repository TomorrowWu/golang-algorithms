package src

import (
	"math/rand"
	"time"
)

//// Solution defines a structure
//type Solution struct {
//	m map[int][]int
//}
//
//// Constructor constructs a object
//func Constructor(nums []int) Solution {
//	m := make(map[int][]int, len(nums))
//	for i, v := range nums {
//		m[v] = append(m[v], i)
//	}
//	return Solution{
//		m: m,
//	}
//}
//
//// Pick returns index number that target at nums Randomly.
//func (s *Solution) Pick(target int) int {
//	m := s.m[target]
//	//rand.Seed(time.Now().UnixNano())
//	return m[rand.Intn(len(m))]
//}

// Solution defines a structure
type Solution struct {
	nums []int
}

// Constructor constructs a object
func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

// Pick returns index number that target at nums Randomly.
func (s *Solution) Pick(target int) int {
	var res []int
	for i, v := range s.nums {
		if v == target {
			res = append(res, i)
		}
	}

	rand.Seed(time.Now().UnixNano())
	return res[rand.Intn(len(res))]
}
