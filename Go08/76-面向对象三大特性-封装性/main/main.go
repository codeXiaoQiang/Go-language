package main

import (
	"demo"
	"fmt"
)

//// 1.定义一个类
//type person struct {
//	name string
//	age int
//}

func main()  {
	// 2.通过类创建对象
	p := demo.Person{}
	//p.name = "lnj"
	// 当一个类把自己的成员变量暴露给外部的时候,那么该类就失去对该成员变量的管理权，别人可以任意的修改你的成员变量
	//p.age = -1

	// 封装之后: 提供了数据的安全性, 将变化隔离到了自己当中, 外界不用关心内部的处理, 只需要拿到方法使用即可
	// 封装的原则: 将不需要对外提供的内容都隐藏起来,把属性都隐藏,提供公共的方法对其访问
	p.SetName("lnj")
	p.SetAge(-1)
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}
