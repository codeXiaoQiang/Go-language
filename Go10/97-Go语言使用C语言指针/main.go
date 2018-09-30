package main
/*
#include <stdio.h>
int num = 123;
int *p1 = &num;
void *p2 = &num;
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	/*
	var p1 *C.int = C.p1
	//fmt.Printf("%T\n", p1)
	//fmt.Println(*p1)
	//var num = *p1
	//fmt.Println(num)
	*/

	// 注意点: 其它类型都可以按照上面的格式来表示
	// 但是void *类型, 必须使用Go语言中的unsafe.Pointer来保存
	//var p2 *C.void = C.p2
	var p2 unsafe.Pointer = C.p2
	fmt.Println(p2)
}
