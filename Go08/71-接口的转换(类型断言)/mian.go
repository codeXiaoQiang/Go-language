package main

import "fmt"

// 1.定义接口
type USB interface {
	start()
}

// 2.定义一个结构体
type Computer struct {
	name string
}

// 利用结构体实现接口中所有的方法
func (cm Computer) start() {
	fmt.Println(cm.name, "启动了")
}
func (cm Computer) say() {
	fmt.Println(cm.name)
}
func main() {
	/*
	1.如果结构体实现了某个接口, 那么就可以使用接口类型来保存结构体变量
	2.如果利用接口类型变量保存了实现了接口的结构体, 那么该变量只能访问接口中的方法, 不能访问结构体中的特有方法, 以及结构体中的属性
	3.如果利用接口类型变量保存了实现了接口的结构体,想要访问结构体中特有的方法和属性, 那么必须进行类型转换, 将接口类型转换为结构体类型
	 */
	// 1.利用接口类型变量保存实现了接口的数据类型
	//var in USB = Computer{"惠普"}
	// 2.利用接口类型变量, 调用接口中的方法
	//in.start()
	// 3.利用接口类型变量, 调用实现了接口的数据类型中特有的方法
	// 会报错
	//in.say()
	// 4.利用接口类型变量, 访问实现了接口的结构体中的属性
	// 会报错
	//fmt.Println(in.name)

	// 5.将接口类型转换为结构体类型
	/*
	// 方式一: 使用类型断言
	// 格式: cm, ok : = 接口类型变量名.(结构体类型)
	//cm, ok := in.(Computer)
	//fmt.Printf("%T\n", cm)
	//fmt.Println(cm)
	//fmt.Println(ok)
	//cm.say()
	//fmt.Println(cm.name)

	//if cm, ok := in.(Computer); ok{
	//	cm.say()
	//	fmt.Println(cm.name)
	//}
	*/

	// 方式二: 使用类型断言
	/*
	// 会将接口类型变量转换为对应的原始类型之后赋值给cm
	// cm := 接口变量名.(type)
	switch cm := in.(type) {
	case Computer:
		fmt.Println(cm)
		cm.say()
	default:
		fmt.Println("不是Person类型")
	}
	*/
	/*
	cm := Computer{"惠普"}
	var sce []interface{} = []interface{}{1, 3.14, false, "lnj", cm}
	// 需求: 获取切片中保存的每一个数据的原始类型
	for key, value := range sce {
		switch temp := value.(type) {
		case int:
			fmt.Println("第",key, "个元素是int类型")
		case float64:
			fmt.Println("第",key, "个元素是float64类型")
		case bool:
			fmt.Println("第",key, "个元素是bool类型")
		case string:
			fmt.Println("第",key, "个元素是string类型")
		case Computer:
			fmt.Println("第",key, "个元素是Computer类型")
			temp.say()
		}
	}
	*/
}
