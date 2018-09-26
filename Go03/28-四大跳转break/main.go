package main

import "fmt"

func main() {
	/*
	1.在C语言中有四大跳转语句 break continue goto return
	2.在Go语言中也有四大跳转语句  break continue goto return

	3.在C语言中break的应用范围
	switch语句和循环语句
	4.在Go语言中一样
	switch语句和循环语句
	但是Go语言的switch语句中默认就不会穿透, 所以break写不写都行, 一般都不写

	5.注意点:
	Go语言中的break可以跳转到指定的标签, 但是标签只能定义在前面, 并且跳转到标签之后循环不会被再次执行
	标签和循环之间不能编写其它语句
	 */

	//switch num := 2; num {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2")
	//	break
	//default:
	//	fmt.Println("OTHER")
	//}

	//for num := 1; num <= 10; num++ {
	//	fmt.Println(num)
	//	if num == 5{
	//		break
	//	}
	//}

lnj:
	//fmt.Println("start") // 标签和循环之间不能添加任何语句
	for num := 1; num <= 10; num++ {
		fmt.Println(num)
		if num == 5{
			break lnj
		}
	}
//lnj: // 标签只能写在前面
	fmt.Println("end")
}
