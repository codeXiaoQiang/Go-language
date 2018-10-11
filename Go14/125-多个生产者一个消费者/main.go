package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义一个数组模拟缓冲区
var buff = make(chan int, 5)

// 定义一个函数模拟生产者
func producer(name string)  {
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	for i:=0; i<10;i++  {
		// 产生随机数
		num := rand.Intn(100)
		fmt.Println(name, "生产者生产了", num)
		// 将生产好的数据放入缓冲区
		buff<- num
		//time.Sleep(time.Millisecond)
	}
}

// 定义一个函数模拟消费者
func consumer(name string)  {
	for i:=0; i<10;i++  {
		num := <-buff
		fmt.Println("-----", name,"-------消费者消费到了", num)
	}
}
func main() {

	// 注意点: print函数是非常耗时的, 所以可能会出现看到的数据混乱的问题
	go producer("生产者一")
	go producer("生产者2")
	go consumer("消费者一")
	go consumer("消费者2")
	for{
		;
	}
}
