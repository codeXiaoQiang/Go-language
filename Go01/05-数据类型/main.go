package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
	1.如何计算数据占用大小
	在C语言中我们可以通过sizeof来计算数据类型占用的内存大小
	在Go语言中我们也可以通过sizeof来计算数据类型占用的内存大小

	2.如何使用sizeof
	2.1导入unsafe
	2.2利用unsafe.Sizeof调用函数

	注意点:
	Go语言中的int类型, 会根据当前系统自动调整, 如果系统是64位的, 那么int就会自动变成int64, 如果系统是32位的, 那么int就会自耦单变成int32
	 */
	 // 1.有符号的整型
	fmt.Println("int size = ", unsafe.Sizeof(int(0)))
	fmt.Println("int8 size = ", unsafe.Sizeof(int8(0)))
	fmt.Println("int16 size = ", unsafe.Sizeof(int16(0)))
	fmt.Println("int32 size = ", unsafe.Sizeof(int32(0)))
	fmt.Println("int64 size = ", unsafe.Sizeof(int64(0)))

	// 2.无符号的整型
	fmt.Println("uint size = ", unsafe.Sizeof(uint(0)))
	fmt.Println("uint8 size = ", unsafe.Sizeof(uint8(0)))
	fmt.Println("uint16 size = ", unsafe.Sizeof(uint16(0)))
	fmt.Println("uint32 size = ", unsafe.Sizeof(uint32(0)))
	fmt.Println("uint64 size = ", unsafe.Sizeof(uint64(0)))
	fmt.Println("uintptr size = ", unsafe.Sizeof(uintptr(0)))

	// 3.字符类型
	// byte相当于C语言中的char, 专门用于保存一个字符  'a'
	// rune本质是int32, 专门用于保存一个中文字符      '李'
	// 注意点: GO语言默认就支持中文,默认就是按照UTF-8来处理 , 在UTF-8中一个中文占3个字节
	// 过去有时候在Windows中一个中文占2个字节, 在Linux中一个中文占3个字节, 非常不方便我们做兼容处理, 所以Go有效的解决了过去的这个问题
	fmt.Println("byte size = ", unsafe.Sizeof(byte(0)))
	fmt.Println("rune size = ", unsafe.Sizeof(rune(0)))

	// 4.浮点 类型
	fmt.Println("float32 size = ", unsafe.Sizeof(float32(0)))
	fmt.Println("float64size = ", unsafe.Sizeof(float64(0)))

	// 5.布尔类型
	fmt.Println("true size = ", unsafe.Sizeof(true))
	fmt.Println("false size = ", unsafe.Sizeof(false))

}
