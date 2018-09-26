package main

func main() {
	/*
	1.在C语言中如何转换
	1.1隐式转换
	int a = 10;
	int b = 3.0
	int num = 3.14;   double res = a/b;
	1.2显示转换
	int a = 10;
	int b = 3.0
	double res = (double)a/b;

	2.Go语言中如何转换
	在Go语言中只有显示转换, 没有隐式转换
	格式: 数据类型(被转换的数据)

	注意点:
	1.数据类型(被转换的数据)格式一般用于除了字符串和布尔类型以外的其它基本数据类型之间转换
	2.不能对一个常量进行强制转换
	3.在Go语言中数据类型必须一模一样, 才能直接赋值
	4.能够直接赋值的特殊情况 byte->uint8, rune->int32
	5.bool类型不能强制转换为整型
	 */

	 // 1.没有隐式转换, 会报错
	 //var num int = 3.14

	 // 2.不能对一个常量进行强制转换
	 //var num int = int(3.14)

	 // 3.正确的做法
	 //var num float64 = 3.14
	 //var value int = int(num)
	 //fmt.Printf("%d\n", value)

	 // 4.注意点:
	 // 4.1在Go语言中数据类型必须一模一样, 才能直接赋值
	// var num int32 = 666
	// //var value int64 = num
	// var value int32 = num
	//fmt.Printf("%d\n", value)

	// 4.2特殊情况
	// byte --> uint8
	//var ch byte = 'a'
	//var num uint8 = ch
	//fmt.Printf("%c\n", num)

	// rune --> int32
	//var ch rune = '李'
	//var num int32 = ch
	//fmt.Printf("%c\n", num)

	// 4.3bool类型不能强制转换为整型
	//var flag bool = false
	//var num int = int(flag)
	//fmt.Printf("%d\n", num)

	// 5.整型也可以通过T(v)转换为字符串类型, 但是在企业开发中不要这么干
	//var num int = 97
	//	//var str string = string(num)
	//	//fmt.Printf("%s\n", str)

}
