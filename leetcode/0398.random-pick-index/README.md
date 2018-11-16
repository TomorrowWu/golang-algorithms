
### 题目描述
给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

**注意:**

数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

**示例:**
```
int[] nums = new int[] {1,2,3,3,3};
Solution solution = new Solution(nums);

// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
solution.pick(3);

// pick(1) 应该返回 0。因为只有nums[0]等于1。
solution.pick(1);
```

### 代码实现
```
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
	res := []int{}
	for i, v := range s.nums {
		if v == target {
			res = append(res, i)
		}
	}

	rand.Seed(time.Now().UnixNano())
	return res[rand.Intn(len(res))]
}
```

### 思路
- 数组大小可能非常大，尽量不使用太多额外空间，所以Solution对象直接用nums字段
- 尝试过在Constructor()只遍历一遍nums，创建map存储nums中每个值的索引，但是Leetcode测试一直时间超过限制，可能是因为使用了额外内存吧
- Pick（）中遍历nums，找到target的索引到res中，rand随机选一个索引，返回


### GitHub
- [戳这里看源码](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0398.random-pick-index/src/random-pick-index.go)
- 各种数据结构及算法的Golang实现, LeetCode解题思路及答案
- [大厂面试题汇总及答案解析](https://github.com/TomorrowWu/interview)

###### 参考资料
[leetcode 398. 随机数索引](https://leetcode-cn.com/problems/random-pick-index/description/)

> 本文为原创文章，转载注明出处，欢迎扫码关注公众号 ```tomorrowwu``` 或者网站[https://lovecoding.club](https://lovecoding.club),第一时间看后续精彩文章，觉得好的话，顺手分享到朋友圈吧，感谢支持。

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/200)

