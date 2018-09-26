package main

func main() {
	/*
	1.C语言中的switch结构
	switch(值){
	case 值:
	   被控制语句;
	  break;
	case 值:
	   被控制语句;
	  break;
	default:
	  被控制语句;
	  break;
	}

	2.Go语言中的Switch结构和C语言非常非常非常像
	switch 初始化表达式; 值{
		case 值:
		   被控制语句;
		case 值:
		   被控制语句;
		default:
		  被控制语句;
	}

	注意点:
	1.和C语言不一样, Go语言的switch可以传递小数
	2.和C语言不一样, Go语言switch中还可以添加初始化表达式
	3.和C语言不一样, Go语言不用编写break语句, 默认就不会穿透
	4.在Go语言中, 如果想要穿透, 必须手动添加一个fallthrough
	5.Go语言中的switch可以当做if来使用
	6.default可以不写, 也可以写到其它位置, 无论写到哪都会最后判断执行
	7.case后面定义变量不用添加大括号{}
	 */

	//age := 22
	//switch age {
	//case 22:
	//	fmt.Println("该上班了")
	//case 18:
	//	fmt.Println("该上大学了")
	//default:
	//	fmt.Println("该干嘛干嘛")
	//}

	//switch age := 3; age {
	//case 22:
	//	fmt.Println("该上班了")
	//case 18:
	//	fmt.Println("该上大学了")
	//default:
	//	fmt.Println("该干嘛干嘛")
	//}

	//switch num := 18.0; num { // 传递一个小数
	//case 22:
	//	fmt.Println("该上班了")
	//case 18:
	//	fmt.Println("该上大学了")
	//	fallthrough // 告诉系统需要穿透
	//default:
	//	fmt.Println("该干嘛干嘛")
	//}

	// 可以当做if结构来使用
	//switch num := 50;{
	//case num >=0 && num <=50:
	//	value := 666
	//	fmt.Println("0~50之间的数")
	//	fmt.Println(value)
	//case num >50 && num <=100:
	//	fmt.Println("51~100之间的数")
	//default:
	//	fmt.Println("其它区间")
	//}
}
