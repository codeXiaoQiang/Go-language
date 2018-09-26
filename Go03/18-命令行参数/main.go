package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
	1.C语言中的命令行参数
	在C语言中main函数可以接收两个参数
	int main(int argc, const char *argv[]){}

	2.Go语言中的命令行参数
	在Go语言中main函数是不接收任何参数的, 但是还是可以通过其它方式来接收命令行参数

	3.Go语言中接收命令行参数方式一:
	3.1导入os包,
	3.2然后通过os包的Args属性获取命令行参数
	注意点:
	和C语言一样, 默认会将当前被执行程序的可执行文件的路径传递给我们
	和C语言一样, 无论外界传入的数据是什么类型, OS包拿到的都是字符串类型

	4.通过os包获取命令行参数的弊端:
	4.1如果用户少传入了参数, 会报错
	4.2无论用户传入的是什么类型, 最终获取到的都是字符串类型
	4.3传递参数的时候, 必须按照业务逻辑的顺序来传递, 传递的顺序不能被打乱

	5.Go语言中接收命令行参数方式二:
	5.1导入flag包
	5.2通过flag包中的stringVar函数获取

	6.通过flag包获取命令行参数的优势:
	6.1如果用户少传入了参数, 不会报错
	6.2传递参数的时候, 不用按照业务逻辑的顺序来传递, 传递的顺序可以被打乱
	 */
	 /*
	 //fmt.Println(os.Args[0])
	 //fmt.Println(os.Args[1])
	 //fmt.Println(os.Args[2])
	 //fmt.Printf("%T\n", os.Args[0])
	 //fmt.Printf("%T\n", os.Args[1])
	 //fmt.Printf("%T\n", os.Args[2])

	 //var exePath = os.Args[0]
	 //var name = os.Args[1]
	 //var age = os.Args[2]
	 //fmt.Println("path = ", exePath, "name = ", name, "age = ", age)
	*/

	/*
	 var name string
	 // 第一个参数: 要把接收到的数据保存到哪
	 // 第二个参数: 用于接收的字段名称   -name=lnj
	 // 第三个参数: 如果用户没有传递参数时的默认值
	 // 第四个参数: 用户输入--help的时候的提示信息
	 flag.StringVar(&name, "name", "默认的名称", "接收用户的姓名")
	 var age string
	 flag.StringVar(&age, "age", "-1", "接收用户的年龄")

	 // 注意点: 通过flag包接收用户输入的命令行参数, 还必须将编写好的参数注册到命令行
	 flag.Parse()

	 // 获取通过命令行参数接收到的数据
	 fmt.Println("name = ", name)
	 fmt.Printf("%T\n", name)
	 fmt.Println("age = ", age)
	fmt.Printf("%T\n", age)
	*/

	/*
	// 1.注册命令行参数
	var name *string = flag.String("name", "默认名称", "接收用户名称")
	var age *string = flag.String("age", "-1", "接收用户年龄")

	// 2.注册命令行参数
	flag.Parse()

	// 3.获取接收到的数据
	fmt.Println(*name)
	fmt.Println(*age)
	*/

	var name string
    flag.StringVar(&name, "name", "默认姓名", "接收姓名")
	var age int
	flag.IntVar(&age, "age", -1, "接收年龄")
	var height float64
	flag.Float64Var(&height, "height", 0.0, "接收身高")

	flag.Parse()

	fmt.Println(name)
	fmt.Printf("%T\n", name)
	fmt.Println(age)
	fmt.Printf("%T\n", age)
	fmt.Println(height)
	fmt.Printf("%T\n", height)

}
