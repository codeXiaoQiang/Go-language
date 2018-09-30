package main

import (
	"Clib"
	"fmt"
	"time"
)

func main() {
	// 隐藏光标
	//Clib.HideCursor()
	//fmt.Println("hello world")
	//Clib.GotoPostion(3, 0)
	//time.Sleep(time.Minute)

	//for{
	//	res := Clib.Direction()
	//	fmt.Printf("res = %c\n", res)
	//
	//	//time.Sleep(time.Minute)
	//}

	fmt.Println("############\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"############\n")

	// 1.隐藏光标
	Clib.HideCursor()
	// 2.将关闭移动到指定的位置
	Clib.GotoPostion(4, 3)
	// 3.在指定位置输入内容
	fmt.Print("*")
	Clib.GotoPostion(5, 3)
	fmt.Print("@")
	time.Sleep(time.Minute)
}
