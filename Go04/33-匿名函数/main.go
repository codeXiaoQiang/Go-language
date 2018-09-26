package main

import "fmt"

func main() {
	/*
	1.什么是匿名函数?
	普通的函数都是有名字的函数, 而匿名函数就是没有名字的函数

	2.匿名函数的应用场景:
	一般情况下某个函数只会被调用一次, 我们就会使用匿名函数
	作为函数的参数, 或者作为函数的返回值

    3.注意点:
	3.1匿名函数定义之后, 如果没有立即调用, 编译器会报错
	3.2也可以将一个匿名函数保存到一个变量中
	在C语言中我们知道, 我们可以定义一个指向函数的指针来保存函数的地址, 然后通过指针来调用函数
	在Go语言中函数是一等公民(可以定义变量, 作为形参, 作为返回值), 所以在Go语言中也可以用函数类型来定义变量(相当于C语言中指向函数的指针)
	3.3由于函数是一等公民, 所以以前定义变量所有的格式, 都可以用于定义函数类型的变量
	var 变量名称 函数类型
	var 变量名称 函数类型 = 匿名函数
	var 变量名称  = 匿名函数
	变量名称  := 匿名函数
	var(
		var 变量名称  = 匿名函数
	)

	C语言中指向函数的指针
	int test(int a, int b){}
	int (*test)(int, int)
	 */
	 // test相当于找到函数的定义, 然后调用函数
	 test()

	 // 直接定义一个函数, 然后立即调用这个函数
	 //func(){
	 //	fmt.Println("匿名函数")
	 //}()

	 // 定义一个函数类型的变量, 保存匿名函数, 然后通过函数类型的变量调用匿名函数
	 //f := func() {
		// fmt.Println("匿名函数")
	 //}
	 var f func() = func() {
		 fmt.Println("匿名函数")
	 }
	 //fmt.Printf("%T\n", f)
	 f()
}

func test(){
	fmt.Println("test")
}
