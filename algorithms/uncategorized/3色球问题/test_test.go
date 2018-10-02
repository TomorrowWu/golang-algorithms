package _色球问题

import (
	"testing"
	"time"
)

/**
 * WuMing
 *2017/9/19 下午8:16
 *
 */

func Test_computeMatch(t *testing.T) {
	//n := computeMatch(1,1,1,2)
	//n := computeMatch(3,3,6,8)
	//n := computeMatch(3,3,6,2)
	t1 := time.Now()
	n := computeMatch(3000, 3000, 6001, 12000)
	duration := time.Since(t1)
	t.Logf("排列组合数量:%v 时长:%v", n, duration)
}

func Test_computeMatch1(t *testing.T) {
	t1 := time.Now()
	n := computeMatch1(30, 30, 31, 90)
	duration := time.Since(t1)
	t.Logf("排列组合数量:%v 时长:%v", n, duration)
}
