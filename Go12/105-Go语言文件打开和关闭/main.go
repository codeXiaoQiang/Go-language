package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	1.如何打开文件?
	在Go语言的os包中提供了一个函数, 叫做Open, 就是专门用于打开某一个文件的
	注意点:
	如果文件不存在不会自动创建
	通过open函数打开的文件只能读取, 不能写入

	2.如何关闭文件?
	在Go语言的os包中提供了一个函数, 叫做Close, 就是专门用于关闭某个打开文件的

	 */

	 // 1.打开文件
	 //os.Open("D:\\lnj.txt")
	 fp, err := os.Open("D:/lnj.txt") // 推荐写法
	 if err != nil {
	 	fmt.Println("打开失败")
	 }else{
	 	fmt.Println(fp)
	 }
	 // 2.关闭文件
	 defer func() {
	 	if err = fp.Close(); err !=nil{
	 		fmt.Println("关闭文件失败")
		}
	 }()

}
