package main

import (
	"fmt"
	"os"
)

func main() {

	finfo, err := os.Stat("D:/WWW")
	if err != nil {
		fmt.Println(err)
	}else{
		//fmt.Println(finfo)
		fmt.Println(finfo.Name())
		fmt.Println(finfo.Size())
		fmt.Println(finfo.ModTime())
		fmt.Println(finfo.IsDir())
	}
}
