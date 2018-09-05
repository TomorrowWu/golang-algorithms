package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/sum-of-subsequence-widths/description/

/*
给定一个整数数组 A ，考虑 A 的所有非空子序列。
对于任意序列 S ，设 S 的宽度是 S 的最大元素和最小元素的差。
返回 A 的所有子序列的宽度之和。
由于答案可能非常大，请返回答案模 10^9+7。
*/

/**
 * https://leetcode.com/problems/sum-of-subsequence-widths/discuss/161267/C++Java1-line-Python-Sort-and-One-Pass
 * For A[i]:
 * There are i smaller numbers,
 * so there are 2 ^ i sequences in which A[i] is maximum.
 * we should do res += A[i] * (2 ^ i)
 * There are n - i - 1 bigger numbers,
 * so there are 2 ^ (n - i - 1) sequences in which A[i] is minimum.
 * we should do res -= A[i] * 2 ^ (n - i - 1)
 */

/**
A[i]作为最大值,有i个更小的数,2^i次最大值
A[i]作为最小值,有n-1-i 个比它大的数,2^(n-1-i)次最小值

所有子序列宽度之和=S1(最大值-最小值)+S2(最大值-最小值)+S3(最大值-最小值).....

所有子序列宽度之和=(最大值-最小值)*次数+.....
*/

const mod = 1e9 + 7

func sumSubseqWidths(a []int) int {
	//[3,2,4,1] 和 排序后的 [1,2,3,4] 宽度之和相同

	sort.Ints(a)
	n := len(a)
	res := 0
	/**
	作为最大值出现的次数
	a[0] a[1] a[2]
	1	2	4

	[2,1,3],3作为最大值进行排列组合
	[2,3],[1,3][,2,1,3],[3]
	*/
	times := 1
	for i := 0; i < n; i++ {
		//a[i]作为最大值出现的次数==a[n-1-i]作为最小值出现的次数
		res += (a[i] - a[n-1-i]) * times
		//res可能非常大,所以取模
		res %= mod
		//times可能非常大,取模
		times = (times << 1) % mod
	}
	return res
}

func main() {
	fmt.Printf("sum:%d\n", sumSubseqWidths([]int{2, 1, 3}))    //6
	fmt.Printf("sum:%d\n", sumSubseqWidths([]int{3, 7, 2, 3})) //35
}
