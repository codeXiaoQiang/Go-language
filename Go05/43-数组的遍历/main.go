package main

import "fmt"

func main() {
	/*
	1.Go语言中遍历数组和C语言一样

	2.除了可以利用传统的for循环遍历数组以外, 还可以使用高级for循环来遍历数组
	(相当于其他语言中的forin循环)
	 */
	 ages := [3]int{1, 3, 5}
	 //for i:=0; i < 3; i++ {
	 //	//fmt.Println(ages[i])
	 //	fmt.Printf("index = %d, value = %d\n", i, ages[i])
	 //}

	 for index, value := range ages {
	 	fmt.Println(index, value)
	 }
}
