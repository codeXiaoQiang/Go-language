package main

import "fmt"

func main() {
	/*
	1.C语言中的逻辑运算符有
	&& || !
	2.Go语言和C语言一样
	&& || !

	注意点:
	1.Go语言中的逻辑运算符返回的是true和false
    2.参与运算的表达式的返回值都必须是布尔类型才行
	表达式1 && 表达式2
	表达式1 || 表达式2
	!表达式

	规则:
	&& 一假则假
	|| 一真则真
	! 真变假假变真
	 */

	 res1 := (10 > 5) && (1 > 1)
	 //res1 := 1 && 1 // 报错, 没有非零即真的概念, &&两遍必须都是布尔类型
	 fmt.Println(res1) // false

	res2 := (10 > 5) || (1 > 1)
	//res2 := 1 || 1 // 报错, 没有非零即真的概念, ||两遍必须都是布尔类型
	fmt.Println(res2) // true

	//res3 := !false
	//res3 := !0 // 报错, 没有非零即真的概念, !右必须是布尔类型\
	res3 := !(10 > 5)
	fmt.Println(res3)
}
