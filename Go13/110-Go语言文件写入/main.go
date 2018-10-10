package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	第二种写入的方式:
	通过bufio包中的Write和WirteString函数读取
	 */

	 // 1.创建一个文件
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

	// 3.创建一个写入缓冲区
	w :=bufio.NewWriter(fp)

	// 4.写入数据
	//buf := []byte{'l','n','j','\r','\n'}
	//len, err := w.Write(buf)
	len, err := w.WriteString("www.it666.com\r\n")
	if err != nil{
		fmt.Println("写入失败")
	}else{
		fmt.Println("写入成功, 写入了", len, "个字节")
	}

	// 注意点:
	// 如果通过带缓冲区的方式写入数据, 那么必须在写入完毕之后刷新一下缓冲区, 才会将缓冲区中的数据真正的写入到文件中
	w.Flush()
}
