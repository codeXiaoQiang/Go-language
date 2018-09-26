package main

import "fmt"

func main() {
	/*
	1.变量的连续定义
	1.1在C语言中我们可以通过逗号来连续定义变量     int a, b, c;
	1.2在Go语言中我们也可以通过逗号来连续定义变量 var a, b, c int
	1.3在Go语言中除了可以利用逗号来连续定义变量以外, 还可以使用变量组的语法来同时定义多个变量
	var(
	   a int
	   b int
	   c int
	)

	2.同时定义多个变量, 也可以在定义的时候初始化
	2.1在C语言中我们可以通过   int a = 10, b = 20, c = 30;
	2.2在Go语言中如果通过逗号定义多个变量, 我们可以通过  var a, b, c = 10, 20, 30
	2.3在Go语言中如果通过变量组定义多个变量, 我们可以通过
	var(
	   a int = 10
	   b int = 20
	   c int = 30
	)

	注意点:
	在变量组中, 不能使用:=
	 */

	 // 1.通过逗号定义多个变量
	 /*
	 var a, b, c int
	 a = 123
	 b = 456
	 c = 789
	 fmt.Printf("%d,%d,%d\n", a, b, c)
	 */

	 // 2.通过变量组定义多个变量
	 /*
	 var(
	 	a int
	 	b int
	 	c int
	 )
	a = 123
	b = 456
	c = 789
	fmt.Printf("%d,%d,%d\n", a, b, c)
	 */

	 // 3.同时定义多个变量, 并且同时初始化
	 //var a, b, c int = 10, 20, 30
	 //var a, b, c = 10, 20, 30
	// a, b, c := 10, 20, 30
	//fmt.Printf("%d,%d,%d\n", a, b, c)


	// 4.同时定义多个变量, 并且同时初始化
	//var(
	//	a int = 10
	//	b int = 20
	//	c int = 30
	//)

	//var(
	//	a  = 10
	//	b  = 20
	//	c  = 30
	//)
	//fmt.Printf("%d,%d,%d\n", a, b, c)


	// 5.注意点:
	var(
		a  := 10
		b  := 20
		c  := 30
	)
	fmt.Printf("%d,%d,%d\n", a, b, c)
}
