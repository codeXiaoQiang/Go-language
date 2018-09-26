package main

import "fmt"

func main() {
	/*
	1.什么是匿名函数
	没有名字的函数我们就称之为匿名函数

	2.注意点:
	匿名函数要么定义之后直接调用, 要么保存到一个变量当中, 否则就会报错
	 */
	 /*
	 //func(){
	 //	fmt.Println("匿名函数")
	 //}()
	 */

	 /*
	 //fn := func(){
		// fmt.Println("匿名函数")
	 //}
	 //fmt.Printf("%T\n", fn) // func()
	 //fn()

	 //fn := func(a int, b int) int{
	 //	return a + b
	 //}
	 */

	 /*
	 // 代表定义了一个名称叫做fn的变量, 这个变量的类型是一个函数  func(int, int) int
	 // 这个函数可以接收两个int类型的形参, 并且可以返回一个int类型的结果
	 // 也就是说变量fn将来可以保存所有接收两个int类型形参返回一个int类型结果的函数
	 var fn func(int, int) int
	 fn = func(a int, b int) int {
		 return a + b
	 }

	 // var fn func(int, int) int = func(a int, b int) int {
		// return a + b
	 //}
	 //相当于
	 //func fn(a int, b int) int {
		// return a + b
	 //}

	 //fmt.Printf("%T\n", fn) // func(int, int) int
	 res := fn(10, 20)
	 fmt.Println(res)

	 // 定义变量的格式: var 变量名称 数据类型
	 // 例如:  var num int
	 // 代表定义了一个int类型的变量, 将来可以保存int类型的值
	*/


	/*
	var fn func(int,int)int
	fn = func(a int, b int) int{
		return a + b
	}
	var method func(int, int)int
	method = fn
	var res int
	res = method(10, 20)
	fmt.Println(res)
	*/
	/*
	var fn func(int,int)int
	fn = func(a int, b int) int{
		return a + b
	}
	test(fn)
	*/

	/*
	var fn func(int,int)int
	fn = func(a int, b int) int{
		return a + b
	}
	result := test(10, 20, fn)
	fmt.Println(result)
	*/


	result := test(10, 20, func(a int, b int) int{
		return a + b
	})
	fmt.Println(result)

}

func test(num int, value int, method func(int, int)int) int {
	var res int
	res = method(num, value)
	//fmt.Println(res)
	return  res
}

/*
func test(method func(int, int)int)  {
	var res int
	res = method(10, 20)
	fmt.Println(res)
}
*/
