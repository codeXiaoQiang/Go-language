package main

import "fmt"

//type Integer int
type Integer = int

func main() {
	/*
	1.在C语言中可以通过typedef给某种类型起一个别名
	格式: typedef 原类型名 新类型名;

	2.在Go语言中可以通过type给某种类型起一个别名
	格式一: type 新类型名 原类型名
	格式二: type 新类型名=原类型名

	注意点:
	如果通过格式一, 代表定义了一个新的类型叫做 `新类型名`
	如果通过格式二, 代表给原类型名起了一个别名叫做`新类型名`
	也就是说通过格式一定义的 新类型 和 原类型在编译器看来是两个不同的类型
	       通过格式二定义的  新类型 和 原类型在编译器看来是同一个类型
	 */
	/*
	var num int = 666
	var value Integer = 888
	fmt.Println(num)
	fmt.Println(value)
	*/
	var num int = 666
	var value Integer
	value = num
	fmt.Println(value)
}
