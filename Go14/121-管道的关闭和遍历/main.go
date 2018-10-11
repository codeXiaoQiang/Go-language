package main

import "fmt"

func main() {
	/*
	管道的遍历:
	可以使用for循环, 也可以使用 for range循环, 以及死循环来遍历
	但是更推荐使用后两者

	因为在企业开发中, 有可能我们不知道管道中具体有多少条数据, 所以如果利用for循环来遍历, 那么无法确定遍历的次数, 并且如果遍历的次数太多, 还会报错
	 */

	/*
	 // 1.创建一个管道
	 myCh := make(chan int, 3)
	 // 2.往管道中写入数据
	 myCh<-1
	 myCh<-2
	 myCh<-3
	 // 注意点: 在企业开发中, 一旦写完了数据, 最好将管道关闭
	 //         只要管道被关闭了, 那么for range循环就不会报错了, 就会自动结束了
	 close(myCh)

	 // 3.遍历管道, 不推荐的
	//for i:=0;i<4 ;i++  {
	//	fmt.Println(<-myCh)
	//}

	// 注意点: 如果被遍历的管道没有关闭, 那么会报错
	//for v := range myCh{
	//	fmt.Println(v)
	//}

	for{
		// 注意点: 如果被遍历的管道没有关闭, 那么会报错
		//         如果管道没有被关闭, 那么会将true返回给ok, 否则会将false返回给Ok
		if v, ok := <-myCh; ok{
			fmt.Println(v)
			fmt.Println(ok)
		}else{
			break
		}
	}
	*/


	// 管道关闭的注意点:
	// 管道关闭之后就不能往管道中写入数据了
	// 但是管道关闭之后还可以从管道中读取数据
	myCh := make(chan int, 3)
	close(myCh)
	//myCh<-1
	//myCh<-2
	//myCh<-3
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
	fmt.Println(<-myCh)
}
