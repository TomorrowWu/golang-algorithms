[0149_直线上最多的点数](https://leetcode-cn.com/problems/max-points-on-a-line/description/)

## 题目描述
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上

示例1:
```
输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
```

示例2:
```
输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
```

## 算法

```Golang
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
```

## 个人思路

```
1. 任意不同的两点确定一条直线,遍历所有点是否在该线上,并计数
2. []Point中有重复的点,所以首先得利用map去重并分别记录点的个数,可以大大减少多层for循环中判断三点是否同一直线的次数,
3. 不同点不超过2个,则所有点在同一条线上
```

## 总结
- 笔者前期采用的思路是,公式 y=ax+b,斜率a相等,且b相等,则是同一直线,然后以(a,b)为key,只需遍历一遍[]Point,一层for循环,看似简单,但是写到后面,发现代码非常的复杂,且存在a为0或b为0等各种特殊情况,数据结构复杂,a为float64,有精度问题
- 目前版本算法,直接使用3层for循环,看似复杂度很高,但是可以先去重,减少for次数,以及代码直白简单,易懂的特点
- 好的代码,应该思路清晰,读者容易看懂,与君共勉

## GitHub
- [项目源码在这里](https://github.com/TomorrowWu/dataStructures-algorithm/tree/master/leetcode)
- 笔者会一直维护该项目,对leetcode中的算法题进行解决,并写下自己的思路和见解,致力于人人都能看懂的算法

## 个人公众号
- 喜欢的朋友可以关注,谢谢支持
- 记录90后码农的学习与生活

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
