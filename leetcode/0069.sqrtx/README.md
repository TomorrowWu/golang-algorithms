### 题目描述
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例1:
```
输入: 4
输出: 2
```

示例2:
```
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。
```

### 暴力版本
```Golang
// 暴力解法
func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		res := i * i
		if res == x {
			return i
		} else if res > x {
			return i - 1
		}
	}
	return -1
}
```

### 二分查找
```Golang
//在有序数组中，找到最后一个小于等于给定值的数
func mySqrt2(x int) int {
	low, high := 0, x
	for low <= high {
		//防止大数相加溢出
		//位运算更高效
		mid := low + (high-low)>>1
		product := mid * mid
		if product > x {
			high = mid - 1
		} else {
			if (mid == x) || (mid+1)*(mid+1) > x {
				//遍历最后一个数 || 下一个数大于目标值
				return mid
			}
			//下一个数小于等于目标值，所以mid不是最后一个数
			low = mid + 1
		}
	}
	return -1
}
```

### 二分查找思路
- 相当于从0-x中找到最后一个平方<=x的整数
- 我们求解的是最后一个小于等于给定值的元素，所以当 product<=x时，需要确认 mid+1 的平方>x
- 如果 mid+1 的平方 <= x ,说明mid肯定不是最后一个，更新low

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0069.sqrtx/src/sqrtx.go)
- 各种数据结构及算法的Golang实现, LeetCode解题思路及答案
- [大厂面试题汇总及答案解析](https://github.com/TomorrowWu/interview)

###### 参考资料
[leetcode 69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/description/)

[极客时间 数据结构与算法之美](https://time.geekbang.org/column/article/42733)

> 本文为原创文章，转载注明出处，欢迎扫码关注公众号 ```tomorrowwu``` 或者网站[https://lovecoding.club](https://lovecoding.club),第一时间看后续精彩文章，觉得好的话，顺手分享到朋友圈吧，感谢支持。

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/200)
