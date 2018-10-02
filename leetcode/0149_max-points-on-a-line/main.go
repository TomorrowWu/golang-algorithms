package main

/**
三点在同一直线上,任意两点之间的斜率相等
*/

//Point Definition for a point.
type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	n := len(points)
	// diffMap 用来过滤掉相同的点，并记录他们的个数
	//map特点: point1和point2,不是同一个对象,但是X和Y属性值相同,则key相等
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
	//高中数学中,判断第三个点在某条线上
	//(p3.Y-p1.Y)/(p3.X-p1.X) == (p2.Y-p1.Y)/(p2.X-p1.X),则p3在p1-p2线上
	return (p3.X-p1.X)*(p2.Y-p1.Y) == (p2.X-p1.X)*(p3.Y-p1.Y)
}

func main() {

}
