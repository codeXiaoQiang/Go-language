package main

import "fmt"

func main() {
	/*
	1.在C语言中, 我们使用printf来输出内容
	2.在Go语言中也可以使用printf输出内容
	但是在Go语言中还有其它更方便的函数, 也可以输出内容

	2.1第一种输出方式:
	fmt.Printf("格式化字符串", 数据列表)
	特点: 不会自动换行, 但是可以自定义输出格式
	2.2第二种输出方式
    fmt.Println(数据列表)
	特点: 会自动换行, 但是不能使用占位符%d%c%s
	 */

	 num, value := 10, 20
	 //fmt.Printf("num = %d, value = %d\n", num, value)
	 //fmt.Printf("------")

	 fmt.Println("num = ", num, "value = ",value)
	 fmt.Println("----")
}
