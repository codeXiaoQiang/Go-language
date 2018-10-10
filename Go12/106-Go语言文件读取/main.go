package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	第一种读取的方式:
	通过os包中的Read函数读取
	func (f *File) Read(b []byte) (n int, err error)
	Read函数会将读取到的数据放到指定的切片中, 并且会一次性将所有的数据都读取进来
	会将读取到的byte个数返回给我们, 并且还会返回一个error,
	如果读取成功, 那么error等于nil, 如果读取失败那么error就不等于nil

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
	 // 3.从文件中读取数据
	 // read函数接收一个切片, 会将读取到的数据放到指定的切片中
	 // 注意点: 会一次性将文件中的数据全部读取进来
	 //buff := make([]byte, 666)
	 //len, err := fp.Read(buff)
	 //if err != nil{
		// fmt.Println("读取数据失败")
	 //}else{
	 //	fmt.Println("len = ", len)
	 //	//fmt.Println(string(buff))
	 //	fmt.Println(string(buff[:len]))
	 //}

	 /*
	 // 读取所有内容
	buff := make([]byte, 7)
	// 先从文件中读取7个字节
	len, err := fp.Read(buff)
	// 判断有没有读取到数据
	for ;len > 0;{
		// 将读到的数据转换为字符串打印
		fmt.Print(string(buff[:len]))
		// 再次利用Read函数去读取7个字节
		len, err = fp.Read(buff)
	}
	 */
	// 定义切片保存读取到的数据
	buff := make([]byte, 7)
	 for{
		 // 先从文件中读取7个字节
		 len, err := fp.Read(buff)
		 if err != nil || len <= 0{
		 	break;
		 }
		 // 将读到的数据转换为字符串打印
		 fmt.Print(string(buff[:len]))
	 }

}
