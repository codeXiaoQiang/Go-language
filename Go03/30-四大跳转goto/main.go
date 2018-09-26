package main

import "fmt"

func main() {
	/*
	1.C语言中的goto可以在同一个函数中随便乱跳
	2.Go语言中的goto也可以在同一个函数中随便乱跳
	所以就是一模一样

	注意点:
	可以往前跳, 也可以往后跳
	标签和跳转语句之间可以编写其它的代码
	 */

//	num := 1
//outer:
//	fmt.Println("start")
//	if(num <= 10){
//		fmt.Println(num)
//		num++
//		goto outer // 死循环
//	}
//	fmt.Println("come here")

	for i:=1; i <= 10; i++{
		fmt.Println(i)
		if i == 5{
			goto LNJ
		}
	}
LNJ:
	fmt.Println("come here")
}
