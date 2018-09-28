package main

import "fmt"

func main() {
	/*
	1.程序不要随意被终止, 只要不是程序不能运行了, 就尽量让改程序继续保持运行

	2.在Go语言中如果panic异常, 那么可以通过defer和recover来实现panic异常的捕获, 让程序继续执行

	注意点:
	1.defer和recover必须在panic抛出异常之前定义
	2.panic异常会随着函数的调用栈向外传递
	例如: A函数调用了B函数, B函数调用了C函数
	      如果在C函数中抛出了一个panic异常, 那么这个异常会一层一层的传递到B和A
	      也就是在B和A函数中也能捕获到这个异常
	 */

	defer func() {
		// defer无论所在的函数是正常结束,还是异常结束都会被执行
		// recover可以捕获panic异常
		if err := recover(); err != nil{
			fmt.Println("recover捕获到了", err)
		}
	}()

	res := div(10, 0)
	fmt.Println(res)
}
// 除法运算
func div(a, b int) (res int) {
	// 在当前函数中捕获
	//defer func() {
	//	// defer无论所在的函数是正常结束,还是异常结束都会被执行
	//	// recover可以捕获panic异常
	//	if err := recover(); err != nil{
	//		fmt.Println("recover捕获到了", err)
	//	}
	//}()
	if b == 0 {
		// 手动终止程序
		panic("除数不能为0")
	}else{
		res = a / b
	}
	// 无效
	//defer func() {
	//	// defer无论所在的函数是正常结束,还是异常结束都会被执行
	//	// recover可以捕获panic异常
	//	if err := recover(); err != nil{
	//		fmt.Println(err)
	//	}
	//}()
	return
}