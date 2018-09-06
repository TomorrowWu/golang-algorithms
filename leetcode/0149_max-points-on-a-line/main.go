package main

/**
三点在同一直线上,任意两点之间的斜率相等
*/

//Point Definition for a point.
type Point struct {
	X int
	Y int
}

//maxPoints 错误的算法,始终无法完全通过测试用例,不过可以参考思路
func maxPoints(points []Point) int {
	size := len(points)
	//size边界条件
	if size == 0 {
		return 0
	}

	if size == 1 {
		return 1
	}

	max := 0

	//分子分母都不为0
	slopes := make(map[float64]map[float64]map[int]bool)

	//分子为0,分母不为0,与X轴平行,y值相等
	yMap := make(map[int]map[int]bool)
	//分子不为0,分母为0,与Y轴平行
	xMap := make(map[int]map[int]bool)
	//分子和分母都为0,两点相同
	xyMap := make(map[int]map[int]map[int]bool)

	for i, _ := range points {
		for j := i + 1; j < size; j++ {
			//分子
			diff1 := points[j].Y - points[i].Y
			//分母
			diff2 := points[j].X - points[i].X

			//分子或分母都不为0
			if diff1 != 0 && diff2 != 0 {
				//直线:y=ax+b
				//斜率相等,b相等,才是同一条直线

				//斜率
				slope := float64(diff1) / float64(diff2)

				//求b
				b := float64(points[j].Y) - slope*float64(points[j].X)

				if _, ok := slopes[slope]; !ok {
					slopes[slope] = map[float64]map[int]bool{}
				}

				if _, ok := slopes[slope][b]; !ok {
					slopes[slope][b] = map[int]bool{}
				}

				slopes[slope][b][i] = true
				slopes[slope][b][j] = true

				count := len(slopes[slope][b])
				if count > max {
					max = count
				}
			}

			//两点相同,坐标相同点,xyMap中key唯一
			if diff1 == 0 && diff2 == 0 {
				if _, ok := xyMap[points[j].X]; !ok {
					xyMap[points[j].X] = make(map[int]map[int]bool)
				}
				if _, ok := xyMap[points[j].X][points[j].Y]; !ok {
					xyMap[points[j].X][points[j].Y] = make(map[int]bool)
				}
				xyMap[points[j].X][points[j].Y][i] = true
				xyMap[points[j].X][points[j].Y][j] = true

				count := len(xyMap[points[j].X][points[j].Y])
				if count > max {
					max = count
				}
			}

			//与X轴平行
			if diff1 == 0 && diff2 != 0 {
				if _, ok := yMap[points[j].Y]; !ok {
					yMap[points[j].Y] = map[int]bool{}
				}

				yMap[points[j].Y][i] = true
				yMap[points[j].Y][j] = true

				count := len(yMap[points[j].Y])
				if count > max {
					max = count
				}
			}

			//与Y轴平行
			if diff1 != 0 && diff2 == 0 {
				if _, ok := xMap[points[j].X]; !ok {
					xMap[points[j].X] = map[int]bool{}
				}
				xMap[points[j].X][i] = true
				xMap[points[j].X][j] = true

				count := len(xMap[points[j].X])
				if count > max {
					max = count
				}
			}

		}
	}

	return max
}

func maxPoints2(points []Point) int {
	n := len(points)
	// diffMap 用来过滤掉相同的点，并记录他们的个数
	diffMap := make(map[Point]int, n)

	for i := 0; i < n; i++ {
		diffMap[points[i]]++
	}

	size := len(diffMap)

	// 不超过 2 个不同的点
	// 则，所有的点都在同一条直线上
	if size <= 2 {
		return n
	}

	max := 0
	// 存在相同的点，
	// 则，提取所有不同的点，可以大大减少后面 3 个嵌套的 for 循环的次数
	if size < n {
		points = make([]Point, 0, size)
		for p := range diffMap {
			points = append(points, p)
		}
	}

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			count := 0
			// 所有的点，都要检查，是否与 i， j 共线
			for k := 0; k < size; k++ {
				if isSameLine(points[i], points[j], points[k]) {
					count += diffMap[points[k]]
				}
			}
			if max < count {
				max = count
			}
		}
	}

	return max
}

func isSameLine(p1, p2, p3 Point) bool {
	return (p3.X-p1.X)*(p2.Y-p1.Y) == (p2.X-p1.X)*(p3.Y-p1.Y)
}
