package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	1.从字符串中提取电话号码
	2.从字符串中提取邮箱
	 */


	str := "今天天气很好97606813@qq.cn你吃饭了吗?linanjiang@it666.com"
	// 1.创建正则表达式对象
	// [a-zA-Z0-9_] 只要是字母或者数字或者_就行
	// + 前面一个规则匹配1次或多次, 越多越好
	// @ 原样匹配
	// \\.  代表将.转义为普通字符.
	regx, _ := regexp.Compile("[a-zA-Z0-9_]+@[a-zA-Z0-9]+\\.[a-zA-Z]+")
	// 2.利用正则表达式提取符合需求的字符串
	res := regx.FindAllString(str, -1)
	fmt.Println(res)
}

func test()  {
	str := "李abc13554499311def江15812345678def"
	// 1.创建一个正则表达式对象
	// [0123456789] 对应位可以是0~9任意的一个
	// {10}将前面一个条件连续匹配10次
	regx, err := regexp.Compile("1\\d{10}")
	if err != nil{
		fmt.Println(err)
	}
	res := regx.FindAllString(str, -1)
	fmt.Println(res)
}
