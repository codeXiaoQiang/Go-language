package main

import "fmt"

func main() {
	/*
	1.增加
	字典名称[key] = 值
	如果字典中没有对应的key, 那么就是增加

	2.修改
	字典名称[key] = 值
	如果字典中已经有对应的key, 那么就是修改

	3.查询
	value, ok = 字典名称[key]
	如果有对应的key, 那么就会把key对应的值赋值给value, 并且ok等于true
	如果没有对应的key, 那么value就等于零值, 并且ok等于false

	4.删除
	delete(字典变量, key)
	会将指定字典中指定key的值删除

	注意点:
	字典中不能有同名的key
	 */
	 /*
	 dict := map[string]string{}
	 fmt.Println(dict)
	 // 增加
	 dict["name"] = "lnj"
	fmt.Println(dict)
	 // 修改
	 dict["name"] = "zs"
	 fmt.Println(dict)
	 */

	 /*
	 // 字典中不能有同名的key
	//dict := map[string]string{"name":"lnj", "name":"zs"}
	//fmt.Println(dict)

	dict := map[string]string{"name":"lnj", "age":"18"}

	//value, ok := dict["name"]
	//fmt.Println(ok) // true
	//fmt.Println(value) // lnj
	//
	//value, ok := dict["gender"]
	//fmt.Println(ok) // false
	//fmt.Println(value) // ""
	 // 查询
	// 企业开发判断字典中有没有某个key的标准写法
	if value, ok := dict["name"]; ok{
		fmt.Println(ok)
		fmt.Println(value)
	}
	 */

	 // 删除
	dict := map[string]string{"name":"lnj", "age":"18"}
	fmt.Println(dict)
	delete(dict, "age")
	fmt.Println(dict)
}
