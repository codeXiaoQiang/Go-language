package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	1.什么是正则表达式
	正则表达式是一个特殊的字符串, 用于过滤其它的字符串

	例如:
	"abcdef13554499311www.12345.com"
	小明[/byte], 明天见
	 */

	// 需求: 取出字符串中所有的123
	// 1.定义一个字符串
	str := "你好123张三123李四"
	// 2.创建一个正则表达式对象
	// 参数: 用于指定匹配的规则
	regx, _ := regexp.Compile("123")
	// 3.利用正则表达式对象, 匹配指定的字符串
	res := regx.FindAllString(str, -1)
	fmt.Println(res)
}
