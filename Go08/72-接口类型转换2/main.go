package main

import "fmt"

// 1.定义接口
type USB interface {
	start()
}

// 2.定义一个结构体
type Computer struct {
	name string
}

// 利用结构体实现接口中所有的方法
func (cm Computer) start() {
	fmt.Println(cm.name, "启动了")
}

// 除了实现了接口中声明的方法以外, 结构体中还有自己特有的方法
func (cm Computer) say() {
	fmt.Println(cm.name)
}
func main() {
	/*
	1.如果结构体实现了某个接口, 那么就可以使用接口类型来保存结构体变量
	2.如果利用接口类型变量保存了实现了接口的结构体, 那么该变量只能访问接口中的方法, 不能访问结构体中的特有方法, 以及结构体中的属性
	3.如果利用接口类型变量保存了实现了接口的结构体,想要访问结构体中特有的方法和属性, 那么必须进行类型转换, 将接口类型转换为结构体类型
	 */

	 /*
	// 1.定义一个接口变量
	var in USB
	// 2.利用接口变量保存结构体变量
	in = Computer{"惠普"}
	// 3.可以利用接口变量调用接口中什么的方法
	in.start()
	// 4.不可以利用接口变量调用结构体中特有的方法
	//in.say()
	// 5.不可以利用结构变量访问结构体中的属性
	//fmt.Println(in.name)
	// 6.如果想调用结构体中特有的方法和访问结构体中特有的属性,
	// 必须将接口类型转换为结构体类型
	// 7.方式一: 通过 value, ok := 接口变量名.(具体数据类型)
	//if value, ok := in.(Computer); ok{
	//	value.say()
	//	fmt.Println(value.name)
	//}

	// 8.方式二: 通过switch value := 接口变量名.(type) { case 具体数据类型: }
	//switch value := in.(type) {
	//case Computer:
	//	value.say()
	//	fmt.Println(value.name)
	//}
	 */

	// 9.方式一除了可以将接口类型转换为具体类型以外, 还可以将抽象接口类型转换为具体的接口类型
	// 9.1定义一个抽象的接口类型
	var in interface{}
	// 9.2利用抽象的接口类型保存结构体
	in = Computer{"惠普"}
	// 9.3不可以利用抽象接口类型变量, 调用具体接口类型中声明的方法
	//in.start()
	// 9.4要么将抽象接口类型转换为具体的接口类型
	if value, ok := in.(USB); ok{
		value.start()
	}
	// 9.6要么将抽象接口类型转换为结构体类型
	if value, ok := in.(Computer); ok{
		value.start()
	}
}
