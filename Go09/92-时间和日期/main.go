package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	1.如何获取当前时间?
	package time
	time.Now() // 获取当前时间

	2.如何单独获取当前时间的年月日时分秒
	now.年月日时分秒的英文

	3.如何按照我们指定的格式生成时间字符串

	 */

	 // 1.获取当前时间
	 //now := time.Now()
	 //fmt.Println(now)
	 // 2.单独获取当前时间的年月日时分秒
	//fmt.Println("年",now.Year())
	//fmt.Println("月",now.Month())
	//fmt.Println("日",now.Day())
	//fmt.Println("小时",now.Hour())
	//fmt.Println("分钟",now.Minute())
	//fmt.Println("秒钟",now.Second())

	// 3.按照我们指定的格式生成时间字符串
	/*
	t := time.Now()
	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	str := fmt.Sprintf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(str)
	*/

	/*
	t := time.Now()
	//fmt.Println(t)
	// 2006/01/02 15:04:05 这个字符串是Go语言规定的, 各个数字都是固定的
	// 字符串中的各个数字可以只有组合, 这样就能按照需求返回格式化好的时间
	// 其它语言 : yyyy-MM-dd HH:mm:ss
	//str := t.Format("2006-01-02 15:04:05")
	//str := t.Format("2006-01-02")
	//str := t.Format("15:04:05")
	str := t.Format("2006")
	fmt.Println(str)
	*/


	// 时间常量的使用
	/*
	// 常见的就是配合休眠函数一起使用
	for{
		// 每隔1秒打印一次
		//time.Sleep(time.Second)
		// 每隔0.1秒打印一次
		//time.Sleep(time.Second * 0.1)
		time.Sleep(time.Millisecond * 100)
		fmt.Println("我被打印了")
	}
	*/

	// 时间戳的时间
	// go语言中的时间戳是从1970年1月1日开始计算的
	// Unix: 返回当前时间距离1970年1月1日有多少秒
	// UnixNano: 返回当前时间距离1970年1月1日有多少纳秒
	//fmt.Println(time.Now().Unix()) // 秒
	//fmt.Println(time.Now().UnixNano()) // 纳秒

	for{
		// 设置随机因子(一定要保证每次运行随机因子都不一样才能真正的生成随机数)
		rand.Seed(time.Now().UnixNano())
		// 随机数
		res := rand.Intn(3)
		fmt.Println(res)
	}

}
