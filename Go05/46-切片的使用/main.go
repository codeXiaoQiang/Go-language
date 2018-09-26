package main

import "fmt"

func main() {
	/*
	1.可以先使用普通数组那样使用切片

	2.可以通过预先定义好的函数来使用切片
	append(切片, 数据)

	3.append函数的作用:
	在len后面追加数据

	4.容量
	如果通过append函数追加数据之后超过了原有的容量, 那么系统内部会自动按照当前容量*2的方式重新定义一个数组作为切片保存数据的模型

	5.为什么append函数要返回一个切片
	因为在扩容的时候会重新定义一个新的切片, 所以需要返回一个切片
	 */
	 /*
	var sce []int = make([]int, 2, 5)
	fmt.Println(sce)
	//sce[0] = 1
	//sce[1] = 2
	// 注意点: 如果像使用数组一样使用切片, 那么索引的长度不能超过切片的len
	//sce[2] = 999
	sce = append(sce, 3)
	//fmt.Println(sce)
	fmt.Printf("%p\n", sce)
	sce = append(sce, 4)
	fmt.Printf("%p\n", sce)
	sce = append(sce, 5)
	fmt.Printf("%p\n", sce)
	fmt.Println(sce)
	fmt.Println(cap(sce))

	// 超过容量
	sce = append(sce, 6)
	fmt.Printf("%p\n", sce)
	fmt.Println(sce)
	fmt.Println(cap(sce))
	 */

	 // 注意点: 容量小于1024的时候扩容, 会按照原有容量的2倍扩容
	 //         容量大于1024的时候扩容, 会按照原有容量1/4倍扩容
	var sce []int = make([]int, 1024, 1024)
	fmt.Printf("%p\n", sce)
	fmt.Println(cap(sce))
	sce = append(sce, 1025)
	fmt.Printf("%p\n", sce)
	fmt.Println(cap(sce))
}
