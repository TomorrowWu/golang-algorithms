package _色球问题

import (
	"testing"
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
	n := computeMatch(3000,3000,6001,12000)
	t.Logf("排列组合:%v",n)
}
