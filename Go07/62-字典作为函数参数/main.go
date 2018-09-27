package main

import "fmt"

func main() {
	/*
	1.字典作为函数的参数
	字典之间的赋值和切片之间的赋值一样, 是地址传递
	所以字典作为函数的参数, 修改形参会影响到实参

	2.注意点:
	字典和数组以及切片不一样, 数组和切片保存的数据是有序的, 而字典保存的数据是无序的
	 */
	var dict1 map[string]string = map[string]string{"name": "lnj", "age": "18"}
	//var dict2 map[string]string
	//dict2 = dict1
	fmt.Println(dict1)
	//dict2["name"] = "zs"
	change(dict1)
	fmt.Println(dict1)
}

func change(dict2 map[string]string)  {
	dict2["name"] = "zs"
}
