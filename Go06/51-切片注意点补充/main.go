package main

import "fmt"

func main() {
	/*
	1.如何创建切片?
	1.1通过一个数组来创建切片
	arr [3]int = [3]int{1, 3, 5}
	sce []int = arr[:]
	1.2通过make函数来创建切片
	sce := make([]int, 0, 3)
	1.3通过go语言的语法糖来创建切片
	sce := []int{1, 3, 5}

	2.通过数组创建切片补充
	格式: 数组名称[low: high, max]

	3.仅仅定义, 没有初识的数组是可以使用的, 但是仅仅定义但是没有初始化的切片是不能使用的
	 */

	/*
	// 					  0  1  2  3  4
	var arr [5]int = [5]int{1, 3, 5, 7, 9}
	//var sce []int = arr[2:4]
	//fmt.Println(sce) // [5 7]
	//fmt.Println(len(sce)) // 2
	//fmt.Println(cap(sce)) // 3  数组以前的长度 - 截取的开始位置 = 容量

	// 切片的时第三个参数max,必须大于或等于第二个参数
	// 如果指定了第三个参数, 那么切片的容量就等于 第三个参数 - 第一个参数
	// 总结: 如果没有指定第三个参数, 那么切片的容量就等于 数组的容量 - 第一个参数
	//       如果指定了第三个参数, 那么切片的容量就等于 第三个参数 - 第一个参数
	var sce []int = arr[2:4:5]
	fmt.Println(sce)
	fmt.Println(len(sce))
	fmt.Println(cap(sce))
	*/

	//// 1.定义一个数组变量
	//var arr [3]int
	//// 2.使用数组变量
	//arr[0] = 1
	//arr[1] = 3
	//arr[2] = 5
	//// 3.打印使用之后的结果
	//fmt.Println(arr)

	// 1.定义一个切片变量
	var sce []int
	// 对切片进行初始
	//sce = make([]int, 3, 5)

	// 2.使用切片变量
	//sce[0] = 1
	//sce[1] = 3
	//sce[2] = 5

	//sce = append(sce, 1)
	//sce = append(sce, 3)
	//sce = append(sce, 5)

	// 3.打印使用之后的结果
	fmt.Println(sce)

	/*
	var arr [3]int = [3]int{1, 3, 5}
	// 第一种方式: 推荐方式
	//length := len(arr)
	//fmt.Println(length)
	// 第二种方式: 不推荐方式
	//size := unsafe.Sizeof(arr)
	//fmt.Println(size)
	//eleSize := unsafe.Sizeof(arr[0])
	//fmt.Println(eleSize)
	//fmt.Println(size/eleSize)
	*/
	/*
	//var sce []int = []int{1, 3, 5, 7, 9}
	// 方式一: 推荐方式
	//length := len(sce)
	//fmt.Println(length)

	// 方式二: 注意点, 切片不能通过sizeof来计算
	// 由于切片是一个结构体内部指向了一个数组, 所以通过sizeof计算的是结构体的大小, 而不是切片的长度
	//size := unsafe.Sizeof(sce)
	//fmt.Println(size)
	*/
}
