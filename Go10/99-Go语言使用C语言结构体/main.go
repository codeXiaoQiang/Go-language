package main
/*
#include <stdio.h>
struct Point {
    float x;
    float y;
};
*/
import "C"
import "fmt"

func main() {
	// 注意点: 如果结构体属性是复杂类型, 可能会有问题
	var pi C.struct_Point = C.struct_Point{1.1, 2.2}
	fmt.Println(pi.x)
	fmt.Println(pi.y)
}
