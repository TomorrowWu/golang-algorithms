package src

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

//归并排序思路（O(min(m,n)))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	res := make([]int, 0, l)
	//归并排序
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i >= len(nums1) {
		res = append(res, nums2[j:]...)
	}
	if j >= len(nums2) {
		res = append(res, nums1[i:]...)
	}

	if l%2 == 0 {
		return float64(res[l/2]+res[l/2-1]) / 2.0
	}
	return float64(res[l/2])
}

//findMedianSortedArrays2 给出两个有序数组，假设两个数组的长度和是 len，如果 len 为奇数，那么我们求的就是两个数组合并后的第 (len >> 1) + 1 大的数，如果 len 为偶数，就是第 (len >> 1) 和 (len >> 1) + 1 两个数的平均数
//  给定两个有序数组，求第k大数,如果我们从 A 和 B 中分别取前 k/2 个元素，其中必然有一部分是是在数组 C 的前 k 个数里
//  设 mid = k / 2，当 A[mid - 1] < B[mid - 1] 时，可以断定 A 的前 mid 个元素是在 C 的前 k 个数里
//  时间复杂度是 O(log(m+n))
//TODO 吴名 2018-12-10 13:33 分治法

// leetcode官方解答，时间复杂度 O(log(min(m,n)))
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	//num1是较短数组
	//
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	if n == 0 {
		panic("ValueError")
	}

	imin, imax, half_len := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := half_len - i
		if i < m && nums2[j-1] > nums1[i] {
			// i is too small, must increase it
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i is too big, must decrease it
			imax = i - 1
		} else {
			// i is perfect

			max_of_left := 0
			if i == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[i-1]
			} else {
				max_of_left = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(max_of_left)
			}

			min_of_right := 0
			if i == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[i]
			} else {
				min_of_right = min(nums1[i], nums2[j])
			}

			return float64(max_of_left+min_of_right) / 2
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
