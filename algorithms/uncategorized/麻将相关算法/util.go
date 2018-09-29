package algorithm

/**
 * WuMing
 *2017/7/6 下午2:41
 *麻将类通用的函数
 */

/**
吴名
2017/7/8 上午10:12
分离2同以及2连(万筒条,适用):返回去除后的数组,以及2同2连数量
*/
func separate2LianAnd2Same(arr []int, bao int, isGeneShunZi bool) ([]int, int, [][]int) {
	shunZi := [][]int{}

	//之前已经去除了对子,所以这里不用考虑对子
	is := false
	canlianOrSameNum := 0

	l := len(arr)
	for i, _ := range arr {
		if arr[i] != 0 && i <= l-2 {
			//只需要1张宝就能形成3连或3同
			d := arr[i+1] - arr[i]
			if isGeneShunZi {
				switch d {
				case 1:
					shunZi = append(shunZi, []int{arr[i], arr[i+1], bao})
				case 2:
					shunZi = append(shunZi, []int{arr[i], bao, arr[i+1]})
				}
			}

			if d <= 2 {
				arr[i] = 0
				arr[i+1] = 0
				is = true
				canlianOrSameNum++
			}
		}
	}

	if is {
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		return r, canlianOrSameNum, shunZi
	} else {
		return arr, 0, shunZi
	}
}

/**
吴名
2017/7/10 下午4:00
分离2连(万筒条,适用):返回去除后的数组,以及2连数量(去除了3连后,再来调用2连)
+++++++不需要生成顺子时,宝传 -1
*/
func separate2Lian(arr []int, bao int, isGeneShunZi bool) ([]int, int, [][]int) {
	shunZi := [][]int{}

	//之前已经去除了对子,所以这里不用考虑对子
	is := false
	canlianOrSameNum := 0

	l := len(arr)
	for i, _ := range arr {
		if arr[i] != 0 && i <= l-2 {
			sub := arr[i+1] - arr[i]

			if isGeneShunZi {
				switch sub {
				case 1:
					shunZi = append(shunZi, []int{arr[i], arr[i+1], bao})
				case 2:
					shunZi = append(shunZi, []int{arr[i], bao, arr[i+1]})
				}
			}

			if sub == 1 || sub == 2 {
				arr[i] = 0
				arr[i+1] = 0
				is = true
				canlianOrSameNum++
			}
		}
	}

	if is {
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		return r, canlianOrSameNum, shunZi
	} else {
		return arr, 0, shunZi
	}
}

/**
吴名
2017/7/7 上午11:34
分离顺子(针对万条筒)
*/
func separate3Lian(arr []int) ([]int, [][]int) {

	shunZi := [][]int{}

	is := false
	for i, _ := range arr {
		//前3张无对子的情况下(101,102,103)
		if i <= len(arr)-3 {
			if arr[i+1] == arr[i]+1 && arr[i+2] == arr[i]+2 {
				shunZi = append(shunZi, []int{arr[i], arr[i+1], arr[i+2]})

				arr[i] = 0
				arr[i+1] = 0
				arr[i+2] = 0
				//log.Println("去除顺子:%v", arr)
				is = true
				break
			}
		}
		//前3张有对子的情况下(101,102,102,103)
		if i <= len(arr)-4 {
			if arr[i+1] == arr[i]+1 && arr[i+3] == arr[i]+2 {
				shunZi = append(shunZi, []int{arr[i], arr[i+1], arr[i+3]})
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

		a, s := separate3Lian(r)

		shunZi = append(shunZi, s...)
		return a, shunZi
	} else {
		return arr, shunZi
	}
}

/**
吴名
2017/7/7 上午10:56
分离3同(万筒条,风,都适用)
*/
func separate3Same(arr []int) []int {
	is := false
	for i, _ := range arr {
		if arr[i] != 0 && i <= len(arr)-3 {
			//因为是已经从小到大排序后的,所以这俩相等,必定是3同
			if arr[i] == arr[i+2] {
				arr[i] = 0
				arr[i+1] = 0
				arr[i+2] = 0
				is = true
			}
		}
	}

	if is {
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		return r
	} else {
		return arr
	}
}

/**
吴名
2017/7/7 下午11:53
分离对子(万筒条,风,都适用):返回去除后的数组,以及对子数量,当然在当前场景下,只用来去除风的对子
*/
func separate2Same(arr []int) ([]int, int) {
	is := false
	duiZiNum := 0

	for i, _ := range arr {
		if arr[i] != 0 && i <= len(arr)-2 {
			if arr[i] == arr[i+1] {
				arr[i] = 0
				arr[i+1] = 0
				is = true
				duiZiNum++
			}
		}
	}

	if is {
		r := []int{}
		for _, v := range arr {
			if v != 0 {
				r = append(r, v)
			}
		}
		return r, duiZiNum
	} else {
		return arr, 0
	}
}
