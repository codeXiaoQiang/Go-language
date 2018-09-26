package main

import "fmt"

//value := 789
func main() {
	/*
	1.相同的作用域内, 无论是全局变量还是局部变量, 都不能出现同名的变量
	2.变量离开作用域就不能使用
	3.局部变量如果没有使用, 编译会报错, 全局变量如果没有使用, 编译不会报错
	4.:=只能用于局部变量, 不能用于全局变量
	5.:=如果用于同时定义多个变量, 会有退化赋值现象
	 */

	 // 1.:=只能用于局部变量
	 //num := 666
	 //fmt.Printf("%d\n", num)

	 // 2.:=退化赋值现象
	 // 2.1 var num int
	 // 2.2 num = 123
	 num := 123
	 //num := 456 // 报错
	 // 如果通过:=定义多个变量, 但是多个变量中有的变量已经在前面定义过了, 那么只会对没有定义过的变量执行:=, 而定义过的变量只执行=操作
	 // 1. var value int
	 // 2. num = 456
	 // 3. value = 789
	num, value := 456, 789
	fmt.Printf("%d, %d", num, value)
}
