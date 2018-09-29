package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	字符串转换
	unc ToUpper(s string) string
	func ToLower(s string) string
	func ToTitle(s string) string
	func ToUpperSpecial(_case unicode.SpecialCase, s string) string
	func ToLowerSpecial(_case unicode.SpecialCase, s string) string
	func ToTitleSpecial(_case unicode.SpecialCase, s string) string
	func Title(s string) string
	 */

	 // ToUpper
	 /*
	 // 作用: 将指定字符串中所有的小写字母转换成大写
	 // 注意点: 会生成一个新的字符串, 不会影响原有的字符串
	 str1 := "www.it666.com"
	 str2 := strings.ToUpper(str1)
	 fmt.Println(str1)
	 fmt.Println(str2)
	 */

	 // ToLower
	 /*
	// 作用: 将指定字符串中所有的大写字母转换成小写
	// 注意点: 会生成一个新的字符串, 不会影响原有的字符串
	str1 := "WWW.IT666.COM"
	str2 := strings.ToLower(str1)
	fmt.Println(str1)
	fmt.Println(str2)
	 */

	 // ToTitle 和 ToUpper 会将小写转换为大写
	 /*
	str1 := "hello world"
	str2 := strings.ToTitle(str1)
	fmt.Println(str1)
	fmt.Println(str2)
	 */


	 // Title
	 // 作用: 将单词的首字母大写, 单词之间用空格或者-隔开
	//str1 := "hello world"
	//str1 := "helloworld"
	str1 := "hello-world"
	str2 := strings.Title(str1)
	fmt.Println(str1)
	fmt.Println(str2)
}
