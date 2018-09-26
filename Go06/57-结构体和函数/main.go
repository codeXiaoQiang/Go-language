package main

import "fmt"

// 1.定义一个结构体类型
type Person struct {
	name string
	age int
}
func main() {
	// 结构体和函数: 结构体类型和数组类型一样, 在Go语言中都是值传递
	//          	 修改形参不会影响到实参

	// 2.定义一个结构体变量
	p1 := Person{"lnj", 18}

	//// 3.再定义一个结构体变量
	//var p2 Person
	//// 4.将p1赋值给p2
	//p2 = p1
	//// 5.修改p2的值
	//fmt.Println(p1)
	//p2.name = "zs"
	//fmt.Println(p1)

	fmt.Println(p1)
	change(p1)
	fmt.Println(p1)
}

func change(pp Person)  {
	pp.name = "zs"
}
