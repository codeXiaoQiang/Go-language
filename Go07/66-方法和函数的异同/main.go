package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 1.定义一个方法
func (Person) say() {
	fmt.Println("Method say hello")
}

// 3.既然可以把接收者看做是一个形参, 那么就存储值传递和地址传递的问题
// 值传递
func (per Person) setName(name string) {
	per.name = name
}

// 地址传递
func (per *Person) setAge(age int) {
	(*per).age = age
}

func test() {
	fmt.Println("test")
}
func main() {
	/*
	1.方法和函数的异同
	1.1方法和函数都是 函数类型
	1.2函数通过 包名.函数名称() 调用 / 方法通过  变量.方法名称() 调用
	1.3方法的接手者可以看做是函数的一个形参, 会将调用者传递给接收者
	   接收者和函数形参一样, 也区分值传递和地址传递
	   到底是值传递还是地址传递, 和函数的形参一样
	注意点:
	接收者地址传递的两种方式
    (*指针变量名).方法名称()
	普通变量名.方法名称()
	只要接收者是指针类型, 那么编译器就会自动将普通变量的地址取出来传递给接收者
	 */
	// 1.方法和函数都是 函数类型
	/*
	p := Person{"lnj", 33}
	//fmt.Printf("say = %T\n", p.say) // func()
	//fmt.Printf("test = %T\n", test) // func()

	// 既然方法和函数都是函数类型, 所以可以定义变量保存函数, 也可以保存方法

	// 定义一个用于保存没有形参没有返回值的函数类型变量
	var fn func()
	fn = test
	fn()

	fn = p.say
	fn()
	*/

	// 3.值传递和地址传递
	per := Person{"lnj", 33}
	//per.setName("zs")

	//p := &per
	//(*p).setAge(666)
	per.setAge(666) // 底层相当于(&per).setAge(666)
	fmt.Println(per)
}
