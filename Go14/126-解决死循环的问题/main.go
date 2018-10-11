package main

import "fmt"

func main() {
	var myCh = make(chan int, 3)
	var exitCh = make(chan bool, 1)
	go func() {
		for i:=0; i<3; i++ {
			fmt.Println("生产了数据", i)
			myCh<-i
		}
		exitCh<-true
	}()
	fmt.Println("exitCh之前的代码")
	<-exitCh // 阻塞
	fmt.Println("exitCh之后的代码")

	//for{
	//	;
	//}
}
