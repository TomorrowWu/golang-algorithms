package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	stocks := load("/Users/wuming/Developer/go/src/golang-algorithms/us_stock/ipo/data.json")

	// 策略1：首日开盘买入，第6交易日开盘卖出
	var profit1 float64
	// 策略2：首日收盘买入，第6交易日开盘卖出
	var profit2 float64
	// 策略3：首日开盘买入，第5交易日开盘卖出
	var profit3 float64
	// 策略4：首日收盘买入，第5交易日开盘卖出
	var profit4 float64
	// 策略5：首日开盘买入，第4交易日开盘卖出
	var profit5 float64
	// 策略6：首日收盘买入，第4交易日开盘卖出
	var profit6 float64
	// 策略7：首日开盘买入，首日高开低走第2交易日开盘卖出，剩下的第6交易日开盘卖出
	var profit7 float64
	// 策略8：首日开盘买入，第2交易日开盘价<首日收盘价,第2交易日开盘卖出，剩下的第6交易日开盘卖出
	var profit8 float64
	// 策略9：首日开盘买入，第3交易日开盘价<首日收盘价,第3交易日开盘卖出，剩下的第6交易日开盘卖出
	var profit9 float64
	// 策略10：首日开盘买入，第4交易日开盘价<首日收盘价,第4交易日开盘卖出，剩下的第6交易日开盘卖出
	var profit10 float64
	// 策略11：首日开盘买入，第5交易日开盘价<首日收盘价,第5交易日开盘卖出，剩下的第6交易日开盘卖出
	var profit11 float64
	// 策略12：首日开盘买入，第3交易日开盘卖出
	var profit12 float64
	// 策略13：首日开盘买入，第6交易日开盘卖出,首日开盘价比招股价高80%不买
	var profit13 float64
	for _, s := range stocks {
		if s.IpoPrice <= 0 {
			continue
		}
		profit1 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		profit2 += 10000.00/s.Day1EndPrice*s.Day6OpenPrice - 10000
		profit3 += 10000.00/s.Day1OpenPrice*s.Day5OpenPrice - 10000
		profit4 += 10000.00/s.Day1EndPrice*s.Day5OpenPrice - 10000
		profit5 += 10000.00/s.Day1OpenPrice*s.Day4OpenPrice - 10000
		profit6 += 10000.00/s.Day1EndPrice*s.Day4OpenPrice - 10000

		if s.Day1EndPrice < s.Day1OpenPrice {
			profit7 += 10000.00/s.Day1OpenPrice*s.Day2OpenPrice - 10000
		} else {
			profit7 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}

		if s.Day2OpenPrice < s.Day1EndPrice {
			profit8 += 10000.00/s.Day1OpenPrice*s.Day2OpenPrice - 10000
		} else {
			profit8 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}

		if s.Day3OpenPrice < s.Day1EndPrice {
			profit9 += 10000.00/s.Day1OpenPrice*s.Day3OpenPrice - 10000
		} else {
			profit9 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}

		if s.Day4OpenPrice < s.Day1EndPrice {
			profit10 += 10000.00/s.Day1OpenPrice*s.Day4OpenPrice - 10000
		} else {
			profit10 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}

		if s.Day5OpenPrice < s.Day1EndPrice {
			profit11 += 10000.00/s.Day1OpenPrice*s.Day5OpenPrice - 10000
		} else {
			profit11 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}

		profit12 += 10000.00/s.Day1OpenPrice*s.Day3OpenPrice - 10000

		if s.Day1OpenPrice/s.IpoPrice <= 1.8 {
			profit13 += 10000.00/s.Day1OpenPrice*s.Day6OpenPrice - 10000
		}
	}
	fmt.Printf("策略1：首日开盘买入，第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit1, len(stocks))
	fmt.Printf("策略2：首日收盘买入，第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit2, len(stocks))
	fmt.Printf("策略3：首日开盘买入，第5交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit3, len(stocks))
	fmt.Printf("策略4：首日收盘买入，第5交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit4, len(stocks))
	fmt.Printf("策略5：首日开盘买入，第4交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit5, len(stocks))
	fmt.Printf("策略6：首日收盘买入，第4交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit6, len(stocks))
	fmt.Printf("策略7：首日开盘买入，首日高开低走第2交易日开盘卖出，剩下的第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit7, len(stocks))
	fmt.Printf("策略8：首日开盘买入，第2交易日开盘价<首日收盘价,第2交易日开盘卖出，剩下的第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit8, len(stocks))
	fmt.Printf("策略9：首日开盘买入，第3交易日开盘价<首日收盘价,第3交易日开盘卖出，剩下的第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit9, len(stocks))
	fmt.Printf("策略10：首日开盘买入，第4交易日开盘价<首日收盘价,第4交易日开盘卖出，剩下的第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit10, len(stocks))
	fmt.Printf("策略11：首日开盘买入，第5交易日开盘价<首日收盘价,第5交易日开盘卖出，剩下的第6交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit11, len(stocks))
	fmt.Printf("策略12：首日开盘买入，第3交易日开盘卖出,每单购买1w美金，总收益:%.2f,股票数:%d\n", profit12, len(stocks))
	fmt.Printf("策略13：首日开盘买入，第6交易日开盘卖出,首日开盘价比招股价高80%%不买，总收益:%.2f,股票数:%d\n", profit13, len(stocks))
}

type Stock struct {
	Code          string  `json:"code"`
	IpoPrice      float64 `json:"ipo_price"` // 招股价
	Day1OpenPrice float64 `json:"day1_open_price"`
	Day1EndPrice  float64 `json:"day1_end_price"`
	Day2OpenPrice float64 `json:"day2_open_price"`
	Day2EndPrice  float64 `json:"day2_end_price"`
	Day3OpenPrice float64 `json:"day3_open_price"`
	Day3EndPrice  float64 `json:"day3_end_price"`
	Day4OpenPrice float64 `json:"day4_open_price"`
	Day4EndPrice  float64 `json:"day4_end_price"`
	Day5OpenPrice float64 `json:"day5_open_price"`
	Day5EndPrice  float64 `json:"day5_end_price"`
	Day6OpenPrice float64 `json:"day6_open_price"`
	Day6EndPrice  float64 `json:"day6_end_price"`
	//Day7OpenPrice float64 `json:"day7_open_price"`
	//Day7EndPrice  float64 `json:"day7_end_price"`
}

func load(filename string) (ret []*Stock) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
