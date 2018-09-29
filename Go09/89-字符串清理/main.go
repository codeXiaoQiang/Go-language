package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	字符串清理
	func Trim(s string, cutset string) string
	func TrimLeft(s string, cutset string) string
	func TrimRight(s string, cutset string) string
	func TrimFunc(s string, f func(rune) bool) string
	func TrimLeftFunc(s string, f func(rune) bool) string
	func TrimRightFunc(s string, f func(rune) bool) string
	func TrimSpace(s string) string
	func TrimPrefix(s, prefix string) string
	func TrimSuffix(s, suffix string) string
	 */

	 // Trim
	 /*
	 // 作用: 返回字符串前后两端去掉指定字符串之后的新字符串
	 //str1 := "123www.it123.com123"
	 str1 := "     www.it123.com "
	 str2 := strings.Trim(str1, " ")
	 fmt.Println(str2)
	 */

	 // TrimLeft 只去除左边的
	 // TrimRight 只去除右边的
	 /*
	str1 := "123www.it123.com123"
	str2 := strings.TrimLeft(str1, "123")
	fmt.Println(str2)
	str3 := strings.TrimRight(str1, "123")
	fmt.Println(str3)
	 */

	 // TrimFunc
	 /*
	str1 := "  www.it123.com   "
	str2 := strings.TrimFunc(str1, custom)
	fmt.Println(str2)
	 */

	 // TrimSpace
	 /*
	 // 作用: 去除两端的空格
	str1 := "  www.it123.com   "
	str2 := strings.TrimSpace(str1)
	fmt.Println(str2)
	 */
	str1 := "01-HelloWorld.mp4"
	str2 := strings.TrimSuffix(str1, ".mp4")
	fmt.Println(str2)

}
func custom(ch rune) bool{
	if ch == ' '{
		return true
	}
	return false
}
