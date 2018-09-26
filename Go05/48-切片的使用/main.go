package main

import "fmt"

func main() {
	// 1.定义一个切片
	//sce := []int{1, 3, 5, 7, 9}

	// 2.修改切片中的值
	// 切片名称[索引] = 值
	//fmt.Println(sce)
	//sce[1] = 666
	//fmt.Println(sce)

	// 3.给切片增加数据
	//fmt.Println(sce)
	//sce = append(sce, 666)
	//fmt.Println(sce)

	// 4.查询切片中的数据
	//for _, value := range sce {
	//	if value == 3 {
	//		fmt.Println("包含数字3")
	//	}
	//}

	/*
	// 5.删除切片中指定的元素
	// {1, 3, 5, 7, 9}
	index := 2 // 需要删除元素的索引
	// 注意点: 不仅能够通过数组生成切片, 还可以通过切片生成切片
	//sce = sce[:index] // {1, 3}
	//fmt.Println(sce)
	//sce = sce[index + 1:] //{7, 9}
	//fmt.Println(sce)
	// 发现规律: 将第二次截取的切片添加到第一次截取的截屏后面即可
	// {1, 3, 7, 9}
	sce = append(sce[:index], sce[index + 1:]...)
	fmt.Println(sce)
	*/

	/*
	// 通过切片生成切片的注意点
	// 1.创建一个数组
	var arr [5]int = [5]int{1, 3, 5, 7, 9} // 数组
	// Go语言中不能直接通过数组名获取数组的地址
	//fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])

	// 2.通过数组创建一个切片
	var sce []int = arr[:]
	// 直接打印sce打印的是sce中指针保存的地址
	// 也就是底层指向的那个数组的地址
	fmt.Printf("%p\n", sce)

	// 3.通过切片创建一个切片
	// 通过切片创建一个切片, 新的切片和老的切片底层指向同一个数组
	var sce2 []int = sce[:]
	fmt.Printf("%p\n", sce2)

	// 4.直接修改数组的值
	//arr[1] = 666
	//sce[1] = 666
	sce2[1] = 666
	fmt.Println(arr)
	fmt.Println(sce)
	fmt.Println(sce2)
	*/


	// 1.什么是可变参数? 底层其实就是一个切片
	//res := sum(10, 20, 30)
	//arr := [3]int{10, 20, 30}
	//res := sum(arr)
	//sce := []int{10, 20, 30}
	//res := sum(sce)
	//fmt.Println(res)

	res:=sum( 10, 20, 30)
	//res:=sum(3.14, 10, 20, 30)
	fmt.Println(res)

	//sce := []int{1, 3, 5, 7, 9}
	//sce = append(sce[:index], sce[index + 1:]...)
}

// 注意点: 如果函数编写了可变参数, 那么可变参数只能放在形参列表的最后
func sum( nums ...int) int{
//func sum(value float32, nums ...int) int{
	//fmt.Println(a)
	//fmt.Printf("%T", a)
	var res int
	for _,value := range nums {
		res += value
	}

	fmt.Println(value)
	return res
}

//func sum(nums []int) int{
//	// 2.遍历切片
//	var res int
//	for _, value := range nums{
//		res += value
//	}
//	return res
//}

//func sum(nums [3]int) int{
//	// 1.获取数组的长度
//	length := len(nums)
//	var res int // 保存计算结果
//	// 2.遍历数组
//	for i := 0; i < length; i++ {
//		res += nums[i]
//	}
//	return res
//}

//func sum(a int, b int, c int) int{
//	var res int = a + b + c
//	return  res
//}
