### 题目描述
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 ```[0,1,2,4,5,6,7]``` 可能变为 ```[4,5,6,7,0,1,2] )```。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 ```-1``` 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例1：
```
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
```

示例2：
```
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
```

### 代码实现
```
//1. 先二分遍历找到分隔点index，特征是 < 前一个元素, >后一个元素;
//2. 把数组分成二个部分，[0,index-1], [index,length-1];
//3. 分别使用二分查找，找到给定的值。
//时间复杂度为log(n). 不确定有什么更好的办法

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums) - 1
	divisionIndex := findDivision(nums)
	if divisionIndex == 0 || divisionIndex == -1 {
		//非旋转排序数组
		return findTarget(nums, 0, n, target)
	}

	res := findTarget(nums, 0, divisionIndex-1, target)
	if res != -1 {
		return res
	}
	return findTarget(nums, divisionIndex, n, target)
}

//找到分割点
func findDivision(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1

		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func findTarget(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
```

### 思路
- 先二分遍历找到分隔点index，特征是 < 前一个元素, >后一个元素;
- 把数组分成二个部分，[0,index-1], [index,length-1];
- 分别使用二分查找，找到给定的值。

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0033.search-in-rotated-sorted-array/src/search-in-rotated-sorted-array.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 参考资料
[leetcode 33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/)

> 本文为原创文章，转载注明出处，欢迎扫码关注公众号 ```tomorrowwu``` 或者网站[https://lovecoding.club](https://lovecoding.club),第一时间看后续精彩文章，觉得好的话，顺手分享到朋友圈吧，感谢支持。

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/200)

