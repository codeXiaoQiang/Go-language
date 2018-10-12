package main

import "fmt"

func main() {
	/*
	默认情况下所有的管道都是双向的管道(可读可写)
	那么在企业开发中, 我们可能会需要将管道作为函数的参数, 并且还需要限制函数中如何使用管道,
	那么这个时候我们就可能会使用单向管道

	双向格式:
	var myCh chan int;
	myCh = make(chan int, 5)
	myCh = make(chan int)

	单向格式:
	var myCh chan<- int; 只写的管道
	var myCh <-chan int; 只读的管道

	双向管道可以转换为单向的管道
	但是单向的管道不能转换为双向的管道
	 */

	 /*
	 // 1.定义一个双向的管道
	var myCh chan int = make(chan int, 5)
	// 2.声明一个只读的单向管道
	var myCh2 <-chan int
	// 3.将双向管道赋值给单向管道
	myCh2 = myCh
	fmt.Println(myCh2)
	// 4.声明一个只写的单向管道
	var myCh3 chan<- int
	// 5.将双向管道赋值给单向管道
	myCh3 = myCh
	fmt.Println(myCh3)
	 */

	var myCh2 <-chan int = make(<-chan int, 5)
	//myCh2<-1 // 会报错
	fmt.Println(myCh2)

	var myCh3 chan<- int = make(chan<- int, 5)
	//<-myCh3 // 会报错
	//fmt.Println(myCh3)
	//myCh3 = myCh2
	//myCh2 = myCh3
	fmt.Println(myCh3)
}
