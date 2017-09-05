package algorithm

import (
	"testing"
)

/**
 * WuMing
 *2017/7/7 上午11:07
 *
 */
func Test_Separate3Same(t *testing.T) {
	arr := []int{1, 1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	t.Logf("原slice:%v", arr)

	t.Logf("去除3同结果:%v", separate3Same(arr))
}

func Test_separate3Lian(t *testing.T) {
	//arr := []int{2, 3, 4, 5, 5, 6}
	//arr := []int{4, 5, 5,6,7}
	arr := []int{2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9, 9}
	t.Logf("原slice:%v", arr)

	a, _ := separate3Lian(arr)
	t.Logf("去除3连结果:%v", a)
}

func Test_getNeedBaoInSub(t *testing.T) {
	//arr := []int{101,101,102,102, 103, 104, 105, 106, 107, 108, 109}
	//arr := []int{101,101,102}
	//arr := []int{209}
	arr := []int{101, 103, 103, 106, 107}
	//arr := []int{401,402,402,405,405,407}
	t.Logf("形成3同或3连最少需要的癞子数量:%v", getNeedBaoNumToZhengPu(arr))
}

func Test_separateFeng3Lian(t *testing.T) {
	arr := []int{401, 404, 404, 405, 406, 407}
	t.Logf("去除风牌3连结果:%v", separateFeng3Lian_ruiJin(arr))
}

func Test_separate2LianAnd2Same(t *testing.T) {
	arr := []int{301, 301, 303, 304, 304, 306, 306, 307, 307}
	arr, num, _ := separate2LianAnd2Same(arr, -1, false)
	t.Logf("万条筒去除2连结果:%v 2连和2同数量:%v", arr, num)
}

func Test_separateFeng2Lian(t *testing.T) {
	arr := []int{404, 405}
	arr, num := separateFeng2Lian_ruiJin(arr)
	t.Logf("风牌去除2连结果:%v 2连数量:%v", arr, num)
}

func Test_getBaoNumToZhengPuJiang(t *testing.T) {
	//arr := []int{401,401,401,404}
	//arr := []int{101,105,105,105}
	//arr := []int{101,101,101,104,105}
	//arr := []int{101,101,101,104,105}
	//arr := []int{101,101,101}
	arr := []int{101, 101, 104}
	//arr := []int{401,404,404,404,404}
	//arr := []int{404,404,404,404,405}
	//arr := []int{101, 104, 104, 104, 104}
	num := getBaoNumToZhengPuJiang(arr)
	t.Logf("形成整扑一将需要的癞子数量:%v", num)
}

func Test_lanHu(t *testing.T) {
	//arr := []int{101, 104, 104, 104, 104}
	arr := []int{101, 104, 104, 107, 201, 204, 209, 302, 305, 309, 403, 406, 409, 501}
	t.Logf("烂胡:%v", lanHu(arr, 104))
}

func Test_duiDuiHu(t *testing.T) {
	arr := []int{101, 101, 104, 104, 201, 201, 209, 209, 305, 306, 407, 407, 407, 407}
	t.Logf("对对胡:%v", duiDuiHu(arr, 407))
}

func Test_zhiFei(t *testing.T) {
	arr := []int{101, 101, 104, 104, 201, 201, 209, 209, 305, 306, 407, 407, 407}
	t.Logf("制飞:%v", zhiFei(arr, 407))
}

func Test_isHU(t *testing.T) {
	arr := []int{101, 101, 104, 104, 201, 201, 209, 209, 305, 306, 307, 407, 407, 407}
	t.Logf("胡:%v", isHU(arr, 407))
}

func Benchmark_isHU(b *testing.B) {

	for i := 0; i < b.N; i++ {
		arr := []int{101, 101, 104, 104, 201, 201, 209, 209, 305, 306, 307, 407, 407, 407}
		isHU(arr, 407)
		//b.Logf("胡:%v", isHU(arr, 407))
	}

	//b.RunParallel(func(pb *testing.PB) {
	//
	//	for pb.Next() {
	//		arr := []int{101, 101, 104, 104, 201, 201, 209, 209, 305, 306, 307, 407, 407, 407}
	//		isHU(arr, 407)
	//		//b.Logf("胡:%v", isHU(arr, 407))
	//	}
	//})
}
