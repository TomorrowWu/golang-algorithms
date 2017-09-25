package algorithm

import "testing"

/**
 * WuMing
 *2017/7/24 下午4:47
 *测试道州麻将
 */

func Benchmark_shiSanlan(b *testing.B) {

	for i := 0; i < b.N; i++ {
		//arr := []int{101, 104, 107, 201, 205, 208, 301, 304, 309,401,402,403,404,405}
		arr := []int{101, 104, 201, 208, 301, 304, 309, 401, 402, 403, 404, 406, 407}
		//hu := shiSanlan(arr, 405, 405)
		hu, s := IsHu(arr, 405, 405, true)
		switch hu {
		case NoHU:
			b.Logf("不胡 手牌中的顺子:%v", s)
		case LanHU:
			b.Logf("普通烂胡 手牌中的顺子:%v", s)
		case QiFengLanHU:
			b.Logf("七风到位十三烂 手牌中的顺子:%v", s)
		}
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

func Benchmark_qiQiaoDui(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//arr := []int{101, 104, 107, 201, 205, 208, 301, 304, 309,401,402,403,404,405}
		arr := []int{101, 101, 201, 201, 301, 301, 309, 309, 402, 402, 404, 404, 406}
		//arr := []int{101, 101, 101, 101, 301, 301, 309, 309, 402, 402, 404, 404, 406}
		//hu := qiQiaoDui(arr,406, 406, true)
		//hu := qiQiaoDui(arr, 405, 405, true)
		hu, s := IsHu(arr, 405, 405, true)
		switch hu {
		case NoHU:
			b.Logf("不胡 手牌中的顺子:%v", s)
		case YingLongDui:
			b.Logf("硬巧对+龙巧对 手牌中的顺子:%v", s)
		case YingDui:
			b.Logf("硬巧对 手牌中的顺子:%v", s) //ok
		case WangWangLongDui:
			b.Logf("王钓王+龙巧对 手牌中的顺子:%v", s) //ok
		case WangLongDui:
			b.Logf("王钓龙巧对 手牌中的顺子:%v", s)
		case LongDui:
			b.Logf("龙巧对")
		case WangWangQiDui:
			b.Logf("王钓王+七巧对 手牌中的顺子:%v", s) //ok
		case WangQiDui:
			b.Logf("王钓七巧对 手牌中的顺子:%v", s)
		case QiDui:
			b.Logf("七巧对 手牌中的顺子:%v", s)
		}
	}
}

func Benchmark_isHu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//arr := []int{101, 102, 201, 202, 203, 301, 302, 303, 402, 402, 404, 404, 404}
		//arr := []int{101, 102, 103, 104, 104, 201, 202, 203, 301, 302, 303, 404, 404}
		arr := []int{101, 102, 103, 104, 104, 201, 202, 203, 301, 302, 303, 404, 404}
		r, s := IsHu(arr, 405, 405, true)
		b.Logf("手牌中的顺子:%v", s)
		switch r {
		case NormalHU:
			b.Logf("NormalHU")
		case ShuangWangZhaoWang:
			b.Logf("ShuangWangZhaoWang")
		case ShuangWangZhao:
			b.Logf("ShuangWangZhao")
		case DanWangZhaoWang:
			b.Logf("DanWangZhaoWang")
		case DanWangZhao:
			b.Logf("DanWangZhao")
		case NoHU:
			b.Logf("NoHU")
		}
	}
}

func Benchmark_numberOfSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//shunZi := [][]int{{101, 102, 103}, {101, 102, 103}}
		//shunZi := [][]int{{101, 102, 103}, {102, 103,104}}
		shunZi := [][]int{{102, 103, 104}, {101, 102, 103}, {104, 105, 106}, {105, 106, 107}}
		n, w := NumberOfSet(shunZi, 103)
		b.Logf("套数:%v 王归位:%v", n, w)
	}
}

func Test_getNeedBaoNumToZhengPu_daoZhou(t *testing.T) {
	arr := []int{101,105,109}
	n,s:= getNeedBaoNumToZhengPu_daoZhou(arr,102)
	t.Logf("需要宝数量:%v,顺子:%v",n,s)
}
