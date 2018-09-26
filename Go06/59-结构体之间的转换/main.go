package main

import "fmt"

func main() {
	/*
	结构体变量之间可以相互转换, 但是必须保证结构体类型的
	属性名称 属性类型 属性顺序 属性的个数 都一样
	 */

	type Person1 struct {
		name string
		age int
	}
	// 属性顺序不同
	type Person2 struct {
		age int
		name string
	}
	// 如果属性名称和类型都一样, 但是顺序不一样, 不能转换
	//var p1 Person1 = Person1{}
	//var p2 Person2
	//p2 = Person2(p1)
	//fmt.Println(p1)
	//fmt.Println(p2)

	type Person3 struct {
		name1 string
		age int
	}
	// 如果属性的类型和顺序都一样, 但是名称不一样, 不能转换
	//var p1 Person1 = Person1{}
	//var p2 Person3
	//p2 = Person3(p1)
	//fmt.Println(p1)
	//fmt.Println(p2)

	type Person4 struct {
		name string
		age int
		score int
	}
	// 如果属性的名称和类型都一样, 但是个数不一样, 不能转换
	//var p1 Person1 = Person1{}
	//var p2 Person4
	//p2 = Person4(p1)
	//fmt.Println(p1)
	//fmt.Println(p2)

	type Person5 struct {
		name [10]byte
		age int
	}
	// 如果属性名称和个数都一样, 但是属性数据类型不一样, 不能转换
	//var p1 Person1 = Person1{}
	//var p2 Person5
	//p2 = Person5(p1)
	//fmt.Println(p1)
	//fmt.Println(p2)

	type Person6 struct {
		name string
		age int
	}

	// 只有属性个数, 属性名称, 属性类型, 属性属性都一样, 才能转换
	var p1 Person1 = Person1{}
	var p2 Person6
	p2 = Person6(p1)
	fmt.Println(p1)
	fmt.Println(p2)
}
