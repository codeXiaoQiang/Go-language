package main

func main() {
	/*
	1.Go语言中的结构体几乎和C语言中的结构体一模一样
	都需要先定义结构体类型, 再定义结构体变量
	都是用来保存一组不同类型数据的

	2.C语言定义结构体类型的格式
	struct 结构体类型名称{
	    属性名称 属性类型;
	    属性名称 属性类型;
	}

	3.Go语言定义结构体类型的格式
	typd 结构体类型名称 struct{
	    属性名称 属性类型;
	    属性名称 属性类型;
	}

	4.C语言中通过结构体类型定义结构体变量, 必须拷贝struct
	struct 结构体类型名称 结构体变量名称;

	5.Go语言中通过结构体类型定义结构体变量, 不用拷贝struct
	var 结构体变量名称 结构体类型名称

	6.和C语言中的结构体一样, 可以定义结构体变量的同时初始化, 也可以先定义结构体变量再初始化

	注意点:
	和切片以及字典不一样, 结构体变量定义之后就可以直接使用了
	 */

	 /*
	 // 1.定义一个结构体类型
	type Person struct{
		name string
		age int
	}
	// 2.通过结构体类型定义结构变量
	var per Person
	per.name = "lnj"
	per.age = 18
	fmt.Println(per)
	 */
	 /*
	type Person struct{
		name string
		age int
	}
	// 定义的同时初始化
	//var per Person = Person{"lnj", 18}
	//fmt.Println(per)

	// 先定义再初始化
	var per Person
	//per = Person{"lnj", 18} // 完全初始化
	//per = Person{"lnj"} // 部分初始化, 必须通过属性名称指定要给谁初始化
	per = Person{name:"lnj"} // 部分初始化, 必须通过属性名称指定要给谁初始化
	fmt.Println(per)
	 */
}
