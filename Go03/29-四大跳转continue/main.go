package main

import "fmt"

func main() {
	/*
	1.在C语言中continue只能用于循环结构
	2.在Go语言中和C语言一样, 也只能用于循环结构

	3.使用的各种特性和C语言都一样, 唯一不一样的就是可以跳转到指定的标签

	注意点:
	如果跳转到指定的标签, 对于单层循环来说没有任何意义, 和不指定标签的效果一样
	如果跳转到指定的标签, 对于多层循环来说才有意义, 结束当前内循环, 跳到下一次外循环
	其它的注意点和break跳转到标签一样
	标签只能定义在前面, 标签和循环之间不能添加其它语句
	 */

	//for num := 1; num <= 10; num++ {
	//	if (num % 2) == 0{
	//		continue
	//	}
	//	fmt.Println(num)
	//}

//LNJ:
//	for num := 1; num <= 10; num++ {
//		if (num % 2) == 0{
//			continue LNJ
//		}
//		fmt.Println(num)
//	}

LNJ:
	for i:=1; i<=5;i++{
		fmt.Println(i)
		for j:=1; j <=10; j++{
			if (j % 2) == 0{
				continue LNJ
			}
			fmt.Println("---------", j)
		}
	}
}
