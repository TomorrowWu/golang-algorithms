package 大话数据结构

import (
	"testing"
	"log"
)

/**
 * WuMing 
 *2017/9/9 上午12:08
 *
 */

func Test_Fbi1(t *testing.T) {
	t.Logf("结果:%v",Fbi1(40))
}

func Test_Fbi2(t *testing.T) {
	t.Logf("结果:%v",Fbi2(40))
}

func Test_index(t *testing.T) {
	//pos := index("abcdef-google","google")
	pos := index("googlegood","google")
	log.Printf("pos: %v",pos)
}
