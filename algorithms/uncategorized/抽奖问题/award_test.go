package 抽奖问题

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
 * WuMing
 *2017/9/5 下午3:42
 *
 */

func Test_GetAwardUserName(t *testing.T) {
	//var users map[string]int64 = map[string]int64{
	//	"a": 10,
	//	"b": 6,
	//	"c": 3,
	//	"d": 12,
	//	"f": 1,
	//}

	//awardCount := make(map[string]int)
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i <= 10000; i++ {
	//	awardName := GetAwardUserName(users)
	//	if count, ok := awardCount[awardName]; ok {
	//		awardCount[awardName] = count + 1
	//	} else {
	//		awardCount[awardName] = 0
	//	}
	//}
	//for n, c := range awardCount {
	//	fmt.Printf("%v:%v\n", n, c)
	//}

	//for i := 0; i <= 4; i++ {
	//	fmt.Println(GetAwardUserName(users))
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func Test_getAwardUser_weight(t *testing.T) {
	var users = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"f": 1,
	}

	//rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i <= 100; i++ {
		awardName := getAwardUser_weight(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 0
		}
	}
	for n, c := range awardCount {
		fmt.Printf("%v:%v \n", n, c)
	}
}

func Test_append(t *testing.T) {
	test_len := 200
	start := time.Now().Unix()
	s := make([]int, 0, test_len)
	for i := 0; i < test_len; i++ {
		s = append(s, i)
	}
	fmt.Println(time.Now().Unix() - start)

	start_1 := time.Now().Unix()
	s1 := make([]int, test_len, test_len)
	for i := 0; i < test_len; i++ {
		s1[i] = i
	}
	fmt.Println(time.Now().Unix() - start_1)
}
