package main
/*
#include <stdio.h>
enum Gender {
   GenderMale,
   GenderFemale,
   GenderYao
};
*/
import "C"
import "fmt"

func main() {
	// Go语言访问C语言中的枚举和访问普通变量一样, 通过C.枚举值 方式即可
	// C语言中的枚举类型, 在Go语言中对应的是 enum_枚举类型名称
	var gender C.enum_Gender = C.GenderYao
	fmt.Println(gender)
}
