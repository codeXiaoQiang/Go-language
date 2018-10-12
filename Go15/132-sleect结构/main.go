package main

import (
	"fmt"
	"time"
)

func main() {
 	// 1.创建一个管道
 	myCh1 := make(chan int, 5)
 	//myCh2 := make(chan int, 5)
	//exitCh := make(chan bool)

 	// 2.开启一个协程生产数据
 	go func() {
 		time.Sleep(time.Second * 5)
		for i := 0; i < 10 ; i++ {
			myCh1<-i
			fmt.Println("生产者1生产了", i)
		}
		close(myCh1)
		//exitCh<-true
	}()
 	/*
	go func() {
		time.Sleep(time.Second * 5)
		for i := 0; i < 10 ; i++ {
			myCh2<-i
			fmt.Println("生产者2生产了", i)
		}
		close(myCh2)
	}()

 	// 2.在主线程中消费数据
	//for i := 0; i < 10 ; i++ {
	//	num := <-myCh
	//	fmt.Println("------消费者消费了", num)
	//}

	//for num := range myCh{
	//	fmt.Println("------消费者消费了", num)
	//}
 	*/

	// 注意点: 在企业开发中, 一般情况下使用select都是用于同时消费多个管道中数据
	//         在企业开发中, 一般情况下select中的default不用写
	//         在企业开发中, 一般情况下使用select来控制退出主线程
	//		   在企业开发中, 一般情况下使用select来处理超时
	for{
		//fmt.Println("start")
		select {
		case num1 := <-myCh1:
			fmt.Println("------消费者消费了myCh1", num1)
		//case num2 := <-myCh2:
		//	fmt.Println("------消费者消费了myCh2", num2)
		//case <-exitCh:
		//	return
		case <-time.After(3):
			fmt.Println("超时了")
			return
		//default:
		//	fmt.Println("生产者还没有生产好数据")
		}
		//fmt.Println("=====================")
		time.Sleep(time.Millisecond)
	}
	fmt.Println("程序结束了")
}
