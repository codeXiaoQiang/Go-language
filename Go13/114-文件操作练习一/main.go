package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 将一个文本文件拷贝到另外一个文件夹中

	// 1.读取文件
	// 由于文件比较小, 所以可以直接使用ioutil中的Read方法来读取
	buf, err := ioutil.ReadFile("D:/lnj.txt")
	if err != nil{
		fmt.Println("读取文件失败")
		return
	}

	// 2.写入文件
	// 由于文件比较小, 所以可以直接使用ioutil中的Write方法来写入
	err = ioutil.WriteFile("D:/abc.txt", buf, 0666)
	if err != nil{
		fmt.Println("写入文件失败")
	}else{
		fmt.Println("写入文件成功")
	}

}
