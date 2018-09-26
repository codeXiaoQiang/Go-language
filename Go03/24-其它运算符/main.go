package main

import (
	"fmt"
)

func main() {
	/*
	1.C语言中的其它运算符
	& 取地址运算符
	* 访问地址运算符

	2.Go语言和C语言一样
	& 取地址运算符
	* 访问地址运算符
	 */

	 var num int
	 num = 10
	 // 定义一个指向int类型变量的指针
	 var p *int
	 // 取出num变量的内存地址, 放到指针变量p中
	 p = &num
	 // 访问指针指向的内存
	 fmt.Println(*p)
}
