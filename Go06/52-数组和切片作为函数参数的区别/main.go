package main

import "fmt"

func main() {
	/*
	1.数组作为函数的参数, 是值传递, 在函数内修改形参, 不会影响函数外的实参
	2.切片作为函数的参数, 是地址传递(指针), 在函数内修改形参, 会影响到函数外的实参
	 */
	//var arr [3]int = [3]int{1, 3, 5}
	//fmt.Println(arr)
	//change1(arr)
	//fmt.Println(arr)

	var sce []int = []int{1, 3, 5}
	fmt.Println(sce)
	change2(sce)
	fmt.Println(sce)
}
func change2(nums []int)  {
	nums[1] = 666
}
func change1(nums [3]int)  {
	nums[1] = 666
}
