package main

import "fmt"

func main() {
	/*
	一种是程序发生异常时, 立刻退出终止程序继续运行

	终止程序也分为两种:
	1.系统自动终止
	2.我们手动终止 (企业开发不常用)
	格式: panic("提示信息")
	 */

	// 系统自动终止
	//arr := [3]int{1, 3, 5}
	//for i := 0; i < 20; i++ {
	//	fmt.Println(arr[i])
	//}

	res := div(10, 0)
	fmt.Println(res)
}
// 除法运算
func div(a, b int) (res int) {
	if b == 0 {
		// 手动终止程序
		panic("除数不能为0")
	}else{
		res = a / b
	}
	return
}
