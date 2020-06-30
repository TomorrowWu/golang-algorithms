package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"time"
)

//var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

type tmp struct {
	A string `json:"a"`
}

func main() {
	//now := time.Now()
	////本月1号
	//curMon := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(curMon.Unix(), now.Unix())
	//
	//var best map[string]interface{}
	//fmt.Println(best)
	//fmt.Println(best == nil)
	//
	//a1 := &tmp{A: "a"}
	//
	//a2 := copyTmp(a1)
	//fmt.Println(a1, a2)
	//
	//a2.A = "b"
	//fmt.Println(a1, a2)
	//
	//arr1 := []int64{1, 2, 3, 4, 5, 6}
	//arr2 := CopyInt64Arr(arr1)
	//fmt.Println(arr1, arr2)
	//arr2 = arr2[1:4]
	//fmt.Println(arr1, arr2)
	//
	//arr2[1] = 1
	//arr2[2] = 1
	//fmt.Println(arr1, arr2)
	//
	//songIDs := []int64{1, 2, 3, 4, 5}
	//if 5 <= len(songIDs) {
	//	// 未越界
	//	songIDs = songIDs[0:5]
	//} else {
	//	// 越界
	//	songIDs = songIDs[0:]
	//}
	//
	//fmt.Println("www.baidu.com?name=" + url.QueryEscape("wu ming"))
	//
	//splitN := strings.SplitN("a,", ",", -1)
	//fmt.Println(splitN, len(splitN), "1"+splitN[1]+"1")
	//
	//sli1 := []int64{1, 2, 3}
	//sli2 := []int64{4, 5, 6}
	//r := append(sli1, sli2...)
	//fmt.Println(r, sli1, sli2)

	//src := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(SplitInt64s(src, 2))
	//fmt.Println(SplitInt64s(src, 1))
	//fmt.Println(SplitInt64s(src, 3))
	//fmt.Println(SplitInt64s(src, 4))
	//fmt.Println(SplitInt64s(src, 5))
	//fmt.Println(SplitInt64s(src, 6))
	//fmt.Println(SplitInt64s(src, 10))
	//fmt.Println(SplitInt64s(src, 11))
	//fmt.Println(SplitInt64s(src, 0))
	//
	//fmt.Println("-----------\n", src)

	fmt.Println("************\n")

	var m map[int64]string
	fmt.Println("-----------", m[1], "----------")

	fmt.Println("************\n")

	//tmp := []struct {
	//	id   int64
	//	join int64
	//}{
	//	{
	//		1, 2,
	//	},
	//	{2, 8},
	//	{3, 243},
	//	{4, 238},
	//	{5, 254},
	//}
	//sort.Slice(tmp, func(i, j int) bool {
	//	// 由大到小
	//	v1 := int64(tmp[i].join / 1)
	//	v2 := int64(tmp[j].join / 1)
	//	return v1 > v2
	//})
	//fmt.Println("----------", tmp, "----------")

	fmt.Println(GetDeviceIdTail("0242C83737B3C7CDB836B9763C9EA984"))
	fmt.Println(time.Unix(time.Now().Unix(), 0), time.Unix(time.Now().Unix(), 0).Format("01/02/2006"))

	s := `{"is_cover":"1"}`
	i := struct {
		IsCover bool `json:"is_cover,string"`
	}{}
	if err := json.Unmarshal([]byte(s), &i); err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

}

func copyTmp(src *tmp) tmp {
	return *src
}

func CopyInt64Arr(src []int64) []int64 {
	n := make([]int64, len(src), len(src))
	copy(n, src)
	return n
	//return src
}

func SplitInt64s(src []int64, size int) (ret [][]int64) {
	if size <= 0 {
		size = len(src)
	}
	for len(src) > size {
		t := src[:size]
		src = src[size:]
		ret = append(ret, t)
	}
	if len(src) > 0 {
		ret = append(ret, src)
	}
	return ret
}

func GetDeviceIdTail(deviceId string) int64 {
	tailNum := crc32.ChecksumIEEE([]byte(deviceId)) % 10
	return int64(tailNum)
}
