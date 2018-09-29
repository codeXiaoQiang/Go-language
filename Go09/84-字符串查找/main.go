package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	查找子串在字符串中出现的位置
	1.从左至右查找
	func IndexByte(s string, c byte) int
	func IndexRune(s string, r rune) int
	func IndexAny(s, chars string) int
	func Index(s, sep string) int
	func IndexFunc(s string, f func(rune) bool) int
	2.从右至左查找
	func LastIndex(s, sep string) int
	func LastIndexByte(s string, c byte) int
	func LastIndexAny(s, chars string) int
	func LastIndexFunc(s string, f func(rune) bool) int
	 */

	 // IndexByte
	 /*
	 // 1.定义一个字符串
	 //str := "www.it666.com"
	 str := "www.知播渔.com"
	 // 2.查找指定字符在字符串中的位置
	 // 注意点: 从左至右的查找, 一旦找到了, 就不会继续查找的
	 //         只能查找字符, 不能查找中文
	 //         如果找到了就返回字符在字符串中的位置, 如果没有找到就会-1
	 //index := strings.IndexByte(str, '6')
	 //index := strings.IndexByte(str, 'p')
	 index := strings.IndexByte(str, '知')
	 // 3.打印查找到的位置
	 fmt.Println(index)
	 */

	 // IndexRune
	 /*
	// 1.定义一个字符串
	str := "www.知播渔.com"
	// 2.查找指定字符在字符串中的位置
	// 注意点: 从左至右的查找, 一旦找到了, 就不会继续查找的
	//         可以查找字符, 也可以查找中文
	//         如果找到了就返回字符在字符串中的位置, 如果没有找到就会-1
	//         在返回找到的索引的时候, 如果前面有中文, 那么一个中文按照3个索引计算
	//index := strings.IndexRune(str, '知') // 4
	//index := strings.IndexRune(str, '播') // 7
	index := strings.IndexRune(str, 'w')
	// 3.打印查找到的位置
	fmt.Println(index)
	 */

	 // IndexAny
	 /*
	// 1.定义一个字符串
	//str := "www.it666.com"
	str := "www.李南江.com"
	// 2.查找指定字符在字符串中的位置
	// 注意点: 从左至右的查找, 一旦找到了, 就不会继续查找的
	//         查找时不是将第二个参数当做一个整体查找, 是拆开来查找, 返回找到的最前面一个值
	//index := strings.IndexAny(str, "6ic")
	index := strings.IndexAny(str, "江南")
	// 3.打印查找到的位置
	fmt.Println(index)
	 */

	 // Index
	 /*
	// 1.定义一个字符串
	//str := "www.it666.it"
	str := "www.江南666.江南"
	// 2.查找指定字符在字符串中的位置
	// 注意点: 从左至右的查找, 一旦找到了, 就不会继续查找的
	//         查找时会将第二个参数当做一个整体查找
	//index := strings.Index(str, "it")
	index := strings.Index(str, "江南")
	// 3.打印查找到的位置
	fmt.Println(index)
	 */

	 // IndexFunc
	 /*
	// 1.定义一个字符串
	//str := "www.it666.it.com"
	str := "www.江南666.江南.com"
	// 2.查找指定字符在字符串中的位置
	// 注意点: 从左至右的查找, 一旦找到了, 就不会继续查找的
	// IndexFunc方法会将字符串转换成一个[]rune, 然后遍历切片,
	// 逐个将切片中的元素传递给自定义的函数, 只要自定义函数返回true就代表找到了
	index := strings.IndexFunc(str, custom)
	// 3.打印查找到的位置
	fmt.Println(index)
	 */


	 // LastIndex
	// 1.定义一个字符串
	str := "www.it666.it.com"
	// 2.查找指定字符在字符串中的位置
	// 注意点: 从右至左的查找, 一旦找到了, 就不会继续查找的
	index := strings.LastIndex(str, "it")
	// 3.打印查找到的位置
	fmt.Println(index)
}
func custom(ch rune) bool {
	if ch == '江'{
		// 一旦return的是true就代表告诉IndexFunc方法, 找到了我们需要查找的字符
		return true
	}
	return false
}
