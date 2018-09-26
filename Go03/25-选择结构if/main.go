package main

import "fmt"

func main() {
	/*
	1.在Go语言中没有C语言中的三目运算符, 所有用三目解决的问题都可以使用ifelse来解决, 所以就没有三目运算符

	2.C语言中的if选择结构有三种格式
	if(条件表达式){
	   被控制的语句;
	}

	if(条件表达式){
	   被控制的语句;
	}else{
	   被控制的语句;
	}

	if(条件表达式){
	   被控制的语句;
	}else if(条件表达式){
	   被控制的语句;
	}else{
	   被控制的语句;
	}

	3.Go语言中的if选择结构也有三种格式, 并且和C语言中的格式非常非常非常像
	if 初始化表达式;条件表达式 {
	   被控制的语句;
	}

	if 初始化表达式;条件表达式 {
	   被控制的语句;
	}else{
	   被控制的语句;
	}

	if 初始化表达式;条件表达式{
	   被控制的语句;
	}else if 初始化表达式;条件表达式{
	   被控制的语句;
	}else{
	   被控制的语句;
	}

	4.C语言中if选择结构的特点:
	4.1条件表达式非零即真
	4.2if后面的{}可以省略
	4.3如果省略大括号, 不能定义行的变量

	5.Go语言if选择结构的特点:
	5.1条件表达式必须是返回布尔类型的值, 没有非零即真的概念
	5.2if后面的{}不能省略
	5.3可以在条件表达式前面添加初始化表达式
    5.4每个条件表达式的前面都可以添加初始化表达式
	5.5多个初始化表达式定义了同名的变量, 访问时采用就近原则(退化赋值)
	 */

	 // 第一种格式:
	 /*
	 //age := 17
	 //if age >=18 {
	 //	fmt.Println("开网卡")
	 //}

	// 注意点: 如果将变量定义到if的初始化表达式中, 那么变量只能在if结构中使用
	// 这样做的目的是为了提升性能, 如果变量只在选择结构中使用, 那么建议定义到初始化表达式中, 这样只要选择结构结束了, 那么变量就释放了
	if age := 19; age >=18 {
		fmt.Println("开网卡")
	}
	 fmt.Println("买饮料")
	//fmt.Println(age) // 报错
	 */

	 // 第二种格式
	 /*
	 //age := 17
	 //if age >= 18 {
	 //	fmt.Println("开网卡")
	 //}else{
	 //	fmt.Println("回家叫家长")
	 //}

	// 注意点: 在if初始化表达式中定义的变量, 是在整个if选择结构中都可以使用的,
	// 不仅仅只是在if{}中可以使用, 在else{}中也可以使用
	if age := 17;age >= 18 {
		fmt.Println("开网卡")
	}else{
		fmt.Println("回家叫家长")
		fmt.Println(age)
	}
	 fmt.Println("买饮料")
	 */

	 // 第三种格式:
	// 注意点:
	// 1.不仅仅if后面可以写初始化表达式, 所有的elseif后面也可以写初始化表达式
	// if后面的初始化表达式在整个选择结构中都可以使用
	// elseif后面的初始化后表达式, 只能在当前的elseif和后续的elseif/else中使用, 不能在初始化表达式定义之前使用
	// 2.如果if初始化表达式定义的变量和elseif初始化表达式定义的变量同名了, 那么会出现覆盖现象(退化赋值)
	if age := 3; age >= 22 {
		fmt.Println("该上班了")
		fmt.Println(age)
	}else if num := 666;  age >= 18{
		fmt.Println("该上大学了")
		fmt.Println(age)
		//fmt.Println(num)
	}else if age := 666; age >= 12{ // 因为前面已经定义过了, 所以相当于age = 666
		fmt.Println("该上中学了")
		fmt.Println(age) // 这里访问的是当前elseif定义的age
		fmt.Println(num) // 这里访问的是上一个elseif定义的num
	}else{
		fmt.Println("该干嘛干嘛")
		fmt.Println(age)
		//fmt.Println(num)

	}
	fmt.Println("other")

}
