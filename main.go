package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
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

	//var arr []*time.Duration
	//fmt.Println(arr, arr == nil)
	//
	//json.Unmarshal([]byte("[{}]"), &arr)
	//fmt.Println(arr, arr == nil)

	//arr = []*time.Duration{}
	//fmt.Println(arr, arr == nil)

	//fmt.Println(fmt.Sprintf("select user_id from user_geo where geohash_value like '%s%%' limit %d", "xxx", 5000))

	//AddMobiMaterial()

	//fmt.Printf("===%s===", strings.Join([]string{}, ","))

	//Url, _ := url.Parse("sm://facebook_content_deferred_link")
	//parameters := url.Values{}
	//parameters.Add("ad_type", "2")
	//parameters.Add("show_guide", "0")
	//parameters.Add("language_code", "en")
	//parameters.Add("sm_id", "123")
	//parameters.Add("first_start_tab", "sing")
	//parameters.Add("uid", strconv.FormatInt(1, 10))
	//parameters.Add("song_id", "123")
	//parameters.Add("mid", strconv.FormatInt(1, 10))
	//parameters.Add("channel_extra", `["country:us","genre:pop"]`)
	//Url.RawQuery = parameters.Encode()

	//fmt.Println(Url.String())

	//url1, _ := url.Parse("sm://appsflyer_content_onelink_link")
	//parameters1 := url.Values{}
	//parameters1.Add("language_code", "en")
	//parameters1.Add("song_id", "123")
	//parameters1.Add("channel_extra", `["country:us","genre:pop"]`)
	//parameters1.Add("df_referrer", "")
	//url1.RawQuery = parameters1.Encode()

	//deeplink := "https://starmaker.onelink.me/xzo7?pid=TT&deep_link_value=sm://appsflyer_content_onelink_link?language_code=%s&song_id=%s&channel_extra=%s&c=tt-deeplink&is_retargeting=true&af_click_lookback=7d"
	//deeplink = fmt.Sprintf(deeplink, contentLan, songId, extra)

	//url2, _ := url.Parse("https://starmaker.onelink.me/xzo7")
	//parameters2 := url.Values{}
	//parameters2.Add("pid", "TT")
	//parameters2.Add("deep_link_value", url1.String())
	//parameters2.Add("c", "tt-deeplink")
	//parameters2.Add("is_retargeting", "true")
	//parameters2.Add("af_click_lookback", "7d")
	//url2.RawQuery = parameters2.Encode()

	//fmt.Println(url2.String(), len(url2.String()))

	//fmt.Println(url.QueryEscape("https://starmaker.onelink.me/xzo7?af_click_lookback=7d&c=tt-deeplink&is_retargeting=true&pid=TT&deep_link_value=sm://appsflyer_content_onelink_link?channel_extra=[\"country:us\",\"genre:pop\"]&df_referrer=&language_code=en&song_id=123"))

	//enEscapeUrl, _ := url.QueryUnescape("ONEPLUS%20A5000")
	//fmt.Println(enEscapeUrl)

	//deepLinkValue := fmt.Sprintf("sm://appsflyer_content_onelink_link?channel_extra=%s&language_code=%s&song_id=%s", url.QueryEscape(`["country:us","genre:pop"]`), "en", "123")
	//s := fmt.Sprintf("https://starmaker.onelink.me/xzo7?af_click_lookback=7d&c=tt-deeplink&is_retargeting=true&pid=TT&deep_link_value=%s", url.QueryEscape(deepLinkValue))
	//fmt.Println(s)

	//fmt.Println(url.QueryEscape(`["country:us","genre:pop"]`))
	//
	//unescape, _ := url.QueryUnescape(url.QueryEscape(`["country:us","genre:pop"]`))
	//unescape, _ = url.QueryUnescape(unescape)
	//fmt.Println(unescape)
	//
	//fmt.Println("=" + strings.Join([]string{}, ",") + "=")
	//
	//ssql := fmt.Sprintf("select m_id, sm_id, nv, source, video_path, hook_path,ext, `type`, recall_score from sm_volume_bought "+
	//	"where dt=? and (language=? or language like '%s-%%') and country=? and t=? order by recall_score desc", "en")
	//fmt.Println(ssql)
	//
	//offset := map[string]int64{
	//	"IN": 5.5 * 3600, // UTC+5:30
	//}
	//fmt.Println(offset["IN"])

	//fn := func() {
	//	fmt.Println("test")
	//}
	//Pulse(fn, time.Second)

	//time.Sleep(time.Minute)

	//var arr []int64
	//_ = json.Unmarshal([]byte("[1,2]"), &arr)
	//fmt.Println(arr)

	//ssql := fmt.Sprintf("SELECT `id` FROM `text_tpl_class` WHERE `tmp_ids` LIKE '%%\"tpl_id\":%d%%'", 1)
	//fmt.Println(ssql)

	//dir, _ := os.Getwd()
	//dirName := "text_template_" + fmt.Sprintf("%d_%d", 1, time.Now().UnixMicro())
	//tmpDir := path.Join(dir, dirName)
	//_ = os.MkdirAll(tmpDir, os.ModePerm)
	//
	//localJsonFile := path.Join(tmpDir, "tpl.json")
	//_ = os.WriteFile(localJsonFile, []byte("{}"), os.ModePerm)
	//out, err := exec.Command("zip", "-r", dirName+".zip", dirName).Output()
	//fmt.Println(out, err, dir)

	//zipDir := path.Join(dir, dirName+".zip")
	//zipFile, _ := os.Create(zipDir)
	//defer zipFile.Close()
	//zipwriter := zip.NewWriter(zipFile)
	//defer zipwriter.Close()
	//zipwriter.Create(dirName + "/")
	//defer func() {
	//	if err := os.RemoveAll(tmpDir); err != nil {
	//	}
	//}()

	//type tmp struct {
	//	Id int64 `json:"id"`
	//}
	//var t = new(tmp)
	//jsonIterator.UnmarshalFromString(`{"id":1}`, t)
	//fmt.Println(t)

	//m := make(map[int][]int64)
	//m[0] = append(m[0], 123)
	//fmt.Println(m)

	data := map[string]map[string]int64{}
	jsonIterator.UnmarshalFromString(`{"7036874317771764":{"background_effects":1665196471,"box_background":1665196471,"box_pendant":1665309742,"profile_pendant":1665198607,"top_part_background":1665470057}}`, &data)
	fmt.Println(data)

	data["1"] = map[string]int64{}

	s, _ := jsonIterator.MarshalToString(data)
	fmt.Println(s)

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

func AddMobiMaterial() {
	// https://ads-admin.static.kunlun.com/api/material?product=starmaker&page=0&per_page=10

	type tmp struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List []struct {
				Id       int64  `json:"id"`
				Country  string `json:"country"`
				Language string `json:"language"`
				Url      string `json:"url"`
				Typ      int64  `json:"type"` // 1是图片，2是视频
			} `json:"list"`
		} `json:"data"`
	}

	var page int64 = 0
	for {
		p := "starmaker"
		d := time.Now().Add(-1 * 24 * time.Hour).Format("2006-01-02")
		pageS := strconv.FormatInt(page, 10)
		pageSize := "50" // 最多50条

		req, _ := http.NewRequest(http.MethodGet, "https://ads-admin.static.kunlun.com/api/material", nil)
		q := req.URL.Query()
		q.Add("start_date", d)
		q.Add("page", pageS)
		q.Add("per_page", pageSize)
		q.Add("product", p)
		q.Add("api_key", "9A726A18BEE7A5C148B80DA0490D2E59")
		req.URL.RawQuery = q.Encode()

		//param := fmt.Sprintf("page=%s&per_page=%s&product=%s&start_date=%s&api_key=9A726A18BEE7A5C148B80DA0490D2E59", pageS, pageSize, p, d)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer 8oR+VRlRGXewaag+hiy4QQ")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("err:%+v", err)
			return
		}

		respBody, err := io.ReadAll(resp.Body)
		fmt.Printf("req:%s,body:%s,err:%+v\n", req.URL.RawQuery, string(respBody), err)

		ret := new(tmp)
		jsonIterator.Unmarshal(respBody, ret)
		_ = resp.Body.Close()

		if len(ret.Data.List) > 0 {
			fmt.Printf("page:%d\n", page)

			//TODO 吴名 2021/10/28 21:17 写入到 material 表(去掉已有的)

			// 还有下一页
			page++
		} else {
			fmt.Printf("stop\n")
			break
		}
	}
}

func Pulse(fn func(), duration time.Duration) {
	fn()
	go func() {
		for range time.NewTicker(duration).C {
			func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Printf("recover err:%v\n", err)
					}
				}()
				fn()
				panic("panic")
			}()
		}
	}()
}
