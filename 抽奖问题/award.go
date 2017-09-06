package 抽奖问题

import "math/rand"

/**
 * WuMing 
 *2017/9/5 下午3:41
 *
 */

//2,用户随机抽奖
// map中，key代表名称，value代表成交单数
//思路:随机问题,一般都是使用rand包下的随机函数得到0到n-1中的一个随机的index,即用户数组中的index,从map中选取随机用户,把map转为数组解决问题,

func GetAwardUserName(users map[string]int64) (name string) {
	size := len(users)
	awardIndex := rand.Intn(size)

	i := 0
	for userName, _ := range users {
		if i == awardIndex {
			name = userName
			return
		}
		i++
	}
	return
}

//3,权重抽奖(用户count越大,获奖概率越大)
//思路:index不再代表某个用户,而且n到n+count这个数字范围代表某个用户,类似数轴的思想,rand随机函数获取一个随机数字,此数字在那个数轴范围内,即为某用户
func getAwardUser_weight(users map[string]int64) (name string) {
	userArr := make([]string, len(users),len(users))
	var sumCount int64 = 0
	index := 0
	for n, c := range users {
		//整理所有用户的count数据为数轴
		userArr[index] = n
		index++
		sumCount += c
	}

	awardIndex := rand.Int63n(sumCount)

	var offset int64
	for _, n := range userArr {
		//判断获奖index落在那个用户区间内
		offset += users[n]
		if offset > awardIndex {
			name = n
			return
		}
	}
	return
}
