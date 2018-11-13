package src

//“求一个数的平方根” 精确到小数点后6位

const d = 0.000001

//二分法查找
func mySqrt3(x int) float32 {
	target := float32(x)
	low, high := float32(0), target
	for low <= high {
		//防止大数溢出
		mid := low + (high-low)/2
		product := mid * mid
		diff := product - target
		if diff > d {
			// product > target
			high = mid
		} else {
			if diff >= 0 {
				//近似 product == target
				//符合要求,保留6位小数
				return float32(int(mid*1000000)) / 1000000
			}
			// product < target
			low = mid
		}
	}
	return -1
}
