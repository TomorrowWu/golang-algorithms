package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

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

	//fmt.Println("************\n")
	//
	//var m map[int64]string
	//fmt.Println("-----------", m[1], "----------")
	//
	//fmt.Println("************\n")

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

	//fmt.Println(GetDeviceIdTail("0242C83737B3C7CDB836B9763C9EA984"))
	//fmt.Println(time.Unix(time.Now().Unix(), 0), time.Unix(time.Now().Unix(), 0).Format("01/02/2006"))
	//
	//s := `{"is_cover":"1","room_id":19170}`
	//i := struct {
	//	IsCover bool  `json:"is_cover,string"`
	//	RoomId  int64 `json:"room_id"`
	//}{}
	//if err := json.Unmarshal([]byte(s), &i); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(i)

	//args := []int{1, 2, 3}
	//args1 := append(args, 4, 5)
	//fmt.Println(args, args1)
	//
	//args1[0] = 5
	//
	//fmt.Println(args, args1)
	//
	//commentContent := fmt.Sprintf("<a id=\"%d\" stage_name=\"%s\" type=\"at\"/>%s", 1, "wuming", "commentContent")
	//
	//fmt.Println(commentContent)
	//fmt.Println(time.Now().Format("20060102"))
	//fmt.Println(time.Now().Format("200601"))
	//
	//fmt.Println(VersionCompare("", "7.7.9"))
	//
	//s := "sm://singersong?singerId=551170&name=أم كلثوم&from=0&isNewData=true"
	//fmt.Println(strings.Contains(s, "%20"))
	//s1 := strings.Replace(s, " ", "20%", -1)
	//fmt.Println(s)
	//fmt.Println(s1)
	//
	//params := url.Values{}
	//params.Add("singerId", strconv.FormatInt(123456, 10))
	//params.Add("name", "أم كلثوم")
	//params.Add("from", "0")
	//params.Add("isNewData", "true")
	//fmt.Println(params.Encode())
	//
	//// 获取当前(当地)时间
	//t := time.Now()
	//// 获取0时区时间
	//t = time.Now()
	//fmt.Println(t)
	//// 获取时区信息
	//locale, _ := time.LoadLocation("America/Los_Angeles")
	//name, offset := t.In(locale).Zone()
	//fmt.Println(name, offset/3600, offset)
	//
	//s = strings.Replace("{user_name} just liked your post💗", "{user_name}", "مرحبا", -1)
	//fmt.Println("\n=====", s, "=====")
	//s = "\u200f" + strings.Replace("{user_name} just liked your post💗", "{user_name}", "مرحبا", -1) + "\u200f"
	//fmt.Println("\n=====", s, "=====")
	//
	//fmt.Println(GetDelayTime("America/Los_Angeles", 1612701866))

	//s := fmt.Sprintf("(a.push_title like '%%%s%%' or a.push_text like '%%%s%%')", "h", "h")
	//fmt.Println(s)

	//	body := `{
	//  "Extra": null,
	//  "StatusCode": 0,
	//  "StatusMessage": "success"
	//}`
	//	tmp := &struct {
	//		StatusCode    int64  `json:"StatusCode"`
	//		StatusMessage string `json:"StatusMessage"`
	//	}{}
	//	if err := json.Unmarshal([]byte(body), &tmp); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Printf("tmp:[%+v]", tmp)

	//	body1 := `{
	//  "Extra": null,
	//  "StatusCode": 1,
	//  "StatusMessage": "success"
	//}`
	//	m := make(map[string]interface{})
	//	if err := json.Unmarshal([]byte(body1), &m); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	m["StatusCode"].
	//	if v, ok := m["StatusCode"].(float64); ok {
	//		fmt.Println("v:", v)
	//	}
	//	fmt.Printf("m:[%+v]", m)

	//t := time.Now()
	//for i := 0; i < 1500; i++ {
	//	wg.Add(1)
	//	go Test_Send()
	//}
	//wg.Wait()
	//fmt.Println("failNum:", failNum)
	//fmt.Println("succNum:", succNum)
	//fmt.Println(time.Since(t))
	//
	//time.Sleep(15 * time.Second)

	//fmt.Println(time.Now().UTC().Clock())
	//
	//sSql := fmt.Sprintf("select user_id from user_geo where geohash_value like '%s%%' limit %d", "wwbh", 5000)
	//fmt.Println(sSql)
	//
	//s := []string{"foo", "bar", "baz"}
	//s = []string{}
	//fmt.Println("[" + strings.Join(s, ", ") + "]")
	//
	//ps := []string{"foo", "bar", "baz", "123"}
	////l := len(ps)
	//rand.Seed(time.Now().UnixNano())
	//randI := rand.Int31n(2)
	//fmt.Println(randI, ps[randI:4])

	//SendFeishuTextMsg("测试消息1", []string{"ming.wu"})

	//parse, _ := url.Parse("https://m-test.starmakerstudios.com/ktv/share?app=sm&is_convert=true&room_id=7566&share_type=message&user_id=5066549353678400")
	//
	//fmt.Println(parse.String())
	//values := parse.Query()
	//fmt.Println(values)
	//fmt.Println(values.Get("user_id"))
	//fmt.Println(values.Get("uid"))
	//
	//values.Del("user_id")
	//values.Del("uid")
	//fmt.Println(values)
	//
	//values.Add("user_sid", "123")
	//fmt.Println(values)
	//
	//parse.RawQuery = values.Encode()
	//fmt.Println(parse.String())

	// 解析 referrer 参数中的，邀请人   encoded: sid%3D123%26source%3Dxxx decoded: sid=123&source=xxx
	//queryStr, err := url.QueryUnescape("sid%3D123%26source%3Dxxx")
	//if err != nil {
	//	return
	//}
	//fmt.Println(queryStr)
	//
	//params, err := url.ParseQuery(queryStr)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//for key, value := range params {
	//	fmt.Printf("  %v = %v \n", key, value)
	//}

	//type T70704Friend struct {
	//	DeviceId  string `json:"device_id"`
	//	CreatedOn int64  `json:"created_on"`
	//	Uid       int64  `json:"uid"`
	//	IsBan     bool   `json:"is_ban"`
	//}
	//var (
	//	invitedUsers string = "[{\"device_id\":\"xxxxxx\",\"created_on\":1617939082,\"uid\":7036874317767700},{\"device_id\":\"xxxxxx1\",\"created_on\":1617939082,\"uid\":7036874317767700}]"
	//	friendList   []*T70704Friend
	//)
	//if err := jsonIterator.UnmarshalFromString(invitedUsers, &friendList); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, f := range friendList {
	//	fmt.Printf("%+v \n", f)
	//}

	//fmt.Println(fmt.Sprintf("%04d", 1))

	// <a id="3634932281" stage_name="whl1001" type="at"/>123
	//s := "<a id=\"3634932281\" stage_name=\"whl1001\" type=\"at\"/>123"
	//s := `<a id="3634932281" stage_name="whl1001" type="at"/>123`
	//for _, s := range strings.Split(s, " ") {
	//	if strings.HasPrefix(s, "id=") {
	//		s = strings.TrimPrefix(s, "id=")
	//		s = strings.Trim(s, "\"")
	//		fmt.Println(s)
	//	}
	//	if strings.HasPrefix(s, "stage_name=") {
	//		s = strings.TrimPrefix(s, "stage_name=")
	//		s = strings.Trim(s, "\"")
	//		fmt.Println(s)
	//	}
	//}

	//fmt.Printf("sub 30 day,%+v ", time.Now().Add(-30*24*time.Hour))

	//fmt.Println("SELECT `status_1`,COUNT(*) num FROM `playlist` WHERE date_format(from_unixtime(`operate_time_1`), '%Y-%m-%d')=? AND country IN(" + "'IN','ID'" + ") AND `status_1` IN(4,5,6) GROUP BY `status_1`")

	//Test_Send()

	//fmt.Println(strings.TrimSpace("   aaa"))
	//fmt.Println(strings.TrimSpace(" aaa  "))
	//fmt.Println(strings.TrimSpace("      "))

	//fmt.Println(strings.Replace("Ur Facebook friend {user_platform_name} released a new playlist🎼", "{user_platform_name}", "12", -1))
	//fmt.Println(strings.Replace("{aaa}", "{aaa}", "12", -1))
	//
	//fmt.Println(strings.Replace("aa", "aa", "12", -1))
	//fmt.Println(strings.Replace("user_platform_name", "user_platform_name", "12", -1))
	//
	//fmt.Println(strings.Replace("Ur Facebook friend {user_platform_name} released a new playlist🎼", "{user_platform_name}", "12", -1))

	// https://improxy.starmakerstudios.com/tools/im/360x/stream_pic/20210620/4222124654367227_226137137_004846_10.jpg
	//s := strings.Split("https://improxy.starmakerstudios.com/tools/im/360x/stream_pic/20210620/4222124654367227_226137137_004846_10.jpg", "/stream_pic/")
	//fmt.Println(s)
	//
	//for _, p := range s {
	//	fmt.Println(p)
	//}

	//d := time.Duration(15*60) * time.Second
	//fmt.Println(d.Seconds())

	var arr []*time.Duration
	fmt.Println(arr, arr == nil)
	
	json.Unmarshal([]byte("[{}]"),&arr)
	fmt.Println(arr, arr == nil)
	
	//arr = []*time.Duration{}
	//fmt.Println(arr, arr == nil)
	
	fmt.Println(fmt.Sprintf("select user_id from user_geo where geohash_value like '%s%%' limit %d", "xxx", 5000))

}

type feishuMsg struct {
	Touser []string `json:"touser"`
	Id     int64    `json:"id"`
	Data   struct {
		Msg string `json:"msg"`
	} `json:"data"`
	Type string `json:"type"`
}

// SendFeishuTextMsg 通过 飞书应用 发消息给指定飞书用户   mentionedList: 企业邮箱前缀 ming.wu 比如 ming.wu@ushow.media
func SendFeishuTextMsg(msg string, mentionedList []string) {
	if len(strings.TrimSpace(msg)) == 0 || len(mentionedList) == 0 {
		return
	}

	sendMsg := &feishuMsg{
		Touser: mentionedList,
		Id:     6,
		Data: struct {
			Msg string `json:"msg"`
		}{
			Msg: msg,
		},
		Type: "Text",
	}

	tMsg, _ := json.Marshal(sendMsg)
	resp, err := http.Post("https://devops.ushow.media/devops-goserver-v1/notify/message/", "application/json", bytes.NewBuffer(tMsg))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//tlog.Info(string(body))

	tmp := &struct {
		Code int64  `json:"code"`
		Data string `json:"data"`
	}{}
	if err := json.Unmarshal(body, &tmp); err != nil {
		fmt.Println(err, string(body))
		return
	}
	if tmp.Code != http.StatusOK {
		fmt.Println(string(body))
		return
	}
}

func GetDelayTime(timezone string, t int64) int64 {
	locale, _ := time.LoadLocation(timezone)
	_, offset := time.Now().In(locale).Zone()
	return t - int64(offset) - time.Now().Unix()
}

func VersionCompare(v1, v2 string) int8 {
	v1s := strings.Split(v1, ".")
	v2s := strings.Split(v2, ".")
	i := 0
	for ; i < len(v1s) && i < len(v2s); i++ {
		v1i, err1 := strconv.ParseInt(v1s[i], 10, 64)
		if err1 != nil {
			return -1
		}
		v2i, err2 := strconv.ParseInt(v2s[i], 10, 64)
		if err2 != nil {
			return 1
		}
		if v1i > v2i {
			return 1
		} else if v1i < v2i {
			return -1
		}
	}
	if i < len(v1s) {
		return 1
	}
	if i < len(v2s) {
		return -1
	}
	return 0
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

var clientChannel chan *http.Client
var wg sync.WaitGroup
var failNum int64
var succNum int64

func init() {
	// 启动HTTP/2协议

	clientChannel = make(chan *http.Client, 10)
	for i := 0; i < 1; i++ {
		clientChannel <- &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:          500,
				MaxIdleConnsPerHost:   500,
				MaxConnsPerHost:       500,
				ResponseHeaderTimeout: 60 * time.Second,

				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	}
}

func getClient() *http.Client {
	client := <-clientChannel
	clientChannel <- client
	return client
}

// FcmMsg represents fcm response message - (tokens and topics)
type FcmResponseStatus struct {
	Ok            bool
	StatusCode    int
	MulticastId   int64               `json:"multicast_id"`
	Success       int                 `json:"success"`
	Fail          int                 `json:"failure"`
	Canonical_ids int                 `json:"canonical_ids"`
	Results       []map[string]string `json:"results,omitempty"`
	MsgId         int64               `json:"message_id,omitempty"`
	Err           string              `json:"error,omitempty"`
	RetryAfter    string
}

// NotificationPayload notification message payload
type NotificationPayload struct {
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Sound        string `json:"sound,omitempty"`
	Badge        string `json:"badge,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Color        string `json:"color,omitempty"`
	ClickAction  string `json:"click_action,omitempty"`
	BodyLocKey   string `json:"body_loc_key,omitempty"`
	BodyLocArgs  string `json:"body_loc_args,omitempty"`
	TitleLocKey  string `json:"title_loc_key,omitempty"`
	TitleLocArgs string `json:"title_loc_args,omitempty"`
}

// FcmMsg represents fcm request message
type FcmMsg struct {
	Data                  interface{}         `json:"data,omitempty"`
	To                    string              `json:"to,omitempty"`
	RegistrationIds       []string            `json:"registration_ids,omitempty"`
	CollapseKey           string              `json:"collapse_key,omitempty"`
	Priority              string              `json:"priority,omitempty"`
	Notification          NotificationPayload `json:"notification,omitempty"`
	ContentAvailable      bool                `json:"content_available,omitempty"`
	DelayWhileIdle        bool                `json:"delay_while_idle,omitempty"`
	TimeToLive            int                 `json:"time_to_live,omitempty"`
	RestrictedPackageName string              `json:"restricted_package_name,omitempty"`
	DryRun                bool                `json:"dry_run,omitempty"`
	Condition             string              `json:"condition,omitempty"`
}

func Test_Send() {
	//defer wg.Done()
	message := new(FcmMsg)
	message.RegistrationIds = []string{"fcZ_J0OgTb6k4ka8Qgbv_U:APA91bGgGMV3B9l3-1xqN-9FX5d80MlOVhp4vVlMsERq5CA9JnmnNN0cCTAmm-PzYA1zY8Pt7HwzXM5OPFfE4nAv1ZjAfNeIgZF6DcObr9W5DvasHGQ30T94hCdbJkDIytQww9aakdXm"}
	message.Data = map[string]interface{}{
		"message": map[string]interface{}{
			"id":               "9001",
			"title":            "wuming555",
			"text":             time.Now().String(),
			"time":             time.Now().Unix(),
			"action_url":       "sm://playrecording_fullscreen?sm_id=1204839388",
			"sound":            0,
			"kind":             "content",
			"target_user_id":   "7036874417786638",
			"media_type":       1,
			"media_url":        "https://improxy.starmakerstudios.com/tools/im/560/production/uploading/recordings/579625246/cover_image.png?ts=1611949680",
			"unique_id":        time.Now().UnixNano(),
			"show_time":        time.Now().Unix() + 5,
			"media_style":      "",
			"priority":         0,
			"source":           "gcm",
			"r_info":           `{"anonymous_user_id":"126","bucket_id":"0","country":"ID","device_id":"b936a9a602584500a0d654e33628f125","gender":0,"is_buy_sm_id":"","is_new_user":"","item_id":"1204839388","labels":[],"lang":"en","msg_type":"","region":"Area_ID","s_ab_type":0,"s_country":"ID","s_editor":"yuqi.zhang","s_item_type":3,"s_push_id":"55047","s_push_time":"2021-06-11 16:20","s_recording_id":"579625246","s_referrer_name":"","s_region":"Area_ID","s_reviewer":"ruihan.wang","s_show_type":"0","s_source":"126","s_token":"9a7f7de4a33dfee3cfff2298f1e3d7a558e31a1c444b9459f10f9197fd3cf083","s_token_version":"8.14.1","scene":"content_push_v17","trace_id":"16233927606239313108396","user_id":"7036874417786638"}`,
			"trigger_type":     0,
			"trigger_sub_type": 0,
			"push_type":        "content",
			"valid_begin":      0,
			"valid_end":        2623400505,
		},
		"type": "command",
		"command": map[string]interface{}{
			"kind": "ddns",
			"ddns": map[string]interface{}{
				"config_host": "push.tspot.vip",
			},
		},
	}
	message.Priority = "high"

	jsonByte, _ := jsonIterator.Marshal(message)

	fmt.Println(string(jsonByte))

	request, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(jsonByte))
	apiKey := "AAAA42HccNU:APA91bFCGZF0Auy0XoVpNXaOGg8T0aNG1TfFJTL2cFZ3DvXS6-8emAoW65xj06UcHj5x1ph85PKG2vxskHvfWFe_-usA6FPVqf-Bl1TPu4NqcN54eoYDzN14cUvLfABSaq_5MXENACGb"
	request.Header.Set("Authorization", "key="+apiKey)
	request.Header.Set("Content-Type", "application/json")

	client := getClient()
	response, err := client.Do(request)
	if err != nil {
		//atomic.AddInt64(&failNum, 1)
		fmt.Println("v:", err)
		return
	}
	fmt.Println("http协议:", request.Proto, response.Proto)
	//atomic.AddInt64(&succNum, 1)
	defer response.Body.Close()
}
