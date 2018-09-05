package main

import "fmt"

//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
//如果没有和至少为 K 的非空子数组，返回 -1 。

//输入：A = [2,-1,2], K = 3
//输出：3

//shortestSubarray  暴力破解法(时间复杂度O(n^2))
func shortestSubarray(A []int, K int) int {

	minLen := 50001

	tail := len(A) - 1
	for i := 0; i <= tail; i++ {
		sum := A[i]
		if sum >= K {
			minLen = 1
			break
		}
		for j := i + 1; j <= tail; j++ {
			sum += A[j]
			if sum >= K {
				l := j - i + 1
				if l < minLen {
					minLen = l
				}
				break
			}
		}
	}

	if minLen != 50001 {
		return minLen
	}

	return -1
}

//GitHub 某MIT研究生的代码(本人不才,没看懂)
func shortestSubarray2(A []int, K int) int {
	size := len(A)

	//sums[0]=0  1...size为前n位的和
	sums := make([]int, size+1)
	for i, n := range A {
		if n >= K {
			return 1
		}
		sums[i+1] = sums[i] + n
	}

	res := size + 1

	//双端队列
	deque := make([]int, size+2)
	l, r := 0, 0

	for i := 1; i < size+1; i++ {
		for r-l >= 0 && sums[i]-sums[deque[l]] >= K {
			// 由于 i 递增
			// 也许以后会有 j>i 使得 sums[j] - sums[deque[l]] >= K
			// 但是由于 j>i, j-deque[l] > i-deque[l] 不会是更短的答案。
			// 所以，可以把 deque[l] 删除
			res = min(res, i-deque[l])
			l++
		}
		for r-l >= 0 && sums[i] <= sums[deque[r]] {
			// 如果存在 j>i>deque[r] 使得 sums[j] - sums[deque[r]] >= K
			// 由于 sums[deque[r]] >= sums[i] 则
			// sums[j] - sums[i] >= K 一定成立，并且 j-i < j-deque[r]
			// 所以，没有必要把 deque[r] 放入队列中
			r--
		}
		r++
		deque[r] = i
	}

	if res == size+1 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//思路: https://blog.csdn.net/Fmuma/article/details/79930981
//shortestSubarray3 滑动窗口解法(正整数数组,有负数的,解决不了)
func shortestSubarray3(A []int, K int) int {
	size := len(A)
	if size == 0 {
		return -1
	}

	//构建滑动窗口,从arr[0]开始
	left := 0
	right := -1 //开始时,假设滑动窗口不存在
	sum := 0    //滑动窗口元素的和

	result := size + 1 //开始时,滑动窗口的大小为size+1,不可能是这个值

	for left < size {
		if right+1 < size && sum < K {
			//sum小于K,窗口右边界扩大
			right++
			sum += A[right]
		} else {
			//sum>=k,窗口从左边界缩小
			sum -= A[left]
			left++
		}
		if sum >= K {
			if right-left+1 < result {
				result = right - left + 1
			}
		}
	}

	if result == size+1 {
		return -1
	}

	return result

}

//和shortestSubarray2 的算法思路是一样的,只是deque双端队列更直白,容易理解
func shortestSubarray5(A []int, K int) int {
	size := len(A)
	sums := make([]int, size+1)
	for i, n := range A {
		if n >= K {
			return 1
		}
		sums[i+1] = sums[i] + n
	}

	res := size + 1
	//存储0----i,有可能是符合条件的最短子串的head
	deque := make([]int, 0, 0)
	for i := 0; i < size+1; i++ {
		for len(deque) > 0 && (sums[i]-sums[deque[0]] >= K) {
			// i 递增
			// 可能有 j>i 使得 sums[j] - sums[deque[0]] >= K
			// 但是由于 j>i,所以deque[0]---i是以deque[0]作为头的最短的符合条件的子串
			// 把 deque[0] 删除
			res = min(res, i-deque[0])
			deque = deque[1:]
		}
		for len(deque) > 0 && (sums[i] <= sums[deque[len(deque)-1]]) {
			// 1,如果存在 j>i>deque[r] 使得 sums[j] - sums[deque[r]] >= K
			// 由于 sums[deque[r]] >= sums[i] 则
			// sums[j] - sums[i] >= K 一定成立，并且 j-i < j-deque[r]
			// 所以，以deque[r]作为头,不可能是最短的符合条件子串,删除

			//根据debug过程中数据变化,如果queue中尾是负数,这里会去掉
			//如果r-i之间元素的和是负数,那么已r为头,满足1条件,不可能是最短的子串
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	if res == size+1 {
		return -1
	}

	return res
}

func main() {
	fmt.Printf("subarrayLen=%d\n", shortestSubarray5([]int{2, -1, 2}, 3))
	fmt.Printf("subarrayLen=%d\n", shortestSubarray5([]int{-28, 81, -20, 28, -29}, 89))
	fmt.Printf("subarrayLen=%d\n", shortestSubarray5([]int{45, 95, 97, -34, -42}, 21))
	fmt.Printf("subarrayLen=%d", shortestSubarray5([]int{-34, 37, 151, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6}, 151))
}
