package main

import (
	"fmt"
)

func singer()  {
	for i:=0; i < 100; i++{
		fmt.Println("我在唱歌", i)
	}
}
func dance(){
	for i:=0; i < 100; i++ {
		fmt.Println("----我在跳广场舞", i)
	}
}
func main() {
	// 默认情况下, 我们编写的程序是串行的
	// 只有前面一个函数执行完, 后面一个函数才会执行
	//singer()
	//dance()

	// 需求: 边唱歌边跳舞
	// 既然要边唱歌边跳舞, 是不是意味着执行singer函数的时候, 也要执行dance函数
	go singer() // 前面写上一个go 就代表着开启一个协程
	go dance()

	// 注意点: 开启协程之后, 主线程(进程)不能结束
	// 一旦主线程结束了, 那么程序就关闭了, 那么进程就不见了, 那么协程就不会被执行了
	//time.Sleep(time.Microsecond)
	for{
		;
	}
}
