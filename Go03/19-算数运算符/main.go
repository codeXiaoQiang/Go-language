package main

import "fmt"

func main() {
	/*
	1.C语言中算数运算符有
	+ - * / %  ++ --
	2.Go语言中的算数运算符和C语言一样
	+ - * / %  ++ --
	并且用法都一样
	 */
	 /*
	 res1 := 1 + 1
	 fmt.Println(res1) // 2
	 res2 := 1 - 1
	 fmt.Println(res2) // 0
	 res3 := 2 * 2
	 fmt.Println(res3) // 4
	 res4 := 10 / 3
	 fmt.Println(res4) // 3
	 res5 := 10 % 3
	 fmt.Println(res5) // 1

	 // 注意点:
	 // 不同类型的常量(字面量)可以进行运算
	 // 不同类型的变量不可以直接进行运算
	 //res6 := 10 / 3.0
	 var num1 int  = 10
	 var num2 float64 = 3.0
	 //var res6 float64 = num1 / num2 // 报错
	 var res6 float64 = float64(num1) / num2
	 fmt.Println(res6)

	 // 注意点:
	 // + 还可以用于拼接字符串
	 var str1 string = "lnj"
	 var str2 string = "it666"
	 var res7 string = str1 + str2
	 fmt.Println(res7)
	 // 只有相同类型才能进行运算
	 //var res8 string = "lnj" + 7 // 报错
	 //fmt.Println(res8)
	 */

	 // ++ --
	 // 注意点: Go语言中的++ -- 只能写在后面, 不能写在前面
	 num := 1
	 fmt.Println(num) // 1
	 num++ // num = num + 1
	fmt.Println(num) // 2
	 num-- // num = num - 1
	 fmt.Println(num) // 1

	 //++num // 报错
	 //--num // 报错
}
