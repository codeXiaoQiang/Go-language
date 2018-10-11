package main

import "fmt"

func main() {
	/*
	1.什么是管道:
	管道就是一个队列, 具备先进先出的原则
	是线程安全的, 也就是自带锁定功能

	2.管道作用:
	在Go语言的协程中, 一般都使用管道来保证多个协程的同步, 或者多个协程之间的通讯

	3.如何声明一个管道, 和如何创建一个管道
	管道在Go语言中和切片/字典一样也是一种数据类型
	管道和切片/字典非常相似, 都可以用来存储数据, 都需要make之后才能使用
	3.1管道声明格式:
	var 变量名称 chan 数据类型
	var myCh chan int
	如上代码的含义: 声明一个名称叫做myCh的管道变量, 管道中可以存储int类型的数据
	3.2管道的创建:
	make(chan 数据类型, 容量)
	myCh = make(chan int, 3);
	路上代码的含义: 创建一个容量为3, 并且可以保存int类型数据的管道

	4.管道的使用
	4.1如何往管道中存储(写入)数据?
	myCh<-被写入的数据
	4.2如何从管道中获取(读取)数据?
	<-myCh
	对管道的操作是IO操作
	例如: 过去的往文件中写入或者读取数据, 也是IO操作
	例如: 过去的往屏幕上输出内容, 或者从屏幕获取内容, 也是IO操作
	stdin / stdout / stderr

	注意点:
	和切片不同, 在切片中make函数的第二个参数表示的切片的长度(已经存储了多少个数据),
	而第三个参数才是指定切片的容量
	但是在管道中, make函数的第二个参数就是指定管道的容量, 默认长度就是0
	 */

	 // 1.声明管道
	 var myCh chan int
	 // 会报错, 没有make的管道不能直接使用
	 //myCh<-666
	 //<-myCh
	 // 2.创建一个管道
	 myCh = make(chan int, 3)
	 // len = 0, cap = 3
	fmt.Println("len", len(myCh), "cap", cap(myCh))
	 // 3.使用管道
	 // 3.1往管道中写入数据
	//myCh<-1
	// 只要往管道中写入了数据, 那么len就会增加
	//fmt.Println("len", len(myCh), "cap", cap(myCh))
	//myCh<-2
	//myCh<-3
	//fmt.Println("len", len(myCh), "cap", cap(myCh))
	// 注意点: 如果len等于cap, 那么就不能往管道中再写入数据了, 否则会报错
	//myCh<-4

	// 注意点: 如果管道中没有数据, 去管道中读取数据也会报错
	//fmt.Println(<-myCh)
	// 只要从管道中读取了数据, 那么len就会减少
	//fmt.Println("len", len(myCh), "cap", cap(myCh))
	//fmt.Println(<-myCh)
	//fmt.Println(<-myCh)
}
