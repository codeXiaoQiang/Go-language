package main

import "fmt"

func main() {
	/*
	1.Go语言中的二维数组和C语言中一样
	2.定义二维数组的格式
	var 数组名称 [行数][列数]数据类型
	var 数组名称 [一维数组的个数][每个一维数组的元素个数]数据类型

	注意点:
	二维数组也可以省略元素个数, 但是只能省略行数, 不能省略列数
	 */
	 /*
	 //ages := [3]int{1, 3, 5}
	 //nums := [3]int{2, 4, 5}

	 var values [2][3]int = [2][3]int{
		 {1, 3, 5},
		 {2, 4, 5},
	 }

	values[0][1] = 666
	fmt.Println(values)

	 //for i:=0; i<2;i++{
	 //	for j:=0;j<3;j++{
	 //		fmt.Println(values[i][j])
		//}
	 //}
	 */

	//values := [...][3]int{
	//	{1, 3, 5},
	//	{2, 4, 5},
	//}

	// 这里代表定义了一个数组, 数组中保存的是数组
	//values := [...][3]int{
	//	{1, 3, 5},
	//	{2, 4, 5},
	//}

	// 这里代表定义了一个切片, 切片中保存的是数组
	//values := [][3]int{
	//	{1, 3, 5},
	//	{2, 4, 5},
	//}

	// 这里代表定义了一个切片, 切片中保存的是切片
	values := [][]int{
		{1, 3, 5},
		{2, 4, 5},
	}
	fmt.Println(values)
}
