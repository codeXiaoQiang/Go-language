package main

import (
	"fmt"
	"sync"
	"time"
)
// 1.创建一个互斥锁
// 添加了互斥锁之后, 我们的并发的程序就编程了串行的程序了
var lock = sync.Mutex{}

// 逐个打印字符的函数
func printer(str string)  {
	// 添加锁(关门)
	lock.Lock()
	for _, ch := range str{
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	// 释放锁(开门)
	lock.Unlock()
}

func person1()  {
	printer("hello")
}
func person2()  {
	printer("world")
}

func main() {
	/*
	需求: 定义个打印的函数, 可以逐个的打印传入字符串的每个字符
	在两个协程中调用这个好方法, 并且还要保持输出的内容有序的
	 */
	go person1()
	go person2()
	for{
		;
	}
}
