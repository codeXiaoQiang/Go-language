package main

import "fmt"

func main() {
	/*
	1.在C语言中循环结构有三个 for/while/dowhile
	2.在Go语言中循环结构只有一个 for
	因为while/dowhile能做的for都能做, 所以Go只保留了for循环

	3.C语言for循环的格式
	for(初始化表达式;条件表达式;增量表达式){
	    需要循环操作的语句;
	}

	4.Go语言的for循环和C语言非常非常非常像
	for 初始化表达式;条件表达式;增量表达式{
	    需要循环操作的语句;
	}

	特点:
	1.没有非零即真的概念, 条件表达式的返回值必须是布尔类型
	2.不能省略大括号
	3.最简单的死循环
	for{}
	 */
	 //fmt.Println("发射子弹")
	 //fmt.Println("发射子弹")
	 //fmt.Println("发射子弹")
	 //fmt.Println("发射子弹")
	 //fmt.Println("发射子弹")

	 for num := 0; num < 5; num++{
		 fmt.Println("发射子弹")
	 }
}
