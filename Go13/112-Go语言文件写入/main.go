package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
	第三种写入的方式:
	通过ioutil包中的Write和WirteString函数读取
	 */
	buf := []byte{'l','n','j','\r','\n'}
	err := ioutil.WriteFile("D:/lnj.txt", buf, 0666)
	if err != nil{
		fmt.Println("写入数据失败")
	}else{
		fmt.Println("写入数据成功")
	}
}
