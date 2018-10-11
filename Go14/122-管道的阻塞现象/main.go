package main

import "fmt"

var myCh = make(chan int, 3)

func test()  {
	/*
	myCh<-1
	myCh<-2
	myCh<-3
	fmt.Println("满之前的代码")
	// 这里不会报错, 会阻塞, 等到有人将管道中的数据读出去之后, 有的新的空间再往管道中写
	myCh<-4
	fmt.Println("满之后的代码")
	*/
	fmt.Println("前的代码")
	// 这里不会报错, 会阻塞, 等到有人往管道中的写入数据之后, 等到有了新的数据才会去读取
	<-myCh
	fmt.Println("后的代码")
}

func demo()  {
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
}
func main() {
	/*
	单独在主线程中操作管道, 写满了会报错, 没有数据去读取也会报错
	只要在go程中操作管道, 无论有没有写满, 无论有没有数据都会发生管道阻塞的现象
	 */

	 // 1.单独在主线程中操作管道
	 //myCh := make(chan int, 3)
	 //myCh<-1
	 //myCh<-2
	 //myCh<-3
	 //myCh<-4 // 会报错

	//<-myCh // 会报错

	// 2.在go程中操作管道
	//go test()


	// 3.既在主线程操作, 又在go程中操作管道
	myCh<-123
	go demo()

	for{
		;
	}

}
