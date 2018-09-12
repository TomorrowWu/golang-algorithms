[0862_和至少为 K 的最短子数组](https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/description/)

## 题目描述
返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。

如果没有和至少为 K 的非空子数组，返回 -1 。

示例1:
```
输入：A = [1], K = 1
输出：1
```

示例2:
```
输入：A = [1,2], K = 4
输出：-1
```

示例3:
```
输入：A = [2,-1,2], K = 3
输出：3
```

Note:
```
1.  1 <= A.length <= 50000
2.  -10 ^ 5 <= A[i] <= 10 ^ 5
3.  1 <= K <= 10 ^ 9
```

## 暴力算法

```Golang
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

```


## 高性能版本算法

```Golang
func shortestSubarray(A []int, K int) int {
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
			// 如果存在 j>i>deque[r] 使得 sums[j] - sums[deque[r]] >= K
			// 由于 sums[deque[r]] >= sums[i] 则
			// sums[j] - sums[i] >= K 一定成立，并且 j-i < j-deque[r]
			// 所以，以deque[r]作为头,不可能是最短的符合条件子串,删除
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	if res == size+1 {
		return -1
	}

	return res
}
```

## 个人思路
1. 相对于暴力破解,重复计算两个元素之间的和,所以此处必须进行优化,只需要循环一遍数组,存储0-i元素的和
2. 假如a<b,sums[b]-sums[a]>=K,b-a则是子串长度
3. deque(双端队列)存储0-i,可能符合条件的最短子串的head
4. 遍历sums,取deque的[0],判断该元素作为子串的head,是否符合条件,如果符合条件,那么它将是该数组元素作为head的符合条件的子串中,最短的子串,从deque中剔除[0],减少后面的sums比次数
5. 取deque的尾,值为r,sums[r]>=sums[i],那么如果后面存在sums[j]-sums[r]>=k,那么j>i>r,i->j之间的子串肯定也符合条件,并且更短,已r为head的子串,不可能是最短的字段,在这里可以删除deque的这个尾,减少比对次数

## 总结
- 在leetcode社区中,有Java版本的算法代码,思路和此处思路完全一致,但是时间效率比Golang版的搞,Golang版170ms左右,Java版本40+ms
- 笔者经过对比代码,主要区别在于deque这个存储容器,在JDK中有对应的数据结构Deque双端队列,数据结构的插入效率导致的,笔者猜测Java中该结构用的双向链表实现的,在插入元素过程中,不需要扩容,内存分配效率很高
- Golang版本代码,deque用的slice切片,在append过程中,不断的扩容,进行内存分配,导致效率低
- 有兴趣的朋友可以在此处优化deque这个数据结构,采用struct对象,构建双向链表结构,相信,效率和Java版本的应该差不多,在此,篇幅限制,双向链表也不是该算法的核心,笔者就到此为止
- 数据结构作为数据的存储容易,对算法的影响也是巨大的

## GitHub
- [项目源码在这里](https://github.com/TomorrowWu/dataStructures-algorithm)
- 笔者会一直维护该项目,对leetcode中的算法题进行解决,并写下自己的思路和见解,致力于人人都能看懂的算法

## 个人公众号
- 喜欢的朋友可以关注,谢谢支持
- 记录90后码农的学习与生活

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
