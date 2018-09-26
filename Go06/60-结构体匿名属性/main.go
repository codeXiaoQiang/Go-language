package main

import "fmt"

func main() {
	/*
	1.什么是匿名结构体属性?
	没有名称的属性就是匿名属性

	2.如何操作匿名属性?
	匿名属性的数据类型就是匿名属性的名称, 所以可以通过匿名属性的数据类型来操作

	3.匿名属性一般都用于结构体的嵌套定义
	结构体的属性又是一个结构体
	 */

	 /*
	type Person struct {
		int // 只有数据类型, 没有名称, 就是匿名属性
		name string
	}

	//var per Person = Person{666, "lnj"}
	//fmt.Println(per)

	var per Person
	per.int = 666
	per.name = "lnj"
	fmt.Println(per)
	 */

	type Date struct {
		year int
		month int
		day int
	}
	type Person struct {
		name string
		//year int
		//month int
		//day int
		Date
	}

	type Animal struct {
		name string
		//year int
		//month int
		//day int
		Date
	}

	//var per Person
	//per.name = "zs"
	//per.Date = Date{2012, 12,12}
	//fmt.Println(per) // {zs {2012 12 12}}

	var per Person = Person{"zs", Date{2012, 12, 12}}
	fmt.Println(per) // {zs {2012 12 12}}
}
