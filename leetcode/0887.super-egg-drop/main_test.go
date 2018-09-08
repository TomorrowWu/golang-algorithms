package main

import "testing"

func Test_superEggDrop(t *testing.T) {
	t.Logf("最小移动次数:%d\n", superEggDrop(1, 2))  //2
	t.Logf("最小移动次数:%d\n", superEggDrop(2, 6))  //3
	t.Logf("最小移动次数:%d\n", superEggDrop(3, 14)) //4
	t.Logf("最小移动次数:%d\n", superEggDrop(1, 3))  //3
}
