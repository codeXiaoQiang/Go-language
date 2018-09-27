package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/*
	 1.什么是空接口?
	Go语言中的空接口, 相当于其他语言的Object类型
	Go语言中的空接口, 可以充当任何类型

	2.空接口类型的格式
	interface{}

	3.正式因为有了空接口类型, 所以也可以让数组和字典保存不同类型的数据
	面试回答: 数组和字典一般用于保存相同类型的数据, 而结构体用于保存不同类型的数据
	          但是也可以通过空接口类型来实现让数组和字典保存不同类型的数据

	注意点: 空接口类型和普通的数据类型, 在使用的时候还是有很大区别的
	 */
	 /*
	// 1.定义一个空接口类型变量
	var value interface{}
	value = 1 // 保存整型
	fmt.Println(value)
	value = 3.14 // 保存浮点类型
	fmt.Println(value)
	value = false // 保存布尔类型
	fmt.Println(value)
	value = 'T' // 保存字符
	fmt.Println(value)
	value = "lnj" // 保存字符串类型
	fmt.Println(value)
	value = [3]int{1, 3, 5} // 保存数组类型
	fmt.Println(value)
	value = []int{2, 4, 5} // 保存切片类型
	fmt.Println(value)
	value = map[string]string{"name": "lnj", "age": "18"} // 保存字典类型
	fmt.Println(value)
	value = Person{"lnj", 33} // 保存结构体类型
	fmt.Println(value)
	 */


	 // 1.定义一个数组
	 var arr [3]interface{}
	 // 2.往数组中存储数据
	 arr[0] = 1
	 arr[1] = "lnj"
	 arr[2] = false
	 // 3.打印数据
	 fmt.Println(arr)
}
