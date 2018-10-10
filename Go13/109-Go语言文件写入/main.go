package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	第一种写入的方式:
	通过os包中的Write函数读取
	 */

	 // 1.打开文件
	 // 注意点: Open函数打开的文件只能读取, 不能写入
	 //fp, err := os.Open("D:/lnj.txt")
	 // Create函数的作用: 创建一个文件
	 // 注意点: 文件不存在, 会自动创建一个新的
	 //         文件存在, 会覆盖以前的(相当于创建了一个新的替换掉了以前的旧的)
	 fp, err := os.Create("D:/lnj.txt")
	 if err != nil{
	 	fmt.Println("打开文件失败")
		 return
	 }
	 // 2.关闭文件
	 defer func() {
	 	if err = fp.Close(); err !=nil{
	 		fmt.Println("关闭文件失败")
		}
	 }()
	 /*
	 // 3.往文件中写入数据
	 // 特点: 会将指定的数据一次性写入到文件中
	 // 所以不适合数据量比较大的情况
	 buf := []byte{'l','n','j','\r','\n'}
	 len, err :=  fp.Write(buf)
	 if err !=nil{
	 	fmt.Println("写入失败")
	 }else{
	 	fmt.Println("写入成功, 写入了", len, "个字节")
	 }
	*/
}
