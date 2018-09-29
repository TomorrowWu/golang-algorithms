package 大话数据结构

/**
 * WuMing 
 *2017/8/31 下午8:34
 *斐波那契额数列实现
 */

//版本1:迭代版本
func Fbi1(i int) int {

	switch i {
	case 0:
		return 0
	case 1:
		return 1
	default:
		arr := make([]int, i+1, i+1)
		arr[0] = 0
		arr[1] = 1
		for j := 2; j <= i; j++ {
			arr[j] = arr[j-1] + arr[j-2]
		}
		return arr[i]
	}
}

//版本2:递归版本
func Fbi2(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return 1
	default:
		//递归调用
		return Fbi2(i-1) + Fbi2(i-2)
	}
}
