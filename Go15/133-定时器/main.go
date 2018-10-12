package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	type Timer struct {
		C <-chan Time
		r runtimeTimer
	}
	 */
	// 1.想使用定时器, 就需要用到time包
	/*
	 // NewTimer作用, 就是让系统在指定时间之后, 往Timer结构体的C属性中写入当前的时间
	 // 让程序阻塞3秒, 3秒之后再执行
	 //start := time.Now()
	 //fmt.Println(start)
	 //timer := time.NewTimer(time.Second * 3) // Timer
	 //fmt.Println(<-timer.C)

	//start := time.Now()
	//fmt.Println(start)
	//timer := time.After(time.Second * 3) // Timer.C
	//fmt.Println(<-timer)
	*/
	/*
	// 以上的定时器都是一次性的定时器, 也就是只会执行一次
	go func() {
		start := time.Now()
		fmt.Println(start)
		timer := time.After(time.Second * 3) // Timer.C
		for{
			fmt.Println(<-timer)
		}
	}()

	for{
		;
	}
	*/

	// 2.周期性的定时器
	start := time.Now()
	fmt.Println(start)
	ticker := time.NewTicker(time.Second * 2)
	for{
		fmt.Println(<-ticker.C)
	}
}
