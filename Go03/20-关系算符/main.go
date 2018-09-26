package main

import "fmt"

func main() {
	/*
	1.C语言中的关系运算符有
	> < >= <= == !=
	2.Go语言和C语言一样
	> < >= <= == !=
	注意点:
	在C语言中非零即真, 但是在Go语言中没有非零即真的概念
	Go语言中所有的关系运算符返回的都是true或者false

	关系运算符的优先级
	不用记, 记住()最高即可
	 */

	res1 := 10 > 5
	fmt.Println(res1) // true
	res2 := 10 < 5
	fmt.Println(res2) // false
	res3 := 10 >= 5
	fmt.Println(res3) // true
	res4 := 10 <= 5
	fmt.Println(res4) // false
	res5 := 10 == 10
	fmt.Println(res5) // true
	res6 := 10 != 10
	fmt.Println(res6) // false
}
