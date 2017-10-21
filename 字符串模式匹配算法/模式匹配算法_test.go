package 字符串模式匹配算法

import (
	"fmt"
	"testing"
)

/**
 * WuMing 
 *2017/10/19 下午6:30
 *
 */

func Test_KMPSearch(t *testing.T) {
	source := "ababab"
	pattern := "ba"
	fmt.Printf("KMP: %d\n", KMPSearch(source, pattern))
}