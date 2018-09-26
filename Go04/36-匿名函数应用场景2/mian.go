package main

import "fmt"

func main() {

	// 1.匿名函数作为函数的参数
	/*
	//var fn func()
	//fn = func() {
	//	fmt.Println("匿名函数")
	//}
	//fmt.Printf("%p\n", fn)
	//test(fn)

	test(func() {
		fmt.Println("匿名函数")
	})
	*/
	/*
	var fn func()
	fn = test()
	fn()
	*/
	var fn func(int, int)int
	fn = test()
	res := fn(10, 20)
	fmt.Println(res)
}

func test()func(int, int)int{
	return func(a int, b int)int {
		return a + b
	}
}


//func test() func() {
//	return func() {
//		fmt.Println("匿名函数")
//	}
//}

//func test(method func())  {
//	fmt.Printf("%p\n", method)
//	method()
//}
