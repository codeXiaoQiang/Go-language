package main

import "fmt"

func main() {
	/*
	1.获取字符串长度
	我们已知字符串在Go语言中的本质是一个切片, 如果想要获取切片的长度, 可以通过len()函数来获取
	所以字符串也可以通过len()函数来获取字符串的长度

	2.注意点:
	2.1len()函数获取字符串长度的时候, 获取的是字节数
	在Go语言中中文是按照UTF-8编码的, 所以一个中文占用3个字节
	2.2如果想要获取的是中文的个数, 而不是字节数, 那么需要将字符串转换为rune类型的切片才行
	 */
	 //str := "lnj"
	 //len := len(str)
	 //fmt.Println(len)

	//str := "李南江"
	//len := len(str)
	//fmt.Println(len)

	//str1 := "李南江"
	//var str2 []rune = []rune(str1)
	//len := len(str2)
	//fmt.Println(len)


	str1 := "lnj李南江"
	var str2 []rune = []rune(str1)
	len := len(str2)
	fmt.Println(len)

}
