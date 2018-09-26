package main

func main() {
	/*
	1.如果将字符串类型转换为基本数据类型
	在Go语言中有两种方式可以将基本数据类型转换为字符串类型
	1.1 strconv.ParseXxx()
	1.2 strconv.Atoi()

	Atoi在C语言中也有
	 */

	 // 1.strconv.ParseXxx()
	 /*
	 //var str string = "9"
	 var str string = "1001"
	 // 第一个参数: 需要转换的字符串
	 // 第二个参数: 被转换的字符串保存的整数是多少进制的
	 // 第三个参数: 期望转换为多少位的整数(不一定准确), 换句话说就是要转换为多少位整数
	 // 注意点: 如果被转换的字符串超出了指定的长度会报错
	 // 返回值:
	 // 返回值的第一个: 转换之后的数值, 是int64类型
	 // 返回值的第二个: 如果转换成功返回nil, 如果转换失败就不是nil
	 // int8   -128~127
	 //num, err := strconv.ParseInt(str, 10, 8)
	 num, err := strconv.ParseInt(str, 2, 8)
	 if err != nil{
	 	fmt.Printf("转换失败\n")
	 }else{
		fmt.Printf("%d\n", num)
		fmt.Printf("%T\n", num)
	 }
	*/

	 /*
	var str string = "3.1234567890123456789"
	// 第一个参数: 需要转换的字符串
	// 第二个参数: 要将字符串中的小数转换为单精度还是双精度, 单精度传入32, 双精度传入64
	// 返回值:
	// 第一个返回值: 转换之后的小数, float64类型
	// 第二个返回值: 转换成功返回nil, 转换失败不为nil
	//num, err := strconv.ParseFloat(str, 32)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil{
		fmt.Printf("转换失败\n")
	}else{
		//fmt.Printf("%f\n", num)
		fmt.Println(num)
		fmt.Printf("%T\n", num)
	}
  */

	 /*
	 var str string = "false"
	 flag, err := strconv.ParseBool(str)
	if err != nil{
		fmt.Printf("转换失败\n")
	}else{
		// 注意点: 在Go语言中%t输出布尔类型
		fmt.Printf("%t\n", flag)
		fmt.Printf("%T\n", flag)
	}
	 */

	 // 2.strconv.Atoi()
	 /*
	 //var str string = "9"
	 var str string = "9a"
	 num, err := strconv.Atoi(str)
	 if err != nil{
		 fmt.Printf("转换失败\n")
	 }else {
		 fmt.Printf("%d\n", num)
	 }
	 */
}
