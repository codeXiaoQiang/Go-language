package main

import "fmt"

func main() {
	/*
	1.在C语言中如果基本数据类型作为函数的参数, 是值传递, 在函数内修改形参不会影响到外界的实参
	2.Go语言和C语言一样/ int float bool

	3.如果想在函数内部修改外界传入的实参, 那么必须传入指针
	 */
	num := 666
	fmt.Println(num)
	//change1(num)
	change2(&num)
	fmt.Println(num)
}
func change1(a int) {
	a = 789
}

func change2(a *int) {
	*a = 789
}
