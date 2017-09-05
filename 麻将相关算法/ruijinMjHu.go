package algorithm

import (
	"sort"
	"log"
)

/**
 * WuMing
 *2017/7/6 下午2:40
 *针对瑞金麻将的函数
 */

var (
	ruiJinmahjongArr = []int{
		101, 102, 103, 104, 105, 106, 107, 108, 109, //#万
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		201, 202, 203, 204, 205, 206, 207, 208, 209, //#饼
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		301, 302, 303, 304, 305, 306, 307, 308, 309, //#条
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		401, 402, 403, 404, 405, 406, 407, //# 东 西 南 北 中 发 白
		401, 402, 403, 404, 405, 406, 407,
		401, 402, 403, 404, 405, 406, 407,
		401, 402, 403, 404, 405, 406, 407,
		501, 502, 503, 504, //花(春夏秋冬)
	}
)

/**
吴名
2017/7/10 下午5:26
3n+2牌型的胡牌
*/
func isHU(arr []int, bao int) bool {
	mjArr := append([]int{}, arr...)

	//宝,万,筒,条,风
	sptArr := seperateRuiJinArr(mjArr, bao)

	baoNum := len(sptArr[0]) //手牌中宝的数量
	if baoNum == 4 || (bao > 500 && baoNum == 3) {
		//飞
		log.Println("胡:所有的宝牌都在一家,飞")
		return false
	}

	//检测牌数量
	if len(mjArr)%3 != 2 {
		//log.Println("牌数量不符合3n+2")
		return false
	}

	needBaoArr := [5]int{}
	for i := 1; i <= 4; i++ {
		a := append([]int{}, sptArr[i]...)
		needNum := getNeedBaoNumToZhengPu(a)
		needBaoArr[i] = needNum
	}

	// 吴名 2017/7/7 下午8:30 将在"万"中
	needBaoNum := needBaoArr[2] + needBaoArr[3] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "万"成整扑一将的癞子数量
		num := getBaoNumToZhengPuJiang(sptArr[1])
		if leftBaoNum >= num {
			return true
		}
	}

	// 吴名 2017/7/7 下午8:30 将在"筒"中
	needBaoNum = needBaoArr[1] + needBaoArr[3] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "筒"成整扑一将的癞子数量
		num := getBaoNumToZhengPuJiang(sptArr[2])
		if leftBaoNum >= num {
			return true
		}
	}

	// 吴名 2017/7/7 下午8:31 将在"条"中
	needBaoNum = needBaoArr[1] + needBaoArr[2] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "条"成整扑一将的癞子数量
		num := getBaoNumToZhengPuJiang(sptArr[3])
		if leftBaoNum >= num {
			return true
		}
	}
	// 吴名 2017/7/7 下午8:31 将在"风"中
	needBaoNum = needBaoArr[1] + needBaoArr[2] + needBaoArr[3]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "风"成整扑一将的癞子数量
		num := getBaoNumToZhengPuJiang(sptArr[4])
		if leftBaoNum >= num {
			return true
		}
	}
	return false
}

/**
吴名
2017/7/12 下午3:05
判断制飞是否成功:来任何一张牌都能胡(2,减去手上一张宝之后,剩下的牌形成整扑 1,减去手上一张宝之后,剩下的牌形成6对子)
*/
func zhiFei(arr []int, bao int) bool {
	mjArr := append([]int{}, arr...)

	//宝,万,筒,条,风
	sptArr := seperateRuiJinArr(mjArr, bao)

	//1:  13张牌,7对的制飞
	if len(mjArr) == 13 {
		danNum := 0
		for i, arr := range sptArr {
			l := len(arr)
			switch i {
			case 0:
				//宝不需要去除对子
				if l == 4 || (bao > 500 && l == 3) {
					log.Println("对对胡:所有的宝牌都在一家,飞")
					return false
				}
			default:
				danArr, _ := separate2Same(arr)
				danNum += len(danArr)
			}
		}
		//4,宝数量>=剩下的单张数量(去掉一张宝后,12张牌形成了6对)
		if len(sptArr[0])-1 >= danNum {
			log.Println("对对胡:制飞成功")
			return true
		}
	}

	//2, 减掉一张宝之后,剩下牌形成整扑

	//检测牌数量
	if len(mjArr)%3 != 1 {
		log.Println("牌数量不符合3n+1,不能飞")
		return false
	}

	needBaoNumToZhengPu := 0
	for i := 1; i <= 4; i++ {
		needBaoNumToZhengPu += getNeedBaoNumToZhengPu(sptArr[i])
	}
	if len(sptArr[0])-1 == needBaoNumToZhengPu {
		log.Println("3n+2胡:制飞成功")
		return true
	}
	return false
}

/**
吴名
2017/7/11 下午1:50
烂胡(在调用这个之前,已经排除了)
*/
func lanHu(arr []int, bao int) bool {

	mjArr := append([]int{}, arr...)

	//1,判断牌长度,必须为14
	if len(mjArr) != 14 {
		return false
	}
	//2,按类型分组
	//宝,万,筒,条,风
	sptArr := seperateRuiJinArr(mjArr, bao)
	//3,万筒条任意两张牌的差必须>=3
	//4,风牌任何两张牌不能相等
	for i, arr := range sptArr {
		l := len(arr)
		switch i {
		case 0:
			//宝不需要处理,只要万筒条风符合要求,不管几个宝,都可以配合
			if l == 4 || (bao > 500 && l == 3) {
				log.Println("烂胡:所有的宝牌都在一家,飞")
				return false
			}
		case 4:
			//风(任何两张不能相等)
			for j, _ := range arr {
				if j <= l-2 && arr[j] == arr[j+1] {
					log.Println("烂胡:风相等")
					return false
				}
			}
		default:
			//万筒条(差必须>=3),seperateRuiJinArr()返回的已经是排序后的
			for k, _ := range arr {
				if k <= l-2 && arr[k+1]-arr[k] <= 2 {
					log.Println("烂胡:万筒条未跳两张")
					return false
				}
			}
		}
	}
	log.Println("符合烂胡")
	return true
}

func duiDuiHu(arr []int, bao int) bool {

	mjArr := append([]int{}, arr...)

	//1,判断牌长度,必须为14
	if len(mjArr) != 14 {
		return false
	}
	//2,按类型分组
	//宝,万,筒,条,风
	sptArr := seperateRuiJinArr(mjArr, bao)
	//3,去掉所有对子,得到单张数量
	danNum := 0
	for i, arr := range sptArr {
		l := len(arr)
		switch i {
		case 0:
			//宝不需要去除对子
			if l == 4 || (bao > 500 && l == 3) {
				log.Println("对对胡:所有的宝牌都在一家,飞")
				return false
			}
		default:
			danArr, _ := separate2Same(arr)
			danNum += len(danArr)
		}
	}
	//4,宝数量>=剩下的单张数量
	if len(sptArr[0]) >= danNum {
		return true
	}
	return false
}

/**
吴名
2017/7/8 下午5:01
万,筒,条,风,成为整扑一将需要的最少癞子数量
*/
func getBaoNumToZhengPuJiang(arr []int) int {
	if len(arr) <= 0 {
		//如果数组为空,至少需要2个癞子组成一对将
		return 2
	}
	//寻找对子

	//吴名 2017/7/8 下午8:09 先去掉顺子的影响
	t := arr[0] / 100 //万筒条风类型
	a := []int{}
	switch t {
	case 4:
		a = separateFeng3Lian_ruiJin(arr)
	default:
		a, _ = separate3Lian(arr)
	}

	l := len(a)
	switch l {
	case 0:
		return 2
	case 1:
		return 1
	default:
		//可能是一张牌,两张牌,或3张牌
		for i, _ := range a {
			switch i {
			//只有2张牌,进入这里,为第一张
			case l - 2:
				if l == 2 && a[i] == a[i+1] {
					//找到对子了
					b := append(a[:i], a[i+2:]...)
					return getNeedBaoNumToZhengPu(b)
				}
				//最后1张牌
			case l - 1:
				//到最后一张牌,还没有找到对子

				//此时的3同参与不了顺子
				subArr := separate3Same(a)
				//2连不能拿去拼对子
				//吴名 2017/7/10 上午11:53 分离2连,再形成对子(101,104,105)
				canLianNum := 0
				switch t {
				case 4:
					subArr, canLianNum = separateFeng2Lian_ruiJin(subArr)
				default:
					subArr, canLianNum, _ = separate2Lian(subArr, -1, false)
				}

				switch len(subArr) {
				case 0:
					return canLianNum + 2
				case 1:
					return canLianNum + 1
				default:
					//101,105,105,105,不能拆3同

					return 1 + canLianNum + getNeedBaoNumToZhengPu(subArr[:len(subArr)-1]) //剩下的,最后1张牌拿去拼将去了(只能最后一张,101,101,101,104)
				}
			default:

				//101,104,104,104,107,不拆开3同(3同对形成顺子无影响)

				//举例总结:3同能参与形成顺子时,拆掉3同需要癞子数<=不拆,3同不能参与形成顺子时,3同利用不到,不能拆3同
				switch t {
				case 4:
					if a[i] == a[i+2] {
						//3同对子,但是对形成顺子有影响(401,401,401,404)
						if i >= 1 && (a[i] <= 404 || a[i-1] >= 405) {
							//401,404,404,404 或 405,406,406,406
							//3同能和前面形成顺子或有影响(包括4同)
							b := append(a[:i], a[i+2:]...)
							return getNeedBaoNumToZhengPu(b)
						}
						if l-i >= 4 && (a[i+3] <= 404 || a[i] >= 405) {
							//403,403,403,404 或 406,406,406,407

							//3同能和后面形成顺子或有影响(包括4同)
							b := append(a[:i], a[i+2:]...)
							return getNeedBaoNumToZhengPu(b)
						}
					} else {
						//纯对子(非3同中的对子)
						if a[i] == a[i+1] {
							if i == 0 || a[i-1] != a[i] {
								b := append(a[:i], a[i+2:]...)
								return getNeedBaoNumToZhengPu(b)
							}
						}
					}
				default:
					if a[i] == a[i+2] {
						//3同对子,但是对形成顺子有影响(101,103,103,103,105)
						if i >= 1 && a[i]-a[i-1] <= 2 {
							//3同能和前面形成顺子或有影响(包括4同)
							b := append(a[:i], a[i+2:]...)
							return getNeedBaoNumToZhengPu(b)
						}
						if l-i >= 4 && a[i+3]-a[i+2] <= 2 {
							//3同能和后面形成顺子或有影响(包括4同)
							b := append(a[:i], a[i+2:]...)
							return getNeedBaoNumToZhengPu(b)
						}
					} else {
						//纯对子(非3同中的对子)
						if a[i] == a[i+1] {
							if i == 0 || a[i-1] != a[i] {
								b := append(a[:i], a[i+2:]...)
								//fmt.Println("b:", b)
								return getNeedBaoNumToZhengPu(b)
							}
						}
					}
				}
			}
		}
	}
	return 0
}

/**
吴名
2017/7/6 下午7:05
万,筒,条,风,成为顺子或者三连需要的癞子数量
*/
func getNeedBaoNumToZhengPu(subArr []int) int {
	length := len(subArr) //万的张数
	switch length {
	case 0:
		return 0
	case 1:
		return 2
	case 2:
		t := subArr[0] / 100 //万筒条风类型
		switch t {
		case 4:
			//风
			if subArr[1] <= 404 {
				//两个都是东南西北(不管做顺子或刻),东南西北任何三个也可以互吃
				return 1
			}
			if subArr[0] >= 405 {
				//两个都是中发白,中发白任何三个可以互吃
				return 1
			}
			//一个东南西北,一个是中法白
			return 4
		default:
			//万,筒,条
			d := subArr[1] - subArr[0] //subArr是已经经过排序的
			if d <= 2 {
				//1万1万,1万2万,1万3万
				return 1
			} else {
				//1万4万
				return 4
			}
		}
	default:
		//3张以上万筒条或风
		//++++++++必须从小到大排序后,先去3同,再去3连,再去2同,再去2连,这些最容易形成整扑的去掉后,然后剩下牌两个一组分割算需要癞子数,这才能得到最少的癞子数量++++++++++
		//1,分离3同
		subArr = separate3Same(subArr)
		if len(subArr) <= 2 {
			//去除3同后剩余牌数<=2,直接结束
			return getNeedBaoNumToZhengPu(subArr)
		}
		t := subArr[0] / 100 //万筒条风类型
		switch t {
		case 4:
			//风
			//2,分离3连
			subArr = separateFeng3Lian_ruiJin(subArr)
			l := len(subArr)
			if l <= 2 {
				return getNeedBaoNumToZhengPu(subArr)
			} else {
				needCount := 0
				//3,分离2同
				subArr, duiZiNum := separate2Same(subArr)
				needCount += duiZiNum //有多少个对子就需要多少个癞子把它变整扑
				//3,分离2连
				subArr, canLianNum := separateFeng2Lian_ruiJin(subArr)
				needCount += canLianNum + getNeedBaoNumToZhengPu(subArr)
				return needCount
			}
		default:
			//万或筒或条
			//2,分离3连
			subArr, _ = separate3Lian(subArr)
			l := len(subArr)
			if l <= 2 {
				return getNeedBaoNumToZhengPu(subArr)
			} else {
				needCount := 0
				//3,分离2同和2连(相当于只需要1癞子就能成的牌都去掉)
				subArr, canLianOrSameNum, _ := separate2LianAnd2Same(subArr, -1, false)
				needCount += canLianOrSameNum + getNeedBaoNumToZhengPu(subArr) //有多少个对子或2连就需要多少个癞子把它变整扑
				return needCount
			}
		}
	}
}

/**
吴名
2017/7/8 上午11:54
分离2连(风,适用):返回去除后的数组,以及2连数量(里面3顺子,对子必须提前已去除)
*/
func separateFeng2Lian_ruiJin(arr []int) ([]int, int) {
	is := false
	lianNum := 0

	l := len(arr)
	for i, _ := range arr {
		if arr[i] != 0 && i <= l-2 {
			if arr[i+1] <= 404 || arr[i] >= 405 {
				//1,东南西北三张互吃 2,中发白互吃(对子前一步已经去除,所以不可能相等)
				arr[i] = 0
				arr[i+1] = 0
				lianNum++
				is = true
				break
			}
		}
	}

	if is {
		//如果祛除过顺子,那么需要清洗0之后继续祛除
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		a, num := separateFeng2Lian_ruiJin(r)
		return a, lianNum + num
	} else {
		return arr, 0
	}
}

/**
吴名
2017/7/7 下午5:45
分离顺子(针对风):东南西北任何三个也可以互吃,中发白任何三个可以互吃
*/
func separateFeng3Lian_ruiJin(arr []int) []int {
	is := false
	for i, _ := range arr {
		//前3张无对子的情况下(401,402,403,404)
		if i <= len(arr)-3 {
			if arr[i+2] <= 404 && arr[i] != arr[i+1] && arr[i+1] != arr[i+2] {
				//连续3张都是中南西北,且三张各不相等,可以互吃
				arr[i] = 0
				arr[i+1] = 0
				arr[i+2] = 0
				//log.Println("去除顺子:%v", arr)
				is = true
				break
			}
			if arr[i] == 405 && arr[i+1] == 406 && arr[i+2] == 407 {
				//连续3张是中发白
				arr[i] = 0
				arr[i+1] = 0
				arr[i+2] = 0
				//log.Println("去除顺子:%v", arr)
				is = true
				break
			}
		}
		//前3张有对子的情况下(401,401,402,403,404)
		if i <= len(arr)-4 {
			if arr[i+3] <= 404 && arr[i] != arr[i+1] && arr[i+1] != arr[i+3] {
				arr[i] = 0
				arr[i+1] = 0
				arr[i+3] = 0
				//log.Println("去除顺子:%v", arr)
				is = true
				break
			}
		}
	}
	if is {
		//如果祛除过顺子,那么需要清洗0之后继续祛除
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		return separateFeng3Lian_ruiJin(r)
	} else {
		return arr
	}
}

/**
吴名
2017/7/6 下午4:27
分割 宝,万,筒,条,风
*/
func seperateRuiJinArr(mjArr []int, bao int) [5][]int {
	result := [5][]int{}
	for _, mj := range mjArr {
		index := mj / 100
		//宝是花牌
		if bao > 500 {
			switch index {
			case 5:
				//宝
				result[0] = append(result[0], mj)
			default:
				result[index] = append(result[index], mj)
			}
		} else {
			switch index {
			case 5:
				//此时花牌处理成宝的本位牌
				i := bao / 100
				result[i] = append(result[i], bao)
			default:
				if mj == bao {
					//宝
					result[0] = append(result[0], mj)
				} else {
					result[index] = append(result[index], mj)
				}
			}
		}
	}
	//升序排列
	for _, arr := range result {
		sort.Sort(sort.IntSlice(arr))
	}
	return result
}
