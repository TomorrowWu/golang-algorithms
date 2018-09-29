package algorithm

import (
	"sort"
	"log"
)

/**
吴名
2017/7/24 下午4:06
胡牌类型
*/
const (
	NoHU = 0 //不胡

	LanHU       = 1 //普通十三烂(点炮人付3番、自摸每人2番)
	QiFengLanHU = 2 //七风到位十三烂(点炮人付6番,自摸每人4番)

	YingLongDui     = 3  //硬巧对+龙巧对(点炮32番，自摸每人16番)
	YingDui         = 4  //硬巧对(点炮12番，自摸每人8番)
	WangWangLongDui = 5  ////王钓王+龙巧对(每人16番)
	WangLongDui     = 6  ////王钓龙巧对(每人8番)
	LongDui         = 7  //龙巧对(点炮人付6番，自摸每人4番)
	WangWangQiDui   = 8  //王钓王+七巧对(每人8番)
	WangQiDui       = 9  //王钓七巧对(每人4番)
	QiDui           = 10 //七巧对(点炮人付3番，自摸每人2番)

	NormalHU = 11 //3n+2 胡牌

	ShuangWangZhaoWang = 12 //双王爪王(8番)
	ShuangWangZhao     = 13 //双王爪(4番)
	DanWangZhaoWang    = 14 //单王爪王(4番)
	DanWangZhao        = 15 //单王爪(2番)
)

/**
 * WuMing
 *2017/7/20 下午3:32
 *道州麻将胡牌
 */
var (
	daoZhoumahjongArr = []int{
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
	}
)

/**
吴名
2017/7/26 下午3:49
胡牌
arr:手牌(不包含摸牌以及上家出的牌)
pai:摸牌或出牌
*/
func IsHu(arr []int, pai, bao int, isZiMo bool) (int, [][]int) {
	shunZi := [][]int{}
	//检测牌数量(不包含摸牌或出牌)
	l := len(arr)

	if (pai > 0 && l == 13) || (pai <= 0 && l == 14 && isZiMo) {
		sSLResult := shiSanlan(arr, bao, pai)
		if sSLResult != NoHU {
			//胡十三烂
			return sSLResult, shunZi
		}

		qQDResult := qiQiaoDui(arr, pai, bao, isZiMo)
		if qQDResult != NoHU {
			//胡七巧对
			return qQDResult, shunZi
		}
	}

	if l%3 != 1 {
		log.Print("牌数量不符合3n+2")
		return NoHU, shunZi
	}

	sptArr := seperateArr_daoZhou(arr, bao, pai, isZiMo, true)

	a := [5][]int{}
	for i, arr := range sptArr {
		a[i] = append(a[i], arr...)
	}
	//canHu里面会修改切片值,所以得复制
	isHu, s := canHu_daoZhou(a, bao, 0)
	shunZi = append(shunZi, s...)
	if isHu {
		// 吴名 2017/7/27 下午3:42 判断3n+2具体胡牌类型

		//首先得根据是否自摸
		if isZiMo {
			//宝,万,筒,条,风
			sptArr := seperateArr_daoZhou(arr, bao, -1, isZiMo, true)
			b := [5][]int{}
			for i, arr := range sptArr {
				b[i] = append(b[i], arr...)
			}

			baoNum := len(b[0]) //手牌中宝的数量

			if baoNum >= 2 {
				//减去手上的2张王,看是否形成3n+2
				if b, _ := canHu_daoZhou(b, bao, 2); b {

					switch pai {
					case bao:
						//双王爪王
						return ShuangWangZhaoWang, shunZi
					default:
						//双王爪
						return ShuangWangZhao, shunZi
					}
				}
			}

			if baoNum >= 1 {
				//减去手上的1张王,看是否形成3n
				needBaoNumToZhengPu := 0
				for i := 1; i <= 4; i++ {
					n, _ := getNeedBaoNumToZhengPu_daoZhou(sptArr[i], bao)
					needBaoNumToZhengPu += n
				}
				if baoNum-1 >= needBaoNumToZhengPu {
					switch pai {
					case bao:
						//单王爪王
						return DanWangZhaoWang, shunZi
					default:
						//单王爪
						return DanWangZhao, shunZi
					}
				}
			}
		}
		return NormalHU, shunZi //普通的自摸或点炮胡
	} else {
		return NoHU, shunZi
	}
}

/**
吴名
2017/7/11 下午1:50
十三烂(王和红中都是本位牌)
arr:包含摸牌,出牌
*/
func shiSanlan(arr []int, bao, pai int) int {
	//1,判断牌长度,必须为14

	l := len(arr)
	if (pai > 0 && l == 13) || (pai <= 0 && l == 14) {
		//2,按类型分组
		//宝(做本位牌,所以不用管是否自摸),万,筒,条,风
		sptArr := seperateArr_daoZhou(arr, bao, pai, false, false)
		//3,万筒条任意两张牌的差必须>=3
		//4,风牌任何两张牌不能相等
		for i, arr := range sptArr {
			l := len(arr)
			switch i {
			case 0:
			case 4:
				//风(任何两张不能相等)
				for j := range arr {
					if j <= l-2 && arr[j] == arr[j+1] {
						//log.Println("烂胡:风相等")
						return NoHU
					}
				}
			default:
				//万筒条(差必须>=3),seperateRuiJinArr()返回的已经是排序后的
				for k := range arr {
					if k <= l-2 && arr[k+1]-arr[k] <= 2 {
						//log.Println("烂胡:万筒条未跳两张")
						return NoHU
					}
				}
			}
		}

		//5,判断是否 7风到位
		if len(sptArr[4]) == 7 {
			//log.Println("七风到位十三烂")
			return QiFengLanHU
		} else {
			//log.Println("普通十三烂")
			return LanHU
		}

	}
	return NoHU
}

/**
吴名
2017/7/21 下午7:27
七巧对
arr手牌(不包含摸牌以及上家出的牌),isZiMo只是为了判断王钓
pai:摸牌或出牌
*/
func qiQiaoDui(arr []int, pai, bao int, isZiMo bool) int {

	//1,判断牌长度,必须为14
	l := len(arr)
	if (pai > 0 && l == 13) || (pai <= 0 && l == 14) {
		mjArr := append(append([]int{}, arr...), pai) //所有牌
		//排序
		sort.Sort(sort.IntSlice(mjArr))
		//2,按类型分组
		//宝,万,筒,条,风
		sptArr := seperateArr_daoZhou(arr, bao, pai, isZiMo, true)
		//3,去掉所有对子,得到单张数量
		danNum := 0
		//+++++++++++++++++++这里,不管红中是做红中,还是王的本位牌,最后剩下的红中都属于单张,都需要癞子搭配成对子+++++++++++++++++++++
		for i, arr := range sptArr {
			switch i {
			case 0:
			default:
				danArr, _ := separate2Same(arr)
				danNum += len(danArr)
			}
		}
		//4,宝数量>=剩下的单张数量
		wangNum := len(sptArr[0])
		if wangNum >= danNum {
			//判断是否王钓(6对+1王)
			isWangdiao := false
			if isZiMo {
				isWangdiao = isWangDiao(arr, bao)
			}
			//判断是否龙巧对(一暗杠,王不能参与)
			isLongQiaoDui := false
			for i, v := range mjArr {
				if i <= 10 && v != bao && mjArr[i] == mjArr[i+3] {
					isLongQiaoDui = true
					break
				}
			}

			switch wangNum {
			case 0:
				if isLongQiaoDui {
					//硬巧对+龙巧对(点炮算12番，自摸每人8番)
					return YingLongDui
				} else {
					//硬巧对(点炮32番，自摸每人16番)
					return YingDui
				}

			default:
				if isLongQiaoDui {
					if isWangdiao {
						if pai == bao {
							//王钓王+龙巧对(每人16番)
							return WangWangLongDui
						} else {
							//王钓龙巧对(每人8番)
							return WangLongDui
						}
					} else {
						//龙巧对(点炮人付6番，自摸每人4番)
						return LongDui
					}
				} else {
					if isWangdiao {
						if pai == bao {
							//王钓王+七巧对(每人8番)
							return WangWangQiDui
						} else {
							//王钓七巧对(每人4番)
							return WangQiDui
						}
					} else {
						//七巧对(点炮人付3番，自摸每人2番)
						return QiDui
					}
				}
			}
		}
	}

	return NoHU
}

/**
吴名
2017/7/29 下午4:28
是否清一色(不能有风,除了王之外,其他牌全是万或筒或条)
*/
func isQingYiSe(sptArr [5][]int, paiNum int) bool {

	if len(sptArr[4]) > 0 {
		return false
	}

	baoNum := len(sptArr[0])

	//宝的数量+万或筒或条的数量=paiNum

	wanNum := len(sptArr[1])
	if wanNum > 0 {
		if baoNum+wanNum == paiNum {
			return true
		} else {
			return false
		}
	}

	tongNum := len(sptArr[2])
	if tongNum > 0 {
		if baoNum+tongNum == paiNum {
			return true
		} else {
			return false
		}
	}

	tiaoNum := len(sptArr[3])
	if tiaoNum > 0 {
		if baoNum+tiaoNum == paiNum {
			return true
		} else {
			return false
		}
	}

	return false
}

/**
吴名
2017/7/29 下午3:29
是否符合3n+2
reduceBaoNum:要减去的宝的数量(为了计算双王爪,单王爪)
*/
func canHu_daoZhou(sptArr [5][]int, bao, reduceBaoNum int) (bool, [][]int) {
	baoNum := len(sptArr[0]) - reduceBaoNum //可用宝的数量

	needBaoArr := [5]int{}
	huShunZi := [5][][]int{}
	for i := 1; i <= 4; i++ {
		a := append([]int{}, sptArr[i]...)
		needNum, s := getNeedBaoNumToZhengPu_daoZhou(a, bao)
		needBaoArr[i] = needNum

		huShunZi[i] = s //存下手上的顺子
	}

	// 吴名 2017/7/7 下午8:30 将在"万"中
	needBaoNum := needBaoArr[2] + needBaoArr[3] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "万"成整扑一将的癞子数量
		num, s := getBaoNumToZhengPuJiang_daoZhou(sptArr[1], bao)
		if leftBaoNum >= num {
			s = append(append(s, huShunZi[2]...), huShunZi[3]...)
			return true, s
		}
	}

	// 吴名 2017/7/7 下午8:30 将在"筒"中
	needBaoNum = needBaoArr[1] + needBaoArr[3] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "筒"成整扑一将的癞子数量
		num, s := getBaoNumToZhengPuJiang_daoZhou(sptArr[2], bao)
		if leftBaoNum >= num {
			s = append(append(s, huShunZi[1]...), huShunZi[3]...)
			return true, s
		}
	}

	// 吴名 2017/7/7 下午8:31 将在"条"中
	needBaoNum = needBaoArr[1] + needBaoArr[2] + needBaoArr[4]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "条"成整扑一将的癞子数量
		num, s := getBaoNumToZhengPuJiang_daoZhou(sptArr[3], bao)
		if leftBaoNum >= num {
			s = append(append(s, huShunZi[1]...), huShunZi[2]...)
			return true, s
		}
	}

	// 吴名 2017/7/7 下午8:31 将在"风"中
	needBaoNum = needBaoArr[1] + needBaoArr[2] + needBaoArr[3]
	if needBaoNum <= baoNum {
		leftBaoNum := baoNum - needBaoNum //剩下可用于去拼 "风"成整扑一将的癞子数量
		num, s := getBaoNumToZhengPuJiang_daoZhou(sptArr[4], bao)
		if leftBaoNum >= num {
			s = append(append(append(s, huShunZi[1]...), huShunZi[2]...), huShunZi[3]...) //风没有顺子
			return true, s
		}
	}
	return false, [][]int{}
}

/**
吴名
2017/7/8 下午5:01
万,筒,条,风,成为整扑一将需要的最少癞子数量
*/
func getBaoNumToZhengPuJiang_daoZhou(arr []int, bao int) (int, [][]int) {
	shunZi := [][]int{}

	l := len(arr)
	switch l {
	case 0:
		//必须得有一对将
		return 2, shunZi
	case 1:
		return 1, shunZi
	default:
		//吴名 2017/7/8 下午8:09 +++++++++++++1,万筒条去掉顺子的影响 2,风去掉3同的影响+++++++++++
		t := arr[0] / 100 //万筒条风类型
		a := []int{}
		switch t {
		case 4:
			//道州麻将中,风不能做顺子,只能3同
			a = separate3Same(arr)
		default:
			s := [][]int{}
			a, s = separate3Lian(arr)
			shunZi = append(shunZi, s...)
		}
		l = len(a)
		if l <= 1 {
			n, s := getBaoNumToZhengPuJiang_daoZhou(a, bao)
			shunZi = append(shunZi, s...)
			return n, shunZi
		}
		//可能是一张牌,两张牌,或3张牌
		for i, _ := range a {
			switch i {
			//只有2张牌,进入这里,为第一张
			case l - 2:
				if l == 2 && a[i] == a[i+1] {
					//找到对子了
					b := append(a[:i], a[i+2:]...)
					n, s := getNeedBaoNumToZhengPu_daoZhou(b, bao)
					shunZi = append(shunZi, s...)
					return n, shunZi
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
					//此时还没有对子,剩下全是风的单张
				default:
					s := [][]int{}
					subArr, canLianNum, s = separate2Lian(subArr, bao, true)
					shunZi = append(shunZi, s...)
				}

				switch len(subArr) {
				case 0:
					return canLianNum + 2, shunZi
				case 1:
					return canLianNum + 1, shunZi
				default:
					//101,105,105,105,不能拆3同

					n, s := getNeedBaoNumToZhengPu_daoZhou(subArr[:len(subArr)-1], bao)
					shunZi = append(shunZi, s...)
					return 1 + canLianNum + n, shunZi //剩下的,最后1张牌拿去拼将去了(只能最后一张,101,101,101,104)
				}
			default:
				//101,104,104,104,107,不拆开3同(3同对形成顺子无影响)

				//举例总结:3同能参与形成顺子时,拆掉3同需要癞子数<=不拆,3同不能参与形成顺子时,3同利用不到,不能拆3同
				b := []int{}
				isFindDui := false
				switch t {
				case 4:
					//纯对子(此时风牌无3同)
					if a[i] == a[i+1] {
						b = append(a[:i], a[i+2:]...)
						isFindDui = true
					}
				default:
					if a[i] == a[i+2] {
						//3同对子,但是对形成顺子有影响(101,103,103,103,105)
						if (i >= 1 && a[i]-a[i-1] <= 2) || (l-i >= 4 && a[i+3]-a[i+2] <= 2) {
							//3同能和前面形成顺子或有影响(包括4同)
							//3同能和后面形成顺子或有影响(包括4同)
							b = append(a[:i], a[i+2:]...)
							isFindDui = true
						}
					} else {
						//纯对子(非3同中的对子)
						if a[i] == a[i+1] {
							if i == 0 || a[i-1] != a[i] {
								b = append(a[:i], a[i+2:]...)
								isFindDui = true
							}
						}
					}
				}
				if isFindDui {
					n, s := getNeedBaoNumToZhengPu_daoZhou(b, bao)
					shunZi = append(shunZi, s...)
					return n, shunZi
				}
			}
		}
	}
	return 0, shunZi
}

/**
 吴名
 2017/7/6 下午7:05
 万,筒,条,成为顺子或者三同需要的癞子数量
风只能做3同
*/
func getNeedBaoNumToZhengPu_daoZhou(subArr []int, bao int) (int, [][]int) {
	shunZi := [][]int{}

	length := len(subArr) //万的张数
	switch length {
	case 0:
		return 0, shunZi
	case 1:
		return 2, shunZi
	case 2:
		t := subArr[0] / 100 //万筒条风类型
		switch t {
		case 4:
			//风
			if subArr[0] == subArr[1] {
				return 1, shunZi
			} else {
				return 4, shunZi
			}
		default:
			//万,筒,条
			d := subArr[1] - subArr[0] //subArr是已经经过排序的
			switch d {
			case 0:
				//1万1万
				return 1, shunZi
			case 1:
				shunZi = append(shunZi, []int{subArr[0], subArr[1], bao})
				//1万2万
				return 1, shunZi
			case 2:
				shunZi = append(shunZi, []int{subArr[0], bao, subArr[1]})
				//1万3万
				return 1, shunZi
			default:
				//1万4万
				return 4, shunZi
			}
		}
	default:
		//3张以上万筒条或风
		//++++++++必须从小到大排序后,先去3同,再去3连,再去2同,再去2连,这些最容易形成整扑的去掉后,然后剩下牌两个一组分割算需要癞子数,这才能得到最少的癞子数量++++++++++
		//1,分离3同
		subArr = separate3Same(subArr)
		if len(subArr) <= 2 {
			//去除3同后剩余牌数<=2,直接结束
			n, s := getNeedBaoNumToZhengPu_daoZhou(subArr, bao)
			shunZi = append(shunZi, s...)
			return n, shunZi
		}
		t := subArr[0] / 100 //万筒条风类型
		switch t {
		case 4:
			//风(只能做3同)
			needCount := 0
			//3,分离2同
			subArr, duiZiNum := separate2Same(subArr)
			needCount += duiZiNum        //有多少个对子就需要多少个癞子把它变整扑
			needCount += len(subArr) * 2 //有多少单张,就需要*2个癞子形成整扑
			return needCount, shunZi
		default:
			//万或筒或条
			//2,分离3连

			subArr, sh := separate3Lian(subArr)
			shunZi = append(shunZi, sh...)

			l := len(subArr)
			if l <= 2 {
				n, s := getNeedBaoNumToZhengPu_daoZhou(subArr, bao)
				shunZi = append(shunZi, s...)
				return n, shunZi
			} else {
				needCount := 0
				//3,分离2同和2连(相当于只需要1癞子就能成的牌都去掉)
				subArr, canLianOrSameNum, s := separate2LianAnd2Same(subArr, bao, true)
				shunZi = append(shunZi, s...)
				needCount += canLianOrSameNum //有多少个对子或2连就需要多少个癞子把它变整扑
				if len(subArr) <= 2 {
					n, s := getNeedBaoNumToZhengPu_daoZhou(subArr, bao)
					needCount += n
					shunZi = append(shunZi, s...)
					return needCount, shunZi
				} else {
					for i := 0; i <= l-1; i += 2 {
						switch i {
						case l - 2, l - 1:
							//最后2张(len为偶数)
							//最后1张(len为基数)
							n, s := getNeedBaoNumToZhengPu_daoZhou(subArr[i:], bao)
							needCount += n
							shunZi = append(shunZi, s...)
						default:
							n, s := getNeedBaoNumToZhengPu_daoZhou(subArr[i:i+2], bao)
							needCount += n
							shunZi = append(shunZi, s...)
						}
					}
					return needCount, shunZi
				}
			}
		}
	}
	return 0, shunZi
}

/**
吴名
2017/7/22 下午5:23
13张牌是否形成王钓
*/
func isWangDiao(arr []int, bao int) bool {
	if len(arr) == 13 {
		//2,按类型分组
		//宝,万,筒,条,风
		sptArr := seperateArr_daoZhou(arr, bao, -1, true, true)
		danNum := 0
		for i, a := range sptArr {
			switch i {
			case 0:
			default:
				danArr, _ := separate2Same(a)
				danNum += len(danArr)
			}
		}
		//4,宝数量>=剩下的单张数量(去掉一张宝后,12张牌形成了6对)
		if len(sptArr[0])-1 >= danNum {
			//log.Println("七巧对:王钓成功")
			return true
		}
	}
	return false
}

/**
吴名
2017/7/28 下午7:34
点炮时分割 宝,万,筒,条,风
mjArr:不包含出的pai
*/
func seperateArr_daoZhou(mjArr []int, bao, pai int, isZiMo, isHandleBao bool) [5][]int {

	arr := append([]int{}, mjArr...)

	result := [5][]int{}

	if pai <= 0 {
		//说明mjArr中包含了摸牌或出牌

	} else {
		switch isZiMo {
		case true:
			arr = append(arr, pai)
		case false:
			//当成本位(宝作为宝的本位,红中还是做红中)
			index := pai / 100
			result[index] = append(result[index], pai)
		}
	}

	for _, mj := range arr {
		index := mj / 100
		switch isHandleBao {
		case true:
			//分离出宝,红中当宝的本位牌处理
			switch mj {
			case bao:
				result[0] = append(result[0], bao)
			case 405:
				//此时红中不为宝,红中做宝的本位牌
				i := bao / 100
				result[i] = append(result[i], bao)
			default:
				result[index] = append(result[index], mj)
			}
		case false:
			result[index] = append(result[index], mj)
		}
	}

	//升序排列
	for _, arr := range result {
		sort.Sort(sort.IntSlice(arr))
	}
	return result
}

/**
 吴名
 2017/8/4 上午10:52
套数(有一套,有两套):王做本位牌的顺子才算有一套中的相等
shunZi:手牌中的顺子
*/
func NumberOfSet(shunZi [][]int, bao int) (number int, wangGuiWei bool) {

	sZ := append([][]int{}, shunZi...)

	l := len(sZ)
	if l < 2 {
		return
	}

	//TODO 吴名 2017/8/11 下午5:08 shunZi排序

	isHaveSet := false
exit:
	for i, s := range sZ {
		for j := i + 1; j <= l-1; j++ {
			jArr := sZ[j]
			//++++++++++++++++连续两个数字相等,哪怕是有王,这时候的王是算本位牌,也算是两个数字相等++++++++++++++++
			//102,103,104 101,102,103
			//101,102,103 102,103,104
			if (s[0] == jArr[0] && s[1] == jArr[1]) || (s[0] == jArr[1] && s[1] == jArr[2]) || (s[1] == jArr[0] && s[2] == jArr[1]) {
				//俩数组形成有一套
				number++
				if !wangGuiWei {
					if (s[0] == bao && s[1] == bao+1) || (s[1] == bao && s[2] == bao+2) || (s[2] == bao && s[1] == bao-1) || (jArr[0] == bao && jArr[1] == bao+1) || (jArr[1] == bao && jArr[2] == bao+2) || (jArr[2] == bao && jArr[1] == bao-1) {
						wangGuiWei = true
					}
				}
				sZ = append(sZ[i+1:j], sZ[j+1:]...)
				isHaveSet = true
				break exit
			}
		}
	}
	if isHaveSet && len(sZ) >= 2 {
		n, w := NumberOfSet(sZ, bao)
		number += n
		if !wangGuiWei && w {
			wangGuiWei = true
		}
	}
	return
}
