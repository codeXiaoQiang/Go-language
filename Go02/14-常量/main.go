package main

import "fmt"

func main() {
	/*
	1.在C语言中可以通过const来定义常量
	2.在Go语言中一样
	3.格式:
	const 变量名称 数据类型 = 值;
	注意点:
	1. 数据类型可以省略, 但是const不能省略
	2.定义常量不能使用 := , := 是专门用于定义局部变量的
	3.定义局部变量没有使用, 编译会报错, 定义全局变量和常量没有使用, 不会报错

	4.在Go语言中可以连续定义多个变量, 所以也可以连续定义多个常量
	格式:
	const 变量名称1, 变量名称2 数据类型 = 值1, 值2;
	const(
	  变量名称1 数据类型 = 值1
	  变量名称2 数据类型 = 值2
	)

	5.常量组的注意点
    在常量组中, 如果常量没有赋值, 那么默认值就是上一行的取值
	 */

	 // 1.定义单个常量
	//const num int  =  666
	//const num  =  666
	//const num  :=  666
	//num = 789
	//fmt.Printf("%d\n", num)

	// 2.定义多个常量
	//const a, b, c int  = 10, 20, 30
	//const a, b, c  = 10, 20, 30
	//a = 666
	//b = 777
	//c = 888

	//const(
	//	a int = 10
	//	b int = 20
	//	c int = 30
	//)
	//const(
	//	a  = 10
	//	b  = 20
	//	c  = 30
	//)
	//a = 666
	//b = 777
	//c = 888


	// 3.常量组的注意点
	// 在常量组中, 如果常量没有赋值, 那么默认值就是上一行的取值
	//const(
	//	a  = 666
	//	b  = 789
	//	c
	//)
	//fmt.Printf("%d, %d, %d", a, b, c)


	//var(
	//	a, b = 10, 20
	//	c, d = 30, 40
	//)
	//fmt.Printf("%d, %d\n", a, b)
	//fmt.Printf("%d, %d\n", c, d)


	const(
		a, b = 10, 20
		c, d
	)
	fmt.Printf("%d, %d\n", a, b)
	fmt.Printf("%d, %d\n", c, d)
}
