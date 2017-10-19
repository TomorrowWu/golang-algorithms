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
	source := "This is Rinc Liu's personal website."
	pattern := "Rinc Liu"
	fmt.Printf("KMP: %d\n", KMPSearch(source, pattern))
}