package main

import "fmt"

func main() {
	/*
	第一种方式
	type 结构体名称1 struct{

	}
	type 结构体名称2 struct{
		结构体名称1
	}

	第二种方式
	type 结构体名称1 struct{

	}
	type 结构体名称2 struct{
		结构体名称1
	}
	type 结构体名称3 struct{
		结构体名称2
	}

	第三种方式
	type 结构体名称1 struct{

	}
	type 结构体名称2 struct{
	}
	type 结构体名称3 struct{
	    结构体名称1
		结构体名称2
	}
	 */

	 // 第一种方式: 没有重名属性的情况
	 /*
	type Person struct {
		name string
		age int
	}
	type Student struct {
		Person
		score float32
	}

	stu := Student{Person{"zs", 18}, 99.5}
	fmt.Println(stu)

	// 结构体嵌套定义时, 如果操作结构体的属性
	//fmt.Println(stu.Person.name)
	//fmt.Println(stu.Person.age)
	//fmt.Println(stu.score)

	// 第二种访问的方式
	// 会先去Student结构体中查询有没有name属性, 如果有就直接访问
	// 如果没有会继续查找嵌套的匿名结构体中有没有, 如果有就访问匿名结构体中的name属性
	//fmt.Println(stu.name)
	//fmt.Println(stu.age)
	//fmt.Println(stu.score)
	*/

	// 第一种方式: 有重名属性的情况
	/*
	type Person struct {
		name string
		age int
	}
	type Teacher struct {
		Person
		name string
		title string
	}

	tea := Teacher{Person{"lnj", 33}, "zs", "老师"}
	fmt.Println(tea)
	fmt.Println(tea.name) // zs
	fmt.Println(tea.Person.name) // lnj
	*/


	// 第二种方式: 没有重名属性的情况
	/*
	type Object struct {
		name string
	}
	type Person struct {
		Object
		age int
	}
	type Student struct {
		Person
		score float32
	}

	stu := Student{Person{Object{"zs"}, 18}, 99.5}
	fmt.Println(stu) // {{{zs} 18} 99.5}
	fmt.Println(stu.score)
	fmt.Println(stu.Person.age)
	fmt.Println(stu.age)
	fmt.Println(stu.Person.Object.name)
	fmt.Println(stu.Person.name)
	fmt.Println(stu.name)
	*/

	// 第二种方式: 有重名属性的情况
	/*
	type Object struct {
		name string
	}
	type Person struct {
		Object
		name string
		age int
	}
	type Student struct {
		Person
		name string
		score float32
	}
	stu := Student{Person{Object{"lnj"}, "zs", 18}, "ww", 99.5}
	fmt.Println(stu)
	fmt.Println(stu.name) // ww
	fmt.Println(stu.Person.name) // zs
	fmt.Println(stu.Person.Object.name) // lnj
	*/

	// 第三种方式: 没有重名属性的情况
	/*
	type Object struct {
		name string
	}
	type Person struct {
		age int
	}
	type Student struct {
		Object
		Person
		score float32
	}
	stu := Student{Object{"lnj"}, Person{18}, 99.5}
	fmt.Println(stu)
	fmt.Println(stu.Object.name)
	fmt.Println(stu.name)
	fmt.Println(stu.Person.age)
	fmt.Println(stu.age)
	*/

	type Object struct {
		name string
	}
	type Person struct {
		name string
	}
	type Student struct {
		Object
		Person
		name string
		score float32
	}

	stu := Student{Object{"zs"}, Person{"ls"}, "lnj",99.5}
	fmt.Println(stu.name)
	fmt.Println(stu.Object.name)
	fmt.Println(stu.Person.name)
}
