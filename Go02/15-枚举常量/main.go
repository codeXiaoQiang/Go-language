package main

import "fmt"

func main() {
	/*
	1.C语言中的枚举本质就是一个整数
	enum　枚举名　{
		枚举元素1 = 值,
		枚举元素2 = 值,
		……
	};

	2.Go语句中没有明确枚举的固定写法, 但是在企业开发中一般都会常量组的形式来表示枚举
	const(
	 枚举元素 = 值
	  枚举元素 = 值
	)
	 */

	//const(
	//	male = 0
	//	female = 1
	//)
	// 注意点:
	// 1.iota迭代器, 默认会从0开始递增

	//const(
	//	male = iota
	//	female = iota
	//	yao = iota
	//)
	// 2.只要常量组中出现了iota, 该常量组中后续的常量都会一次递增1
	//const(
	//	male = iota
	//	female
	//	yao
	//)
	// 3.如果常量组中的iota被打断了, 那么就不会继续递增了, 会按照常量组的默认方式处理(上一行的值)
	//const(
	//	male = iota
	//	female = 666
	//	yao
	//)

	// 4.如果常量组中的itoa被打断了, 但是后续又被回复了, 那么前面有多少行就会递增多少
	const(
		male = iota
		female = 666
		yao = iota
	)
	fmt.Printf("%d\n", male)
	fmt.Printf("%d\n", female)
	fmt.Printf("%d\n", yao)

	//var num int = male
	//fmt.Printf("%d\n", num)
}
