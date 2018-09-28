package main

import (
	"fmt"
)

var num = 666

func test() {
	fmt.Println("main test")
}

type person struct {
	name string
}

func (person) eat() {
	fmt.Println("main eat")
}

func main() {
	/*
	1.在Go语言中, 同一个包中的内容可以随意访问
	2.在Go语言中, 要想访问其它包中的内容, 那么内容名称的首字母必须大写

	3.要想访问全局变量, 导入对应的包之后, 可以通过 `包名.变量名` 的方式访问
	4.要想访问方式, 导入对应的包之后, 可以通过 `包名.方式名` 的方式访问
	5.
	 */
	/*
	// 访问当前包的全局变量
   fmt.Println(num)

   // 访问其它包的全局变量
   //fmt.Println(demo.value1)
   fmt.Println(demo.Value2)
	*/

	/*
   // 访问当前包的函数
   test()
   // 访问其它包的函数
   //demo.say1()
   demo.Say2()
	*/

	/*
	// 访问当前包的数据类型
	p := person{"zs"}
	fmt.Println(p)
   // 访问其它包的数据类型
   //stu1 := demo.student1{"ls"}
   //fmt.Println(stu1)
   stu2 := demo.Student2{"ls"}
   fmt.Println(stu2)
	*/

	/*
	// 访问当前包的方法
	p := person{"zs"}
	p.eat()
	// 访问其它包的方法
	stu2 := demo.Student2{"ls"}
	//stu2.eat1()
	stu2.Eat2()
	*/
}
