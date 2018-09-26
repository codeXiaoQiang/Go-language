package main

import "fmt"

func main() {
	/*
	1.数组在Go语言中是值类型, 所以数组之间的赋值是拷贝关系, 而不是指向的关系
	2.如果想让两个数组可以相互赋值, 那么两个数组的类型必须一致
	类型一致包括元素个数, 元素个数也是Go语言中数组类型的一部分
	3.如果数组中存储的元素类型支持== != 操作,而且两个数组的类型是一个样的,  那么数组也支持== !=操作
	 */
	 /*
	 ages := [3]int{1, 3, 5}
	 var nums [3]int

	 nums = ages
	 nums[1] = 666

	 fmt.Println(ages)
	 fmt.Println(nums)
	 */

	 /*
	ages := [3]int{1, 3, 5}
	change(ages)
	fmt.Println(ages)
	 */

	 /*
	var ages [3]int = [3]int{1, 3, 5}
	var nums [3]int
	nums = ages
	fmt.Println(ages)
	fmt.Println(nums)
	 */

	var ages [3]int = [3]int{1, 3, 5}
	var nums [3]int = [3]int{1, 2, 5}
	res := ages == nums
	fmt.Println(res)
}
func change(nums [3]int)  {
	nums[1] = 666
}
