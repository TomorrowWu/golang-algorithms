package _色球问题

import (
	"log"
)

/**
 * WuMing 
 *2017/9/19 下午8:07
 *
 */

/**
 吴名
 2017/9/19 下午下午8:08

三色球问题。有红、黄、绿三种颜色的球，其中红球3个，黄球3个，绿球6个。现将这12个球混放在一个盒子里，从中任意摸出8个球，编程计算摸出球的各种颜色搭配。
*/
func computeMatch(redCount, yellowCount, greenCount, total int) int {
	if total <= 0 || total > redCount+yellowCount+greenCount {
		return 0
	}
	num := 0
	for r := 0; r <= redCount; r++ {
		for y := 0; y <= yellowCount; y++ {
			for g := 0; g <= greenCount; g++ {
				if r+y+g == total {
					log.Printf("red:%v yellow:%v green:%v", r, y, g)
					num++
				}
			}
		}
	}
	return num
}
