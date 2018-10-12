package main

func main() {
	// 1.创建一个有缓冲的管道
	/*
	myCh := make(chan int, 3)
	myCh<-1
	myCh<-2
	myCh<-3
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	*/

	// 2.没有缓冲的管道
	/*
	myCh := make(chan int, 0)
	// 注意点:
	// 没有缓冲的管道不能直接存储数据
	// 没有缓冲的管道不能直接获取数据
	// 注意点:
	// 想使用没有缓冲的管道, 必须保证读和写同时存在, 而且还必须保证读和写是在不同的go程中
	// 并且读必须写在写的前面
	go func() {
		fmt.Println("读之前的代码")
		fmt.Println(<-myCh)
		fmt.Println("读之后的代码")
	}()
	fmt.Println("写之前的代码")
	time.Sleep(time.Second * 5)
	//myCh<-1
	fmt.Println("写之后的代码")
	//fmt.Println(<-myCh)
	//go func() {
	//	fmt.Println(<-myCh)
	//}()
	*/

	// 2.没有缓冲的管道
	/*
	// 如果是在go程中使用无缓冲的管道, 那么就可以单个使用(只有读, 或者只有写)
	myCh := make(chan int, 0)
	go func() {
		fmt.Println("123")
		myCh<-998
		//<-myCh
		fmt.Println("abc")
	}()
	*/

	// 管道总结:
	// 管道一般都在go程中使用, 不会直接在主线程中使用, 无论是有缓冲的还是没有缓冲的
	// 只要是在go程中使用, 无论是有缓冲的还是没有缓冲的, 都会出现阻塞现象
	for{
		;
	}
}
