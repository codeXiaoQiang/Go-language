package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义一个数组模拟缓冲区
//var buff [10]int
var buff = make(chan int, 5)

// 定义一个函数模拟生产者
func producer(name string)  {
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	for i:=0; i<5;i++  {
		// 产生随机数
		num := rand.Intn(100)
		fmt.Println(name, "生产者生产了", num)
		// 将生产好的数据放入缓冲区
		buff<-num
		//time.Sleep(time.Millisecond * 300)
	}
}

// 定义一个函数模拟消费者
func consumer(name string)  {
	for i:=0; i<5;i++  {
		num := <-buff
		fmt.Println(name, "-------消费者消费到了", num)
	}
}
var myCh = make(chan int, 5)

func main() {
	go producer("生产者1")
	go producer("生产者二")
	go consumer("消费者1")
	go consumer("消费者二")
	for{
		;
	}

	//var myCh = make(chan int, 5)
	/*
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	myCh<-1
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	myCh<-2
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	myCh<-3
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	myCh<-4
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	myCh<-5
	fmt.Println("len", len(myCh), "cap", cap(myCh))

	fmt.Println(<-myCh)
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	*/

	// 只在主线程中操作管道, 那么满了再添加就会报错, 那么没有数据获取也会报错
	/*
	//fmt.Println(<-myCh)
	//myCh<-1
	//myCh<-2
	//myCh<-3
	//myCh<-4
	//myCh<-5
	//myCh<-6
	*/

	// 只要在其它go程中操作管道, 那么满了再添加就会阻塞, 那么没有数据获取也会阻塞
	/*
	go func() {
		//fmt.Println(<-myCh)
		myCh<-1
		myCh<-2
		myCh<-3
		myCh<-4
		myCh<-5

		fmt.Println("123")
		//<-myCh
		myCh<-6
		fmt.Println("abc")
	}()
	for{
		;
	}
	*/
}
