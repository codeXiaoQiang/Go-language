package main

import (
	"errors"
	"fmt"
)

func main()  {
	/*
	1.什么是异常?
	不按照我们期望执行的都可以称之为异常

	2.在Go语言中如何处理异常?
	一种是程序发生异常时, 将异常信息反馈给使用者
   一种是程序发生异常时, 立刻退出终止程序继续运行

	3.将异常信息反馈给使用者
	创建方式一: fmt.Errorf("提示的内容")
	创建方式二: errors.New("提示的内容")
	注意点: 本质上两个方法底层的实现原理都是一样的

	package builtin中定义了一个接口
	type error interface {
		Error() string
	}

	package errors中定义了一个结构体
	type errorString struct {
		s string
	}
	errorString结构体实现了builtin中定义的接口
	func (e *errorString) Error() string {
		return e.s
	}
	所以errorString结构体实现了error接口

	func New(text string) error {
		return &errorString{text}
	}
	 */
	//if res, err := div(10, 5); err == nil{
	//	fmt.Println(res)
	//}
	if res, err := div(10, 0); err == nil {
		fmt.Println(res)
	}else{
		fmt.Println(err)
	}

}
// 除法运算
func div(a, b int) (res int, err error) {
	if b == 0 {
		// 创建一个异常信息
		//fmt.Println(fmt.Errorf("除数不能为0"))
		err = fmt.Errorf("除数不能为0")
		err = errors.New("除数不能为0")
	}else{
		res = a / b
	}
	return
}