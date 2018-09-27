package main

/*
//type USB interface {
//	start() // 方法的声明
//	// 下面的代码会报错
//	func end(){ // 方法的实现
//		fmt.Println("end")
//	}
//}
*/
/*
//type USB interface {
//	start() // 方法的声明
//	// 下面的代码会报错
//	name string // 声明一个变量
//}
*/
/*
// 1.定义一个接口, 声明两个方法
type USB interface {
	start()
	end()
}

type Phone struct {

}
// 以下代码并没有实现接口中的方法
//func (Phone)start(name string)  {
//	fmt.Println("启动")
//}
func (Phone)start()  {
	fmt.Println("启动")
}
func (Phone)end()  {
	fmt.Println("启动")
}
*/

/*
// 子集接口
type aer interface {
	//aer // 会报错
	start()
}
// 超集接口
type ber interface {
	aer // 在ber接口中嵌套了aer接口
	end()
}
type Phone struct {
	name string
}
// Phone这个结构体实现了aer接口的所有方法
// Phone此刻只实现了ber接口的一个方法
//func (p Phone)start()  {
//	fmt.Println(p.name, "启动了")
//}
//  Phone这个结构体实现了aer接口的所有方法
//  Phone这个结构体实现了ber接口的所有方法
func (p Phone)start()  {
	fmt.Println(p.name, "启动了")
}
func (p Phone)end()  {
	fmt.Println(p.name, "停止了")
}
*/

type USB interface {
	start()
	//start() // 会报错
}
type CMD interface {
	USB
	//start() // 会报错
}

func main() {
	/*
	注意点:
	1.接口中只能有方法的声明, 不能有方法的实现
	2.接口中只能有方法的声明, 不能有变量的声明
	3.只有实现了接口中声明的所有方法, 才算实现了接口, 才能使用接口变量保存
	4.在实现接口的时候, 方法名称,形参列表,返回值列表必须一模一样
	5.接口和结构体一样, 可以嵌套
	6.接口和结构体一样, 嵌套时不能嵌套自己(自己搞自己)
	7.可以将集接口变量赋值给子集接口变量,不可以将子集接口变量赋值给超集接口变量(无论实际的数据类型是否已经实现了超集的所有方法)
	8.接口中不能出现同名的方法声明
	 */

	 //var in USB = Phone{}
	 //fmt.Println(in)

	 /*
	 //var in aer = Phone{"华为"}
	 //var in ber = Phone{"华为"}
	 //in.start()

	// // 1.用超集接口变量保存结构体
	//var b ber = Phone{"华为"}
	//// 2.将超集接口变量赋值给子集接口变量
	//var a aer
	//a = b
	//a.start()

	// 1.用子集接口变量保存结构体
	var a aer = Phone{"华为"}
	// 2.将子集接口变量赋值给超集接口变量
	var b ber
	b = a // 会报错
	b.start()
	 */


}
