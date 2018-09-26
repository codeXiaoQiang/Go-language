package main

import "fmt"

func main() {
	/*
	1.在C语言中函数只能返回一个值
	返回值类型 函数名称(形参列表){
	   逻辑语句;
	}

	2.在Go语言中函数可以返回多个值
	func 函数名称(形参列表) (返回值列表){
	  逻辑语句;
	}
	 */

	 // 1.返回一个值
	 //var num int = getValue()
	 //fmt.Printf("%d\n", num)

	 // 1.var num1 int
	 //   var num2 int
	 // 2. num1 = 10
	 //    num2 = 20
	 num1, num2 := calculate()
	 fmt.Printf("%d, %d\n", num1, num2)
}

func getValue()(int) {
	return  666
}

func calculate()(int, int){
	return 10, 20
}
