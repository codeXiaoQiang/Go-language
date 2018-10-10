package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 将一个视频拷贝到另外一个文件夹中
	// 1.读取视频
	// 由于视频文件比较大, 所以应该使用带缓冲区的方式读取
	rp, err := os.Open("F:/demo/day04/abc.mp4")
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	defer func() {
		if err = rp.Close(); err != nil{
			fmt.Println("关闭文件失败")
		}
	}()

	// 2.写入读取到的视频
	// 由于视频文件比较大, 所以应该使用带缓冲区的方式写入
	wp, err := os.Create("D:/def.mp4")
	if err != nil{
		fmt.Println("创建文件失败")
		return
	}
	defer func() {
		if err = wp.Close(); err != nil{
			fmt.Println("关闭文件失败")
		}
	}()

	// 3.创建读取的缓冲区
	r := bufio.NewReader(rp)
	// 4.创建写入的缓冲区
	w := bufio.NewWriter(wp)

	// 5.利用io包中的copy函数实现文件的拷贝
	len, err := io.Copy(w, r)
	if err != nil{
		fmt.Println("拷贝文件失败")
	}else{
		fmt.Println("拷贝数据成功, 总共拷贝了", len, "个字节")
	}
}
