package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	字符串拆合
	1.字符串切割
	func Split(s, sep string) []string
	func SplitN(s, sep string, n int) []string
	func SplitAfter(s, sep string) []string
	func SplitAfterN(s, sep string, n int) []string
	2.按照空格切割字符串
	func Fields(s string) []string
	func FieldsFunc(s string, f func(rune) bool) []string
	3.字符串合并
	func Join(a []string, sep string) string

	4.重复生成字符串
	func Repeat(s string, count int) string

	5.替换字符串
	func Replace(s, old, new string, n int) string
	 */

	 // Split
	 /*
	 // 作用: 按照指定字符串切割原有字符串
	 // 注意点: 切割之后的结果, 不包含指定的字符串
	 str := "abc,def,123,456"
	 sce := strings.Split(str, ",")
	 fmt.Println(sce)
	 */

	 // SplitN
	 /*
	// 作用: 按照指定字符串切割原有字符串, 切割为指定的份数
	// 注意点: 切割之后的结果, 不包含指定的字符串
	str := "abc,def,123,456"
	//sce := strings.SplitN(str, ",", 2)
	sce := strings.SplitN(str, ",", 0)
	fmt.Println(len(sce))
	fmt.Println(sce)
	 */

	 // SplitAfter
	 /*
	// 作用: 按照指定字符串切割原有字符串
	// 注意点: 切割之后的结果, 包含指定的字符串
	str := "abc,def,123,456"
	sce := strings.SplitAfter(str, ",")
	fmt.Println(len(sce))
	fmt.Println(sce)
	 */

	 // SplitAfterN
	 /*
	// 作用: 按照指定字符串切割原有字符串, 切割为指定的份数
	// 注意点: 切割之后的结果, 包含指定的字符串
	str := "abc,def,123,456"
	sce := strings.SplitAfterN(str, ",", 2)
	fmt.Println(len(sce))
	fmt.Println(sce)
	 */

	 // Fields
	 /*
	// 作用: 按照空格切割原有字符串
	// 注意点: 连续的多个空格会按照一个空格来处理
	str := "abc def        123 456"
	sce := strings.Fields(str)
	fmt.Println(len(sce))
	fmt.Println(sce)
	 */

	 // FieldsFunc
	 /*
	 // 作用: 函数返回true就切片, 返回false就不切割
	//str := "abc def       123 456"
	str := "abc-def-123-456"
	sce := strings.FieldsFunc(str, custom)
	fmt.Println(len(sce))
	fmt.Println(sce)
	 */

	// Join
	/*
	// 作用: 按照指定字符连接切片中的元素
	sce := []string{"www", "it666", "com"}
	str := strings.Join(sce, ".")
	fmt.Println(str)
	*/

	// Repeat
	/*
	// 作用: 将原有字符串重复指定次数后生成一个新的字符串
	str1 := "123"
	str2 := strings.Repeat(str1, 2)
	fmt.Println(str2)
	*/

	// Replace
	// 作用: 将原有字符串中的, 指定字符串替换为新的字符串
	//      最后一个参数用于指定替换多少个, 如果传入-1全部都替换
	str1 := "abc123abc456def"
	//str2 := strings.Replace(str1,"abc","www", 1)
	//str2 := strings.Replace(str1,"abc","www", 2)
	str2 := strings.Replace(str1,"abc","www", -1)
	fmt.Println(str2)
}
func custom(ch rune) bool{
	//if ch == ' '{
	//	return true
	//}
	if ch == '-'{
		return true
	}
	return false
}