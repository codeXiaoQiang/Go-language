package main

/*
#include <stdio.h>
char ch = 'b';
int num = 123;
double value = 3.14;
 */
import "C"
import "fmt"

func main() {
	// 通过C.变量名称 方式访问C语言中的变量
	//fmt.Println(C.ch)
	//fmt.Println(C.num)
	//fmt.Println(C.value)

	//fmt.Printf("%T\n", C.ch)
	//fmt.Printf("%T\n", C.num)
	//fmt.Printf("%T\n", C.value)

	var ch byte
	ch = byte(C.ch)
	fmt.Println(ch)

	var num int
	num = int(C.num)
	fmt.Println(num)

	var value float64
	value = float64(C.value)
	fmt.Println(value)
}
