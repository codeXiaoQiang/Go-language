package main

import "fmt"

type Person struct {
	name string
	age int
}
func (p *Person)say()  {
	fmt.Println(p.name, p.age)
}

// Go语言中的继承就是通过匿名属性来实现的
type Student struct {
	Person // 匿名属性
	score int
}
func (stu *Student)say()  {
	fmt.Println(stu.name, stu.age, stu.score)
}

func main()  {
	// 总结:
	// 子类可以使用父类的属性和方法
	// 如果子类和父类出现了重名的属性, 那么采用就近原则
	// 如果子类和父类出现了重名的属性, 要想访问父类的属性, 必须逐级查找
	// 如果子类和父类出现了重名的方法, 那么采用就近原则 (方法重载)
	// 如果子类和父类出现了重名的方法, 要想访问父类的方法, 必须逐级查找 (方法重载)
	stu := Student{}
	// 属性的继承
	stu.name = "zs"
	stu.age = 18
	stu.score = 66
	//fmt.Println(stu)
	// 方法的继承
	stu.say()
	stu.Person.say()
}