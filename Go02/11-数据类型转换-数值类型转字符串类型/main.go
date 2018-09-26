package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	1.如果将基本数据类型转换为字符串类型
	在Go语言中有两种方式可以将基本数据类型转换为字符串类型
	1.1 strconv.FormatXxx()
	1.2 strconv.Itoa()

	Itoa在C语言中也有
	 */

	 // 1.strconv.FormatXxx()
	 /*
	 //var num int = 9
	 // 第一个参数: 需要转换的整数, 必须是int64类型的
	 // 第二个参数: 转换为多少进制的字符串
	 //var str string = strconv.FormatInt(int64(num), 10) //  "9"
	 //var str string = strconv.FormatInt(int64(num), 2)  // "1001"
	 //fmt.Printf("%s\n", str)
	 */

	/*
	 //var num float32 = 3.1234567890123456789
	 // 第一个参数: 需要转换的小数, 必须是float64类型的
	 // 第二个参数: 按照什么格式转换 'f'小数格式 'e' 指数的格式
	 // 第三个参数: 保留多少位小数, 传入-1按照原始类型的精度保留
	 // 第四个参数: 原始类型的标志  float32 --> 32  float64 --> 64
	 //var str string = strconv.FormatFloat(float64(num), 'f', -1, 32) // "3.1234567"
	 //var str string = strconv.FormatFloat(float64(num), 'f', -1, 64) // "3.123456789012345"
	// var str string = strconv.FormatFloat(float64(num), 'f', 2, 32)  // "3.12"
	//fmt.Printf("%s\n", str)
	*/

	/*
	//var flag bool = false
	//var str string = strconv.FormatBool(flag) // "false"
	//fmt.Printf("%s\n", str)
	*/

	// 2.strconv.Itoa()
	//var num int = 666
	var num int64 = 666
	var str string = strconv.Itoa(int(num))
	fmt.Printf("%s\n", str)
}

