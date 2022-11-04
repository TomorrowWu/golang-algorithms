package main

import (
	"flag"
	"fmt"
)

var (
	isTest = flag.Bool("is_test", false, "is_test online ")
	num    = flag.Int64("aaa", 123, "num")
)

func main() {
	//flag.Parse()

	//fmt.Println("===", *isTest, "===")
	//fmt.Println("===", *num, "===")
	//fmt.Printf("%s\n%s\n", "===", "===")
	//
	//err := errors.New("aaaaaa33333")
	//fmt.Printf("%s\n", err)

	var a, b float64 = 0.3, 0.6
	fmt.Println(a + b)

	//FloatPrecision()
	//FloatPrecision1()
	//KahanSummation()

	//var a float32 = 1.0
	//fmt.Println(a == 1)
}

func FloatPrecision() {
	var a, b float32 = 20000000.00, 1.0
	var c float32 = a + b
	fmt.Println("c is ", c)
	var d float32 = c - a
	fmt.Println("d is ", d)
}

func FloatPrecision1() {
	var sum float32 = 0.0
	for i := 0; i < 20000000; i++ {
		var x float32 = 1.0
		sum += x
	}
	fmt.Println("sum is ", sum)
}

func KahanSummation() {
	var sum, c float32 = 0.0, 0.0
	for i := 0; i < 20000000; i++ {
		var x float32 = 1.0
		var y float32 = x - c
		var t float32 = sum + y
		c = t - sum - y
		sum = t
	}
	fmt.Println("sum is ", sum)
}
