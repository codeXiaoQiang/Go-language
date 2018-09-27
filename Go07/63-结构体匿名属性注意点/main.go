package main

import "fmt"

func main() {
	/*
	1.结构体的嵌套定义的注意点:
	1.1结构体的属性类型不能是当前结构体的类型
	1.2只有匿名属性才会向上查找, 非匿名属性不会向上查找
	 */
	//type Person struct {
	//	Person // 会报错, 不能自己搞自己
	//	name string
	//}

	//type Person struct {
	//	*Person // 不会报错, 链表
	//	name string
	//}

	type Person struct {
		name string
	}
	type Student struct {
		//Person // 匿名属性
		per Person // 匿名属性
		age int
	}

	stu := Student{Person{"lnj"}, 18}
	//fmt.Println(stu.per.name)
	fmt.Println(stu.name)
}
