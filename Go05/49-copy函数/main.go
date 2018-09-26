package main

import "fmt"

func main() {
	//sce1 := []int{1, 3, 4} // 源切片
	//sce1 := []int{1, 3, 4, 5} // 源切片
	sce1 := []int{1, 3} // 源切片
	sce2 := []int{0, 0, 0} // 目标切片
	fmt.Println(sce2)
	// 第一个参数:目标切片
	// 第二个参数:源切片
	// 会将源切片中的数据拷贝到目标切片中
	// 拷贝的时候会以目标切片为主
	copy(sce2, sce1)
	fmt.Println(sce2)
}
