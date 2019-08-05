package main

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// 尾递归版本
func TailRecursion(N int) int {
	return fibTail(N, 0, 1)
}

func fibTail(N, num1, num2 int) int {
	if N == 0 {
		return num1
	}
	return fibTail(N-1, num2, num1+num2)
}
