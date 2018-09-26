package main

import "fmt"

func main() {
	/*
	1.什么是闭包?
	闭包就是一个特殊的匿名函数
	只要匿名函数中用到了外界的变量, 那么我们就把匿名函数叫做闭包

	注意点:
	只要闭包还在使用外界的变量, 那么外界的变量就会一直存在
	 */

	 /*
	 num := 10 // 11 12
	 a := func() {
	 	num++
	 	fmt.Println(num)
	 }
	 a() // 11
	 a() // 12
	 a() // 13
	 */
	 // addUpper会返回一个闭包, 那么fn中保存的就是这个闭包, 那么fn就等于闭包
	 // 所以按照Go语言的规则, 只要fn还在, addUpper中被闭包引用的变量就不会释放
	 fn := addUpper()
	 fn()
	 fn()
	 fn()
}
func addUpper() func() {
	x := 1
	return func() {
		x++
		fmt.Println(x) // 2 , 3 , 4
	}
}
