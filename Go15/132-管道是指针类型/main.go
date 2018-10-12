package main

import "fmt"

func main() {
	var myCh1 chan int = make(chan int, 5)
	//fmt.Println(myCh1) // 0xc042094000
	//fmt.Printf("%p\n", myCh1) // 0xc042094000
	//fmt.Printf("%p\n", &myCh1) // 0xc042082018
	myCh1<-1
	myCh1<-2

	var myCh2 <-chan int
	myCh2 = myCh1 // 将双向的管道转换成单向的管道
	// 打印单向管道的长度和容量
	fmt.Println("len", len(myCh2), "cap", cap(myCh2))
	fmt.Println(<-myCh2)
	// 打印双向管道的长度和容量
	fmt.Println("len", len(myCh1), "cap", cap(myCh1))

}
