package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	字符串包含
	func Contains(s, substr string) bool
	func ContainsRune(s string, r rune) bool
	func ContainsAny(s, chars string) bool
	func HasPrefix(s, prefix string) bool
	func HasSuffix(s, suffix string) bool
	 */

	// Contains
	/*
	// 作用: 判断字符串中是否包含指定的子串
	// 如果包含返回true, 如果不包含返回false
	str := "www.it666.江南.com"
	//res := strings.Contains(str, "66")
	//res := strings.Contains(str, "mm")
	res := strings.Contains(str, "江南")
	fmt.Println(res)
	*/

	// ContainsRune
	/*
   // 作用: 判断字符串中是否包含指定的字符
   // 如果包含返回true, 如果不包含返回false
   str := "www.it666.江南.com"
   //res := strings.ContainsRune(str, 't')
   //res := strings.ContainsRune(str, 'p')
   res := strings.ContainsRune(str, '江')
   fmt.Println(res)
	*/

	// ContainsAny
	/*
	// 作用: 判断字符串中是否包含指定的子串中任意一个字符
	// 如果包含返回true, 如果不包含返回false
	str := "www.it666.江南.com"
	res := strings.ContainsAny(str, "p江e")
	fmt.Println(res)
	*/

	// HasPrefix
	// 作用: 判断字符串是否以指定字符串开头
	str := "01-够浪HtlloWorld.mp4"
	//res := strings.HasPrefix(str, "01")
	// HasSuffix
	// 作用: 判断字符串是否以指定字符串结尾
	res := strings.HasSuffix(str, ".mp4")
	fmt.Println(res)
}
