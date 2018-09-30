package main
/*
#include <stdio.h>
#include <stdlib.h>

char *str1 = "www.it666.com";
char str2[20] = "www.it666.com";
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	/*
	//fmt.Println(C.str1)
	//fmt.Println(C.str2)

	var str1 string
	str1 = C.GoString(C.str1)
	fmt.Println(str1)

	// 注意点: 不能将数组直接通过GoString方法转换为Go语言的字符串
	var str2 string
	//str2 = C.GoString(C.str2)
	str2 = C.GoString(&C.str2[0])
	fmt.Println(str2)
	*/

	// 注意点: 如果将一个Go语言的字符串转换为C语言的字符串之后, 这个字符串是不受Go语言的GC控制的
	// GC 垃圾回收器
	var str string = "www.itzb.com"
	str2 := C.CString(str)
	fmt.Println(str2)
	// 手动释放C语言的字符串
	C.free(unsafe.Pointer(str2))
}
