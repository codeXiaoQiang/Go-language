package main

import "fmt"

func main() {
	/*
	1.切片可以再生成新的切片, 两个切片底层指向同一个数组
	2.切片和数组不同, 切片不支持== !=操作
	3.在Go语言中字符串的底层实现就是切片, 所以可以通过字符串来生成切片
	 */
	 //var arr1 [3]int = [3]int{1, 3, 5}
	 //var arr2 [3]int = [3]int{1, 3, 5}
	 //res1 := arr1 == arr2
	 //fmt.Println(res1)


	//var sce1 []int = []int{1, 3, 5}
	//var sce2 []int = []int{1, 3, 5}
	//res2 := sce1 == sce2
	//fmt.Println(res2)

	//var str string = "www.it666.com"
	////str[0] = 'M'
	////sce := str[:]
	//sce := str[4:]
	//fmt.Printf("%T\n", sce)
	//fmt.Println(sce)

	var str string = "www.it666.com"
	//sce := make([]byte, 3)
	//copy(sce, str)
	////fmt.Println(sce)
	//fmt.Printf("%c\n", sce[0])
	//fmt.Printf("%c\n", sce[1])
	//fmt.Printf("%c\n", sce[2])

	sce := make([]byte, len(str))
	copy(sce, str)
	sce[0] ='M'
	fmt.Printf("%s\n", sce)
}
