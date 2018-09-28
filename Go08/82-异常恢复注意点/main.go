package main

import "fmt"

func main() {

	/*
	捕获异常注意点:
	1.同一个函数中, 多个panic异常, 只有第一个会被捕获
	 */
	 /*
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err) // 异常1
		}
	}()
	panic("异常1")
	panic("异常2")
	panic("异常3")
	panic("异常4")
	 */
	 test2()
}
func test2()  {
	// 如果有异常写在defer中, 但是defer后面还有其它异常, 那么捕获到的是其它的异常
	// 如果其它异常是写在defer前面, 那么和同一个函数中, 多个panic异常, 只有第一个会被捕获
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err) // 异常A
		}
	}()

	defer func() {
		panic("异常B")
	}()
	panic("异常C")
}