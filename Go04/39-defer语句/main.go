package main

import "fmt"

func main() {
	/*
	1.defer语句的格式
	defer 语句

	2.defer语句的作用
	延迟执行, 会在所属函数执行完毕之后再执行
	常用于资源释放, 解锁, 错误处理等等
	以前在C语言中我们说过, 通过malloc申请的存储空间, 一定更要释放free
	但是由于我们不能写完malloc之后立即free,所以经常导致我们忘记释放
	malloc -- >申请存储空间
	... ... --> 使用存储空间
	free --> 释放存储空间

	注意点:
	如果在同一个函数中编写了多个defer语句, 那么会遵守先进后出的原则
	先注册的defer语句后执行, 后注册的defer语句先执行
	 */

	 //fmt.Println("申请存储空间")
	 //defer fmt.Println("释放存储空间")
	 //fmt.Println("使用存储空间")
	 //fmt.Println("使用存储空间")
	 //fmt.Println("使用存储空间")


	 defer fmt.Println("第一条语句")
	 defer fmt.Println("第二条语句")
	 defer fmt.Println("第三条语句")
	 defer fmt.Println("第四条语句")
}
