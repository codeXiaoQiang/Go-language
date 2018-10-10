package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	第一种读取的方式:
	通过bufio包中的ReadBytes和ReadString函数
	func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
	 */
	// 1.打开文件
	fp, err := os.Open("D:/lnj.txt")
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	// 2.关闭文件
	defer func() {
		if err = fp.Close(); err != nil{
			fmt.Println("关闭文件失败")
		}
	}()

	// 创建缓冲区
	// 将打开的文件句柄传递给NewReader函数, 会返回一个新的句柄
	// 缓冲区默认的大小是4096
	r := bufio.NewReader(fp)
	/*
	// 利用缓冲区的句柄来读取数据
	// 接收一个参数, 这个参数用于告诉ReadBytes函数, 读取到什么地方接收
	// 会将读取到的数据放到一个切片中返回给我们
	//buff, err :=  r.ReadBytes('\n')
	//if(err != nil){
	//	fmt.Println("读取失败")
	//}else{
	//	fmt.Println(string(buff))
	//}
	*/

	/*
	abcdefg\n
	1234567
	第一次读取: buf = abcdefg\n
	第二次读取: buf = 1234567
	但是由于第二次读取一直读到了文件的末尾都没有读到\n, 所以ReadBytes函数就给err赋值了一个io.EOF
	所以如果1234567后面没有\n, 并且是先判断错误, 那么就会少输出一次
	 */

 /*
	for{
		buff, err :=  r.ReadBytes('\n')
		fmt.Print(string(buff))
		if err == io.EOF{
		//if err != nil{
			//fmt.Print("读取失败")
			//fmt.Print(err)
			break
		}

	}
 */

    str, err := r.ReadString('\n')
    if(err != nil){
    	fmt.Println("读取失败")
	}else{
		fmt.Print(str)
	}
}
