package main

import "fmt"

/*
1.Go中的注释和C语言中的注释一样, 有多行注释和单行注释
所有的注意点, 细节都一样
// 从//开始到这一行的末尾
/星 星/ 从/星 到 星/为止
 */

	/*
	1.C语言是由于什么组成的?
	C语言是由函数组成的
	同样Go语言也是由函数组成的
	*/
	/*
	2.C语言程序的入口是谁?
	C语言程序的入口是main函数
	同样Go语言的程序入口也是main函数
	*/
	/*
	3.C语言main函数的注意点?
	一个程序只能有一个main函数, 没有main函数程序不能执行
	同样Go语言也只能有一个main函数, 没有main函数程序不能执行
	*/
	/*
	4.C语言main函数的格式?
	int main(){
	逻辑语句;
	}

	fun: 代表这是一个函数
	main: 函数的名称
	(): 形参列表
	{}: 函数体
	fun main(){
		逻辑语句;
	}
	*/
	/*
	5.C语言中如何往控制台输出内容?
	5.1#include <stdio.h> 告诉系统去哪找输出函数的实现
	5.2 printf("要输出的内容"); 告诉系统要输出什么内容

	在GO语言中如何往控制台输出内容?
	5.1 import "fmt" 告诉系统去哪找输出函数的实现
	5.2 fmt.Printf("要输出的内容"); 告诉系统要输出什么内容

	注意点:
	Go语言的main函数只能写在package main的这个包中
	*/
func main()  {
	fmt.Printf("Heldlo LddNJ1,nihao\n")
	fmt.Printf("Heldlo LddNJ2,nihao\n")
	fmt.Printf("Heldlo LddNJ3,nihao\n")
}