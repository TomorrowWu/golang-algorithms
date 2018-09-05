# [两个排序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/)

## 题目描述
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

你可以假设 nums1 和 nums2 均不为空。

示例1：
```
nums1 = [1, 3]
nums2 = [2]

中位数是 2.0
```

示例2：
```
nums1 = [1, 2]
nums2 = [3, 4]

中位数是 (2 + 3)/2 = 2.5
```

## [官方解答](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/)

## 个人思路
- 对两个数组合并,并进行排序,再取中位数
- 难度在于:要求算法的时间复杂度为 O(log (m+n))
- 算法的时间主要是在排序上,普通的排序算法,时间复杂度较高,此处是两个有序数组,从这个点入手,降低时间复杂度

