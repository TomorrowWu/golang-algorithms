package main

func superEggDrop(K, N int) int {
	moves := 0
	dp := make([]int, K+1) // 1 <= K <= 100
	// dp[i] = n 表示， i 个鸡蛋，利用 moves 次移动，最多可以检测 n 层楼
	for dp[K] < N {
		for i := K; i > 0; i-- {
			//逆序从K---1,dp[i] = dp[i]+dp[i-1] + 1 相当于上次移动后的结果,dp[]函数要理解成抽象出来的一个黑箱子函数,跟上一次移动时鸡蛋的结果有关系
			dp[i] += dp[i-1] + 1
			// 以上计算式，是从以下转移方程简化而来
			// dp[moves][k] = 1 + dp[moves-1][k-1] + dp[moves-1][k]
			// 假设 dp[moves-1][k-1] = n0, dp[moves-1][k] = n1
			// 首先检测，从第 n0+1 楼丢下鸡蛋会不会破。
			// 如果鸡蛋破了，F 一定是在 [1：n0] 楼中，
			// 		利用剩下的 moves-1 次机会和 k-1 个鸡蛋，可以把 F 找出来。
			// 如果鸡蛋没破，假如 F 在 [n0+2:n0+n1+1] 楼中
			// 		利用剩下的 moves-1 次机会和 k 个鸡蛋把，也可以把 F 找出来。
			// 所以，当有 moves 个放置机会和 k 个鸡蛋的时候
			// F 在 [1, n0+n1+1] 中的任何一楼，都能够被检测出来。
		}
		moves++
	}
	return moves
}

/**
我们可以把M层楼 / N个鸡蛋的问题转化成一个函数 F（M，N），其中楼层数M和鸡蛋数N是函数的两个参数，而函数的值则是最优解的最大尝试次数。
假设我们第一个鸡蛋扔出的位置在第X层（1<=X<=M），会出现两种情况：
1.第一个鸡蛋没碎
那么剩余的M-X层楼，剩余N个鸡蛋，可以转变为下面的函数：
 F（M-X，N）+ 1，1<=X<=M

2.第一个鸡蛋碎了
那么只剩下从1层到X-1层楼需要尝试，剩余的鸡蛋数量是N-1，可以转变为下面的函数：
F（X-1，N-1） + 1，1<=X<=M

整体而言，我们要求出的是 M层楼 / N个鸡蛋 条件下，最大尝试次数最小的解，所以这个题目的状态转移方程式如下：
 F（M，N）= Min（Max（ F（M-X，N）+ 1， F（X-1，N-1） + 1）），1<=X<=M
*/

//https://www.itcodemonkey.com/article/4915.html
//https://www.itcodemonkey.com/article/5279.html

//superEggDrop2 动态规划原始版(超出leetcode时间限制)
//
//三层循环,时间复杂度是O(M*M*N)
//
//二维数组:空间复杂度是O(M*N)
func superEggDrop2(K, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	//备忘录，存储K个鸡蛋，N层楼条件下的最优化尝试次数
	//cache := [K + 1][N + 1]int{}
	cache := make([][]int, K+1)
	//把备忘录每个元素初始化成最大的尝试次数
	for i := 0; i <= K; i++ {
		cache[i] = make([]int, N+1)
		for j := 1; j <= N; j++ {
			cache[i][j] = j
		}
	}
	for n := 2; n <= K; n++ {
		for m := 1; m <= N; m++ {
			//假设楼层数可以是1---N,
			min := cache[n][m]
			for k := 1; k < m; k++ {
				//M层,N鸡蛋,F（M，N）= Min（Max（ F（M-X，N）+ 1， F（X-1，N-1） + 1））
				//(动态规划)
				//鸡蛋碎了
				max := cache[n-1][k-1] + 1
				if cache[n][m-k]+1 > max {
					max = cache[n][m-k] + 1 //鸡蛋没碎
				}
				if max < min {
					min = max
				}
			}
			cache[n][m] = min
		}
	}
	return cache[K][N]
}

func main() {

}
