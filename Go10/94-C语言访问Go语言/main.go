package main

/*
#include <stdio.h>
void demo(){
// 声明Go语言的函数
extern void say();
say();
}
 */
import "C"

func main() {
	/*
	1.Go语言定义的这个函数, 和C语言中调用这个函数不能在同一个文件中
	2.将C语言的代码写到单行注释或者多行注释中
	3.在C语言代码紧随其后的位置写上import "C"

	4.在需要被C语言调用的Go语言函数对应的文件中写上import "C"
	5.在需要被C语言调用的Go语言函数的上方写上 //export Go语言函数名称
	5.在C语言代码中调用之前写上extern Go语言函数返回值类型 Go语言函数名称(Go语言函数形参列表);
	6.直接调用extern声明好的函数即可
	 */
	C.demo()
}

