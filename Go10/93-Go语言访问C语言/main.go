package main
import "C"
/*
#include <stdio.h>

void say(){
	printf("hello world\n");
}
*/
import "C"

func main() {
	/*
	1.C语言的代码都需要利用单行注册或者多行注释注释起来
	2.在C语言代码紧随其后的位置写上import "C"
	3.就可以在go代码中通过C.函数名称 方式来访问C语言的函数

	注意点:
	1,C语言代码的注释和import "C"之间不写能任何其他的内容
	2.C语言的代码可以利用多行注释注释起来, 也可以利用单行注释注释起来
	 */
	 C.say()
}
