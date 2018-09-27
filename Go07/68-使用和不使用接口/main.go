package main

import "fmt"

// 0.定义一个接口
type USB interface {
	start()
	end()
}

// 1.电脑
type Computer struct {
	name string
}

func (cm Computer) start() {
	fmt.Println(cm.name, "被打开了")
}
func (cm Computer) end() {
	fmt.Println(cm.name, "被关闭了")
}

// 2.手机
type Phone struct {
	name string
}

func (ph Phone) start() {
	fmt.Println(ph.name, "被打开了")
}
func (ph Phone) end() {
	fmt.Println(ph.name, "被关闭了")
}

func main() {
	/*
	需求:  定义一个函数实现开关某种电器
	 */
	/*
	// 1.打开关闭电脑
	ctr := Computer{"惠普"}
	//ctr.start()
	//ctr.end()
   Option1(ctr)

	// 2.打开关闭手机
	pn := Phone{"华为"}
	//pn.start()
	//pn.end()
	Option2(pn)
	*/
	/*
	// 1.定义一个接口类型变量
	var in USB
	// 2.利用接口类型变量保存实现了接口的数据
	in = Computer{"惠普"}
	// 3.利用接口类型调用接口中声明的方法
	in.start()
	in.end()
	in = Phone{"华为"}
	in.start()
	in.end()
	*/
	ctr := Computer{"惠普"}
	Option(ctr)
	pn := Phone{"华为"}
	Option(pn)

	/*
	接口重点掌握的知识:
	1.如何定义一个接口
	type USB interface {
		函数名称(形参列表)(返回值列表)
		函数名称(形参列表)(返回值列表)
	}
	2.如何实现一个接口
	Go语言中实现一个接口, 不需要做额外的声明, 只要某种数据类型绑定了所有接口中什么的方法就是实现了这个接口

	3.实现接口之后的特性
	只要某种数据类型实现了接口, 那么就可以使用接口变量保存这种数据类型
	只要某种数据类型实现了接口, 那么就可以使用接口变量调用接口中声明的方法
	 */
}
func Option(u USB) {
	u.start()
	u.end()
}

func Option2(p Phone) {
	p.start()
	p.end()
}

func Option1(c Computer) {
	c.start()
	c.end()
}
