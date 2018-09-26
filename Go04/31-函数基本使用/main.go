package main

import "fmt"

func main() {
	/*
	1.C语言是由函数组成的
	2.Go语言也是由函数组成的
	3.C语言函数的格式
	返回值类型 函数名称(形参列表){
		被控制的语句;
	}
	4.Go语言函数的格式
	func 函数名称(形参列表)(返回值列表){
		被控制的语句;
	}

	5.C语言函数的类型
	5.1没有返回值没有形参的函数
	5.2有返回值没有形参的函数
	5.3没有返回值有形参的函数
	5.4有返回值有形参的函数

	6.Go语言函数的类型和C语言一样

	注意点:
	1.在C语言中如果函数的定义写在函数的调用之前必须先声明再调用
	在Go语言中如果函数的定义写在函数的调用之前, 不用声明就可以调用, Go语言没有声明
	2.Go语言中如果函数只有一个返回值, 那么返回值列表的()可以省略
	3.Go语言中如果有返回值, 可以在返回值列表中定义变量, 然后在函数体中直接使用
	4.如果在返回值列表中定义了变量,  那么return后面可以不用写任何内容, 默认就会将返回值列表中定义的变量返回
	5.Go语言中一个函数可以有多个返回值
	6.如果一个函数有多个返回值, 但是我们只想使用其中的一个, 那么可以在接收的时候使用_来忽略, 不需要的那个返回值
	7.因为_在Go语言中有特殊的含义, 所以在定义标识符的时候_不能单独作为标识符
	8.如果返回值列表中有多个返回值, 并且连续的多个返回值的数据类型都一样, 那么只用在最后一个变量后面添加数据类型即可
	9.如果形参列表中有多个形参, 并且连续的多个形参的数据类型都一样, 那么只用在最后一个变量后面添加数据类型即可
	 */
	//say()
	//res := getNumber()
	//fmt.Println(res)

	//a, b := getValue()
	//fmt.Println("a = ", a, "b = ", b)
	//a, _ := getValue()
	//fmt.Println("a = ", a)

	//sum(10, 20)

	res := minus(10, 20)
	fmt.Println(res)
}
// 1.没有返回值没有形参的函数
func say(){
	fmt.Println("没有返回值没有形参")
}
// 2.有返回值没有形参的函数
// 2.1只有一个返回值的函数
//func getNumber()(int){
//func getNumber()int{
//	return 888
//}
func getNumber() (num int){
	num = 999
	return
}
// 2.2有多个返回值的函数
//func getValue() (int, int){
//	return 666, 888
//}
//func getValue() (num int,value int){
func getValue() (num ,value int){
	//num = 666
	//value = 888
	//return
	return 666, 888
}
//func test()(num int, temp float64, value int){
//func test()(num , value int, temp float64){
//
//}

// 3.没有返回值有形参的函数
//func sum(a int, b int){
func sum(a , b int){
	res := a + b
	fmt.Println(res)
}

// 4.有返回值有形参的函数
//func minus(a, b int)(res int){
//	res = a - b
//	return
//}

func minus(a int, b int)(int){
	res := a - b
	return res
}
