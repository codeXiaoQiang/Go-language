package main

import "fmt"

// 1.定义一个接口
type Animal interface {
	eat()
}
type Dog struct {
	name string
}

func (d Dog)eat()  {
	fmt.Println(d.name, "吃东西")
}

// 3.定义一个猫的结构体
type Cat struct {
	name string
}

func (c Cat)eat()  {
	fmt.Println(c.name, "吃东西")
}

func main()  {
	/*
	1.什么是多态?
	多态就是某一类事物的多种形态

	2.在Go语言中通过接口来实现多态
	 */
	//// 1.666的int形态
	//var num int = 666
	//// 2.66的接口形态
	//var value interface{} = 666


	//// 1.旺财狗的形态
	//var d Dog = Dog{"旺财"}
	//d.eat()
	//// 2.旺财动物的形态
	//var a1 Animal = Dog{"旺财"}
	//a1.eat()
	//
	//// 1.外星人猫的形态
	//var c Cat = Cat{"外星人"}
	//c.eat()
	//// 2.外星人动物的形态
	//var a2 Cat = Cat{"外星人"}
	//a2.eat()

	var d Dog = Dog{"旺财"}
	var c Cat = Cat{"外星人"}
	eat(d)
	eat(c)
}
func eat(a Animal)  {
	a.eat()
}