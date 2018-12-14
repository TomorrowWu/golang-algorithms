package src

import "math"

//质数有如下特点
//1,如果是合数，必然存在两个约数p1和p2,p1<=sqrt(n),p2>=sqrt(n),所以判断是否质数，只要判断 2--sqrt(n)

// isPrime 是否质数（素数）
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	//如果是合数，必然存在两个约数p1和p2,p1<=sqrt(n),p2>=sqrt(n),所以判断是否质数，只要判断 2--sqrt(n)
	sqrt := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}
