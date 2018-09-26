package main

import "fmt"
import "test"

// 2.
const MAX_VALUE  = 666

// 3.
var num int = getNumber()

// 4.
func init()  {
	fmt.Println("init函数1")
}
func init()  {
	fmt.Println("init函数2")
}
// 5.
func main() {
	/*
	1.在Go语言中保留了两个函数, 一个叫做main,另外一个叫做init
	main函数是由系统调用的, init函数也是由系统调用的

	2.init函数的作用:
	对当前文件进行初始化

	3.调用顺序
	main包 --> 常量 --> 全局变量 --> init函数 --> main函数 --> 执行逻辑代码 --> Exit程序

	4.注意点:
	main函数只能在main包中(package main), 并且一个程序只能有一个,
	但是init函数每个包中都可以有, 并且可以有多个(但是企业开发推荐只写一个)
	 */
	 fmt.Println("main函数")

	 test.getValue()
}

func getNumber() int {
	fmt.Println("返回全局变量需要的值")
	return 888
}
