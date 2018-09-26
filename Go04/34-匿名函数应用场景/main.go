package main

import "fmt"

func main() {
	/*
	1.作为其他函数的参数
	2.作为其他函数的返回值
	 */
	 /*
	 res1 := sum(10, 20)
	 res2 := minus(10, 20)
	 fmt.Println(res1)
	 fmt.Println(res2)
	 */

	 // 需求: 定义一个匿名函数, 实现两个数相加
	//fn := func(a, b int) int {
	//	return  a + b
	//}
	//fmt.Printf("%T\n", fn)

	//var fn func(int, int) int = func(a, b int) int {
	//	return  a + b
	//}
	//res := fn(10, 20)
	//fmt.Println(res)

	//var method func(int, int) int = fn
	//res := method(10, 20)
	//fmt.Println(res)

	//var num int = 6
	//var value int = num
	//fmt.Println(value)

	//result := calculate(10, 20, fn)

	result := calculate(10, 20, func(a, b int) int {
		return  a + b
	})
	fmt.Println(result)
}
func calculate(a, b int, method func(int, int) int) (int) {
	res := method(a, b)
	return res
}

func sum(a, b int) int {
	return  a + b
}
func minus(a, b int) int {
	return  a - b
}
