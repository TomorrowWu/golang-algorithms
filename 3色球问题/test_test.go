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
	n := computeMatch(1,1,1,2)
	t.Logf("排列组合:%v",n)
}
