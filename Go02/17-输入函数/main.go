package main

import "fmt"

func main() {
	/*
	1.在C语言中, 我们使用scanf来接收输入的内容
	2.Go语言中也可以使用scanf来接收输入的内容
	但是在Go语言中还有其它更方便的函数, 也可以接收输入的内容


	2.1第二种方式
	fmt.Scan(地址列表)
	特点: 如果接收的不是字符串类型(%c), 会忽略空格和TAB和回车, 相当于C语言的scanf

	2.2第一种方式:
	fmt.Scanf(格式化字符串, 地址列表)
	特点: 如果接收的不是字符串类型(%c), 会忽略空格和TAB, 但是不会忽略回车

	2.3第三种方式
	fmt.Scanln(地址列表)
	特点: 如果接收的不是字符串类型(%c), 会忽略空格和TAB, 但是不会忽略回车
	 */

	 var num int;
	 var value int;
     //fmt.Scanf("%d", &num)
     //fmt.Scan(&num, &value)
     fmt.Scanln(&num, &value)
     fmt.Println(num, value)
}
