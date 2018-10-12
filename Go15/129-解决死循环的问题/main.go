package main

import (
	"fmt"
)

func main() {
	// 1.定义一个有缓冲区的管道
	myCh := make(chan int, 5)
	// 2.定义一个没有缓冲的管道
	exitCh := make(chan bool)

	go func() {
		for i:= 0; i < 5; i++ {
			myCh<-i
			fmt.Println("生产了", i)
		}
		exitCh<-true
	}()
	//for{
	//	;
	//}
	//time.Sleep(time.Second)
	<-exitCh
}
