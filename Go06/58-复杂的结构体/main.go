package main

import "fmt"

func main() {
	type Person struct {
		age int
		score float32
		name string
		arr [3]int
		sce []int
		dict map[string]string
	}

	var per Person
	// 常规的数据类型, 我们都可以直接操作, 完全没有问题
	per.age = 18
	per.score = 100
	per.name = "lnj"
	per.arr[0] = 1
	per.arr[1] = 3
	per.arr[2] = 5
	fmt.Println(per)
	// 注意点: 如果结构体的属性是切片和字典类型, 那么就不能直接操作
	// 必须先给切片初始化(创建切片)
	per.sce = make([]int, 2)
	per.sce[0] = 123
	per.sce[1] = 456
	//per.sce[2] = 789 // 通过索引操作切片, 不能超出切片的长度
	fmt.Println(per)

	per.dict = make(map[string]string)
	per.dict["name"] = "zs"
	fmt.Println(per)
}
