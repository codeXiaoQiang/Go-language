package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
	第三种读取的方式:
	通过bufio包中的Read函数读取
	 */

	 // 会将指定路径中的文件一次性读取进来
	 // 接收一个参数: 用于传递需要读取的文件路径
	 // 返回两个值:
	 // 第一个返回值: 读取到的数据
	 // 第二个返回值: 读取失败就不等于nil
	 buf, err := ioutil.ReadFile("D:/lnj.txt")
	 if err != nil{
	 	fmt.Println("读取失败")
	 }else{
	 	fmt.Println(string(buf))
	 }

}
