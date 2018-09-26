package main

import "fmt"

func main() {
	/*
	1.和C语言一样,Go语言中也有数组的概念, Go语言中的数组也是用于保存一组相同类型的数据
	2.和C语言一样,Go语言中的数组也分为一维数组和多维数组
	3.C语言中定义数组的格式
	int ages[3];
	4.Go语言中定义数组的格式
	var ages [3]int
	5.C语言中操作数组的格式:
	ages[索引];
	6.Go语言中操作数组的格式和C语言一样

	7.C语言中数组的初始化方式有好几种
	7.1先定义再初始化
	只能逐个初始化, 不能一次性初始化
	7.2定义的同时初始化
	int ages[3] = {1, 3, 5}  // 完全初始化
	int ages[3] = {1, 3} // 部分初始化
	int ages[3] = {[2] = 3} // 指定元素, 部分初始化
	int ages[] = {1, 3, 5} // 如果定义的同时初始化, 元素个数可以省略

	8.Go语言中数组的初始化方式有好几种
	8.1先定义再初始化
	Go语言中的数组可以先定义再一次性初始化
	8.2定义的同时初始化

	注意点:
	1.给数组一次性赋值的时候一定要在值的前面加上数组的类型
	var 数组名称 数据类型 = 数据类型{x , y, z}
	而且必须要和数据类型一模一样
	2.C语言中的数组如果没有初始化保存的是垃圾数据, 但是Go语言的数组 没有初始化保存的不是垃圾数据, 而是对应类型的零值


	零值:
	int/int8/int16/int32/int64/uint/uint8/uint16/uint32/uint64/byte/rune/uintptr的默认值是0
	float32/float64的默认值是0.0
	bool的默认值是false
	string的默认值是""
	pointer/function/interface/slice/channel/map/error的默认值是nil
	 */

	 /*
	 // 1.定义一个数组
	 var ages [3]int
	 // 2.给数组指定索引(下标)的元素赋值
	 //ages[0] = 123
	 //ages[1] = 456
	 //ages[2] = 789
	 // Go语言中的数组可以先定义再一次性初始化
	 ages = [3]int{1, 3, 5}
	 */

	 // 1.定义的同时初始化
	//var ages [3]int = [3]int{1, 3, 5} // 完全初始化
	//var ages [3]int = [3]int{1, 3} // 部分初始化
	//var ages [3]int = [3]int{2:5} // 指定元素初始化
	//var ages []int = []int{1, 3, 5} // 省略元素个数
	//var ages = []int{1, 3, 5} // 省略元素个数, 这里不是数组是切片
	//ages := []int{1, 3, 5} // 省略元素个数, 这里不是数组是切片
	// [...]的含义, 让编译器根据赋值的元素个数自动计算数组的元素个数
	ages := [...]int{1, 3, 5} // 省略元素个数

	 // 3.取出数组中保存的值
	 fmt.Println(ages)
	 fmt.Println(ages[0])
	 fmt.Println(ages[1])
	 fmt.Println(ages[2])
}
