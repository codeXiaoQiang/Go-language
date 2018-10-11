package main

import (
	"fmt"
	"runtime"
	"time"
)

func singer()  {
	for i:=0; i < 10; i++{
		fmt.Println("我在唱歌", i)
		time.Sleep(time.Millisecond * 500)
		//runtime.Gosched() // 出让本次的执行权限
		// 应用场景: 让某些协程执行的次数少一次, 让某些协程执行的次数多一些

		// 唱5次就不唱了
		if i == 5{
			//break
			//return
			// 终止调用它的go程。其它go程不会受影响
			runtime.Goexit()
		}
	}
}
func dance(){
	for i:=0; i < 10; i++ {
		fmt.Println("----我在跳广场舞", i)
		time.Sleep(time.Millisecond * 500)
		//runtime.Gosched()
	}
}
func main() {
	//go singer()
	//go dance()
	//for{
	//	;
	//}

	// 了解就可以了
	// Go1.8之后, 系统会自动设置为最大值
	// num := runtime.NumCPU()
	// runtime.GOMAXPROCS(num)

	// 获取本地机器的逻辑CPU个数。
	//num := runtime.NumCPU()
	//fmt.Println(num)

	// 设置可同时执行的最大CPU数
	// 返回上一次设置的值
	//num := runtime.GOMAXPROCS(1)
	//fmt.Println(num)
	//num = runtime.GOMAXPROCS(8)
	//fmt.Println(num)
}
