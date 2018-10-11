package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 创建一把互斥锁
var lock = sync.Mutex{}

// 定义一个数组模拟缓冲区
var buff [10]int

// 定义一个函数模拟生产者
func producer()  {
	// 上锁
	// 这里锁定的是当前go程, 也就是当前函数
	// 意味着其它的go程不能执行当前的函数, 只有当前锁定这个函数的go程才能执行这个函数
	lock.Lock()
	/*
	为什么在生产者和消费者中都上锁之后, 就可以实现生产完再消费?
	因为生产者和消费者中的锁都是同一把, 都是全局变量lock

	调用Lock()函数的作用: 是修改Mutex结构体中的state属性的值, 将它改为一个非0的值
	每次上错的时候都会判断有没有被锁定, 如果已经锁定就不锁了, 并且不会执行后面的代码
	调用Unlock函数的作用: 是修改Mutex结构体中的state属性的值, 将它改为0
	 */
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	for i:=0; i<10;i++  {
		// 产生随机数
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		// 将生产好的数据放入缓冲区
		buff[i] = num
		time.Sleep(time.Millisecond * 300)
	}

	// 解锁
	lock.Unlock()
}

// 定义一个函数模拟消费者
func consumer()  {
	// 上锁
	lock.Lock()
	for i:=0; i<10;i++  {
		num := buff[i]
		fmt.Println("-------消费者消费到了", num)
	}
	// 解锁
	lock.Unlock()
}
func main() {

	go producer()
	// 我们想要的是, 只有生产者生产了, 我们才能消费
	// 注意点: 在多go程中, 如果生产者生产的太慢, 那么消费者就会消费到错误的数据
	//go producer()
	go consumer()
	// 注意点: 看上去通过给生产者以及消费者同时加锁就能解决, 只有生产完了才能消费
	//         但是取决于谁想执行加锁操作, 所以不完美
	for{
		;
	}
}
