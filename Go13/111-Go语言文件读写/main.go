package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	os包中的Open函数只能读取不能写入
	os包中的Create函数可以读取也可以写入, 但是每次执行都会覆盖原有的文件, 每次执行都是一个空文件
	所以: 如果想要往文件中追加数据, 利用如上两个方法是不行的
	所以: 需要通过os包中的另一个函数: OpenFile

	作用: 打开一个文件, 并且文件不存在可以创建(注意是可以不是自动)
	第一个参数: 需要打开文件的路径
	第二个参数: 以什么模式打开文件 (只读/只写/可读可写/追加/)
	  			模式可以同时指定多个, 多个之间使用|分隔, 例如 O_CREATE | O_WRONLY
	第三个参数: 指定文件的权限, 只对Linux系统有效, 在Windows下无效

	func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

	const (
		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	)

	0.没有任何权限
	1.执行权限(如果是可执行程序, 可以运行)
	2.写权限
	3.写权限和执行权限
	4.读权限
	5.读权限和执行权限
	6.读权限和写权限
	7.读权限和写权限以及执行权限
	一般情况下写 0666

	注意点: OpenFile和Open函数一样, 打开文件之后需要手动关闭文件
	 */

	 fp, err := os.OpenFile("D:/lnj.txt", os.O_APPEND | os.O_RDWR, 0666)
	 if err != nil{
	 	fmt.Println("创建失败")
	 }else{
	 	fmt.Println(fp)
	 }
	 defer func() {
	 	if err := fp.Close(); err != nil{
	 		fmt.Println("关闭文件失败")
		}
	 }()

	//len, err := fp.WriteString("www.itzb.com")
	//if err != nil{
	//	fmt.Println("写入数据失败")
	//}else{
	//	fmt.Println("写入成功, 写入了", len, "个字节")
	//}

	//buf := make([]byte, 1024)
	//_, err = fp.Read(buf)
	//if err != nil{
	//	fmt.Println("读取数据失败")
	//}else{
	//	fmt.Println(string(buf))
	//}

	w := bufio.NewWriter(fp)
	len , err := w.WriteString("\r\n123456")
	if err != nil{
		fmt.Println("写入数据失败")
	}else{
		fmt.Println("写入成功, 写入了", len, "个字节")
	}
	w.Flush()

}
