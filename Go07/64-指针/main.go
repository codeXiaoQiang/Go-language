package main

import "fmt"

func main() {
	/*
	1.普通指针
	Go语言中的普通指针和C语言中的普通指针一样, 通过指针也可以间接操作指向的存储空间
	C语言中指针的格式  int *p
	Go语言中指针的格式 var p *int

	2.指向数组的指针
	2.1Go语言中指向数组的指针和C语言中指向数组的指针差不多, 通过指针也可以间接操作数组
	C语言中指向数组指针的格式   int[3] *p
	Go语言中指向数组指针的格式 var p *[3]int
	2.2和C语言, 只要一个指针指向数组之后, 就可以通过指针来操作数组
	格式: (*p)[0]    /    p[0]

	注意点: C语言中数组名/&数组名/&数组首元素 都是同一个地址
	        Go语言中&数组名/&数组首元素 都是同一个地址, 但是通过数据名不能直接获取数组的地址
	注意点: 在Go语言中只有相同类型的数据才能赋值
	注意点: 在Go语言中指向数组的指针不支持+1 -1的操作

	3.指向切片的指针
    3.1一定要清楚切片的本质是一个结构体, 结构体中的一个属性是指针, 这个指针指向了底层的一个数组

	注意点:
	如果一个指针指向了数组, 那么可以通过 (*p)[0]和p[0]来操作指向的数组
	但是如果一个指针指向了切片, 只能通过(*p)[0]来操作指向的切片

	4.指向字典的指针
	和普通指针一样, 只能通过(*p)[key]来操作字典

	5.指向结构体的指针
	和C语言中指向结构体的指针一样, 可以通过(*p).attr和p.attr来操作结构体

	6.指针作为函数的参数和返回值
	6.1指针作为函数的参数和C语言一样, 修改形参会影响实参
	6.2指针作为函数的返回值, 如果指针指向的是一个局部变量, 那么不建议返回
	因为当函数调用完毕局部变量就释放了
	 */
	 // 1.普通指针
	 /*
	var num int = 666
	var value float32 = 3.14
	var flag bool = false

	var p1 *int
	p1 = &num
	fmt.Println(num)
	*p1 = 678
	fmt.Println(num)


	var p2 *float32 = &value
	fmt.Println(value)
	*p2 = 6.66
	fmt.Println(value)

	p3 := &flag
	fmt.Println(flag)
	*p3 = true
	fmt.Println(flag)
	 */

	// 2.指向数组的指针
	/*
	var arr [3]int = [3]int{1, 3, 5}
	//fmt.Printf("arr = %p\n", arr)
	//fmt.Printf("&arr = %p\n", &arr)
	//fmt.Printf("&arr[0] = %p\n", &arr[0])

	//fmt.Printf("&arr = %T\n", &arr) // *[3]int
	//fmt.Printf("&arr[0] = %T\n", &arr[0]) // *int

	var p *[3]int
	//fmt.Printf("p = %T\n", p) // *[3]int
	p = &arr
	//p = &arr[0] // 因为数据类型不同, 所以不能赋值
	//fmt.Println(arr)
	//fmt.Println(*p)
	//arr[0] = 666
	//(*p)[0] = 666
	//p[0] = 678
	//*(p + 1) = 678 // 会报错, C语言中支持, 但Go语言中不支持
	fmt.Println(arr)
	*/

	// 3.指向切片的指针
	/*
	var sce []int = []int{1, 3, 5}
	//// 如果直接打印sce, 得到的是结构体中指向底层数组的指针保存的值
	//fmt.Printf("sce = %p\n", sce) // sce = 0xc0420600a0
	//// 如果打印&sce, 得到的是切片结构体自己的地址
	//fmt.Printf("&sce = %p\n", &sce) // &sce = 0xc04205c3e0

	//fmt.Printf("sce = %T\n",sce) // []int
	//fmt.Printf("&sce = %T\n",&sce) // *[]int


	var p *[]int
	//fmt.Printf("p = %T\n", p) // *[]int
	p = &sce
	// 结论: p == &sce / *p == sce
	//fmt.Printf("p = %p\n", p) // p = 0xc04205c3e0
	//fmt.Printf("*p = %p\n", *p) // *p = 0xc0420600a0
	//sce[0] = 666
	(*p)[0] = 666
	// p[0] = 666 // 会报错, 因为p等于&sce, 而操作切片需要sce来操作
	fmt.Println(sce)
	*/

	// 4.指向字典的指针
	/*
	var dict map[string]string = map[string]string{"name":"lnj", "age":"18"}
	var p *map[string]string
	p = &dict
	//dict["name"] = "zs"
	(*p)["name"] = "zs"
	//p["name"] = "zs" // 会报错
	fmt.Println(dict)
	*/

	// 5.指向结构体的指针
	/*
	type Person struct {
		name string
		age int
	}
	var per Person = Person{"lnj", 18}
	var p *Person
	p = &per
	//(*p).name = "zs"
	p.name = "zs"
	fmt.Println(per)
	*/

 	// 6.指针作为函数的参数
 	//num := 123
	//fmt.Println(num)
	//change(&num)
 	//fmt.Println(num)

 	// 7.指针作为函数的返回值
 	var ptr *int = getPtr()
	fmt.Println(*ptr)
}

func getPtr() *int {
	var num int = 123
	var p *int = &num
	return p
}

func change(value *int)  {
	*value = 666
}