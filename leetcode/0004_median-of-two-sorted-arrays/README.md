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

## 分治法
O(log(m+n))
```
a. 我们重新看题目，要找中位数，就是要找第 k 大的数（k = (L / 2 + 1)，其中L 是上面提到的合并后新数组的长度，当 L 是偶数时，要求第 (L / 2) 大和第 (L / 2 + 1) 大的两个数）。当我们舍弃掉一部分，假设舍弃部分的长度为 length，那么接下来就是在剩下的数组里求第 (k - length) 大的数。逐层缩小范围，直到两数组其中一个走完，或者要求的是第 1 大的元素，就可以直接返回结果了。

b. 那如何“选择”要舍弃哪部分呢？既然是要找合并后的数组 C 的第 k 大元素，即 C[k-1]，那如果我们从 A 和 B 中分别取前 k/2 个元素，其中必然有一部分是是在数组 C 的前 k 个数里。设 mid = k / 2，当 A[mid - 1] < B[mid - 1] 时，可以断定 A 的前 mid 个元素是在 C 的前 k 个数里（此处可用反证法得证），那么我们则舍弃 A 的前 mid 个元素。反之则舍弃 B 的前 mid 个元素。现在数组 A 或者 B 已经舍弃掉 k/2 个元素，缩小查找范围了，那接下来可以按照同样的方法继续选择吗？当然！现在剩下总共 (L - mid) 个元素，且 A 和 B 依旧有序，要找的是第 (k - mid) 大的元素，所以我们可以按照上面的方法继续递归选择下去，直到找到目标元素！

c. 复杂度分析：每次从合并后数组 C 里减少 k/2 个元素，直到找到目标元素。所以时间复杂度是 O(log L) = O(log (m + n)) ！
```

## 分治法升级版
O(log(min(m,n)))
[LeetCode官方解答](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/)

## 参考资料
[UBER 面试题 两个排序数组的的中位数](https://www.jiuzhang.com/qa/1768/)
