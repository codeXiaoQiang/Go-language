package demo

import "fmt"

var value1 = 888
var Value2 = 999

func say1()  {
	fmt.Println("demo say1")
}
func Say2()  {
	fmt.Println("demo say2")
}

type student1 struct {
	name string
}

type Student2 struct {
	// 注意点: 不仅结构体的类型名称首字母要大写, 结构体属性的名称首字母也要大写
	Name string
}

func (Student2)eat1()  {
	fmt.Println("demo eat1")
}
func (Student2)Eat2()  {
	fmt.Println("demo Eat2")
}