package main

import "fmt"

/*
1.如何快速新建一个文件或者文件夹
Ctrl + Alt + Ins
 */

 /*
 1.保存源代码的文件不同
 1.1C语言的源代码保存在.c文件中
 1.2Go语言的源代码保存在.go文件中
  */
  /*
  2.代码的管理方式不同
  2.1C语言通过文件来管理代码, 会将不同的功能(模块)的代码, 放到不同的文件中, 然后声明.h, 然后include导入.h文件使用对应的代码
  2.2Go语言通过包来管理代码, 会将不同功能(模块)的代码, 放到不同的包中, 然后通过import导入包来使用对应的代码

  什么是包?
  1.在Go语言中, 一个文件夹就是一个包, 只需要将不同的类型的.go文件放到不同的文件夹中, 然后通过package声明包名即可.
  2.注意点:
  2.1Go语言规范规定, 包名要和.go文件所在的文件夹的名称保持一致
  2.2Go语言规定, 报名不能重复

  Go语言标准的项目格式
  项目文件夹(GOPATH对应的文件夹)
  -------|---------src文件夹(用于存储.go文件的文件夹)
  ----------------------|---------main文件夹(存放package main对应文件的文件夹)
  ----------------------|---------other文件夹(存放package 其它对应文件的文件夹)
  -------|---------bin文件夹(用于存放可执行程序的文件夹)
  -------|---------pkg文件夹(用于存放二进制文件的文件夹)
   */

   /*
   3.main函数的书写格式同
   int main(){}
   fun main(){}
    */

    /*
    4.函数的调用格式不同
    4.1C语言只要include之后, 就可以直接通过函数名称调用了
    4.2Go语言通过import之后, 还需通过报名.函数名称的方式来调用
     */
     /*
     5.语句的结束方式不同
     5.1C语言每条语句后面都必须添加分号
     5.2Go语言每条语句后面不用添加分号, 但是如果两条语句在同一行, 就必须添加分号
      */

      /*
      6.对代码的严格程度不同
      6.1C语言中include了一个.h文件, 哪怕没有使用也不会报错
      6.2Go语言中import了一个包, 如果没有使用, 就会报错

      6.3C语言中定义变量没有使用不会报错
      6.4Go语言中定义变量没有使用就会报错

      6.5C语言中函数的{可以和函数名称在同一行, 也可以不再同一行
      6.6Go语言中函数的{必须和函数名称在同一行, 否则会报错
       */
func main() {
	fmt.Printf("Hello LNJ\n")
	fmt.Printf("Hello it666\n")
}
