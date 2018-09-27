package main

import "fmt"

// 1.定义一个接口
type USB interface {
	// 制定一个方法的标准
	// 名称必须叫做start, 必须传入一个string类型参数, 必须没有返回值
	start(name string)
	// 名称必须叫做end, 必须传入一个string类型参数, 必须没有返回值
	end(name string)
}

type Computer struct {
}
// 实现接口中声明的方法
func (Computer)start(name string)  {
fmt.Println(name, "启动了")
}
func (Computer)end(name string)  {
	fmt.Println(name, "关闭了")
}

type Phone struct {

}
// 实现接口中声明的方法
func (Phone)start(name string)  {
	fmt.Println(name, "启动了")
}
func (Phone)end(name string)  {
	fmt.Println(name, "关闭了")
}

func main() {
	/*
	1.什么是接口?
	Go语言中的接口和现实生活中的USB3.0接口很像, 用于定义某种规范
	现实生活中, 只要生产厂家按照USB3.0接口的标准来生产, 那么无论什么品牌的电脑, 都可以使用
	Go语言中, 只要使用者按照我们接口定义的标准来实现, 那么无论谁实现的我们都可以使用

	2.接口的作用?
	在Go语言中, 接口专门用于定义函数的声明
	也就是规定函数的形参, 函数的返回值, 函数的名称

	3.定义接口的格式:
	type 接口的名称 interface{
		函数的声明
	}

	4.注意点:
	4.1接口是虚的, 只有方法的声明, 没有方法的实现 (就和制定USB3.0标准的组织一样, 只负责制定标准, 不负责生产)
	4.2接口中声明的方法, 只能通过和某种数据类型绑定的方法来实现, 不能通过函数来实现 (就和现实生活中只能由富士康, 等具备生产能力的公司来生成一样, 这里和和某种数据类型绑定的方法就是局部生产能力的公司, 函数就是不具备生产能力的公司)
	4.3在Go语言中, 只要某个数据类型实现了接口中声明的`所有`方法, 那么就说这个数据类型实现了这个接口
	4.4只要一个数据类型实现了某个接口, 那么就可以使用这个接口类型的变量来保存这个类型的数据
	4.5只要一个数据类型实现了某个接口, 那么保存这个类型的数据之后, 就可以使用接口类型变量调用接口中的方法
	 */

	 /*
	 // 0.定义一个Computer类型变量
	 var cm Computer
	 cm = Computer{}

	 // 1.定义一个接口类型的变量
	 var in USB
	 // 2.利用接口类型的变量保存实现了该接口的数据
	 in = cm
	 in.start("惠普")
	 in.end("惠普")
	 */

	 /*
	// 1.定义一个接口类型的变量
	var in USB
	// 2.利用接口类型的变量保存实现了该接口的数据
	in = Computer{}
	// 3.利用接口类型调用接口中的方法
	in.start("惠普")
	in.end("惠普")
	 */

	 /*
	// 1.定义一个Computer类型变量
	var cm Computer
	cm = Computer{}
	// 2.利用Computer类型变量调用Computer中的方法
	cm.start("惠普")
	cm.end("惠普")
	 */

	 /*
	// 1.定义一个接口类型的变量
	var in USB
	// 2.利用接口类型的变量保存实现了该接口的数据
	in = Phone{}
	in.start("华为")
	in.end("华为")
	 */


	 cm := Computer{}
	 Option(cm, "惠普")
	 pp := Phone{}
	 Option(pp, "华为")

}

//func Option1(in Computer, name string)  {
//	in.start(name)
//	in.end(name)
//}
//func Option2(in Phone, name string)  {
//	in.start(name)
//	in.end(name)
//}

// 该函数的作用, 就是开关某种电器
func Option(in USB, name string)  {
	in.start(name)
	in.end(name)
}

