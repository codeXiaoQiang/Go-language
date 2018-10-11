package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义一个数组模拟缓冲区
var buff = make(chan int, 5)

// 定义一个函数模拟生产者
func producer()  {
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	for i:=0; i<10;i++  {
		// 产生随机数
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		// 将生产好的数据放入缓冲区
		buff<- num
		//time.Sleep(time.Millisecond)
	}
}

// 定义一个函数模拟消费者
func consumer()  {
	for i:=0; i<10;i++  {
		num := <-buff
		fmt.Println("-------消费者消费到了", num)
	}
}
func main() {

	go producer()
	go consumer()
	for{
		;
	}
}
