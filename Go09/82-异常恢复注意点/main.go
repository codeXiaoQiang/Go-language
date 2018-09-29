package main

import "fmt"

func main()  {
	/*
	1.同一个函数中有多个终止异常, 那么只有第一个会被捕获
	2.如果defer中也有终止异常,但是defer前面也有终止异常, 那么只会捕获到defer前面的那个终止异常
	3.如果defer中也有终止异常,但是defer后面也有终止异常,那么只会捕获到defer中的那个终止异常
	 */
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	//panic("异常A")
	//panic("异常B")
	//panic("异常C")
	//panic("异常D")

	//panic("异常A")
	defer func() {
		panic("异常B")
	}()
	panic("异常A")
}