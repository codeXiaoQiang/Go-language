package main

import (
	"fmt"
	"time"
)
// 利用管道的阻塞现象实现串行
var myCh = make(chan bool, 1)

// 逐个打印字符的函数
func printer(str string)  {
	for _, ch := range str{
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1()  {
	printer("hello")
	// 往管道中写入数据
	// 只有printer函数执行完毕才会往管道中写入数据
	myCh<-true
}
func person2()  {
	// 从管道中读取数据
	<-myCh // 因为是在go程中读取的, 并且从来没有人写入过数据, 所以会阻塞
	printer("world")
}

func main() {
	go person1()
	go person2()
	for{
		;
	}
}
