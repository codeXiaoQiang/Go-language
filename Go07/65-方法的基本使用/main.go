package main

import "fmt"

// 1.定义一个结构体
type Person struct {
	name string
	age int
}
// 2.定义一个方法, 和Person结构体绑定在一起
// 已经将say这个函数和Person这个结构体绑定在一起了
//func (Person)say()  {
//	fmt.Println("Method say hello")
//}

// 2.定义一个方法, 和Person结构体绑定在一起
// 如果指定了接收者的名称, 那么调用方法时会将调用者传递给接收者
// 简单理解: 可以把接收者也看做是函数的一个形参
func (per Person)say()  {
	fmt.Println("name is ", per.name, "age is ", per.age)
}

// 3.既然可以把接收者看做是一个形参, 那么就存储值传递和地址传递的问题
func (per Person)setName(name string)  {
	per.name = name
}
func (per *Person)setAge(age int)  {
	(*per).age = age
}
func say()  {
	fmt.Println("function say hello")
}

func main() {
	/*
	1.什么是Go语言的方法?
	Go语言的方法就是一个特殊的函数, 函数是独立存在的, 而方法是和某种数据类型绑定在一起的

	2.如何将一个函数和某种数据类型绑定在一起?
	如下格式的含义就是将某个函数和接收者类型绑定在一起

	func (接收者名称 接收者类型)函数名称(形参列表)(返回值列表){
			逻辑语句;
	}

	3.Go语言中的函数可以和任何类型绑定, 但是一般用于和结构体绑定

	注意点:
	1.方法和函数的区别在于, 函数可以直接调用(通过包名.函数名称), 而方法只能通过绑定的数据类型对应的变量来调用(变量.函数名称)
	2.函数名称和方法名称可以重名
	 */
	 p := Person{"lnj", 18}
	 //p.setName("zs")
	 var ptr *Person = &p
	 ptr.setAge(33)
	 p.say()

	 say()
}
