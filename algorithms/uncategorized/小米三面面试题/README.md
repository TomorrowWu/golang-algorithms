> 一副从1到n的牌，每次从牌堆顶取一张放桌子上，再取一张放牌堆底，直到手上没牌，最后桌子上的牌是从1到n有序，设计程序，输入n，输出牌堆的顺序数组。

#### 正面角度

“取1张存1张”，说白了就是跳过一张取嘛。比如n=3时，考虑{牌1, 牌2, 牌3}，第一张取了“牌1”，那么第二张取的就是“牌3”。那只用给“牌1”标“1”、“牌3”标“2”、“牌2”标“3”就行了。换句话说，就是跳1位标数字。标过的数字就相当于桌子上的牌，没标的就相当于手上的牌。所以我们可以编写如下程序：
```Golang
func positiveSolution(n int) []int {
	result := make([]int, n)
	//next表示是否需要跳过这位数字
	next := false
	//i就是待标记的数字
	i := 1
	for i <= n {
		for p := 0; p < n; p++ {
			if result[p] == 0 {
				if !next {
					result[p] = i
					i++
				}
				next = !next
			}
		}
	}

	return result
}
```

###### 参考资料
[算法多解:小米三面面试题](https://toutiao.io/posts/i53sto/preview)
