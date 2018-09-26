package main

import "fmt"

func main() {
	/*
	1.什么是字典?
	和数组以及切片一样, 字典是用来保存一组相同类型的数据的
	数组和切片可以通过索引获得对应元素的值
	字典可以通过key获得对应元素的值

	数组和切片的索引是系统自动从0开始递增生成的
	字典的key必须我们自己指定

	2.如何定义一个字典
	格式: var 字典名称 map[key的数据类型]value的数据类型


	3.如何给字典初始化
	3.1通过语法糖来创建
	3.2通过make(字典数据类型, 长度)
	3.3.通过make(字典数据类型)

	4.如何使用字典
	和数组很像, 通过 字典变量名称[key名称]的方式来操作字典

	5.注意点:
	和切片一样, 没有初始化的字典是不能使用的
	只要能够进行== !=判断的数据类型, 都可以作为字典的key
	 */

	 // 1.创建字典的第一种方式
	 /*
	 // 1.定义一个字典变量
	 var dict map[int]int
	 // 2.通过语法糖来创建一个字典(初始化)
	 dict = map[int]int{0:1, 1:3, 2:5}
	 // 3.使用创建好的字典变量
	fmt.Println(dict)
	dict[0] = 666
	fmt.Println(dict)
	 */

	 // 2.创建字典的第二种方式
	 /*
	// 1.定义一个字典变量
	var dict map[int]int
	// 2.通过make函数创建一个字典(初始化)
	 // 在企业开发中, 如果事先知道字典中需要保存多少条数据, 那么建议指定字典的长度
	dict = make(map[int]int, 3)
	// 3.使用创建好的字典变量
	fmt.Println(dict)
	dict[0] = 123
	dict[1] = 456
	dict[2] = 789
	dict[3] = 999
	fmt.Println(dict)
	 */
	 // 3.创建字典的第三种方式
	 /*
	// 1.定义一个字典变量
	var dict map[int]int
	// 2.通过make函数创建一个字典(初始化)
	dict = make(map[int]int)
	// 3.使用创建好的字典变量
	fmt.Println(dict)
	dict[0] = 123
	dict[1] = 456
	dict[2] = 789
	dict[3] = 999
	fmt.Println(dict)
	 */

	 /*
	 var dict map[string]string = map[string]string{"name":"lnj", "age":"18"}
	 fmt.Println(dict["name"])
	 fmt.Println(dict["age"])
	 */

	 // 遍历字典
	 dict := map[string]string{"name":"lnj", "age":"18"}
	 for key, value := range dict{
	 	fmt.Println("key = ", key, "value = ", value)
	 }
}
