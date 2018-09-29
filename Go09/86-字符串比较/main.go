package main

func main() {
	/*
	字符串比较
	func Compare(a, b string) int
	func EqualFold(s, t string) bool
	 */

	 // Compare
	 /*
	 //str1 := "123"
	 //str2 := "126"
	str1 := "李南江"
	str2 := "李江南"
	 // Compare作用:
	 // 比较两个字符串
	 // 如果两个字符串相等, 那么返回0
	 // 如果两个字符串不相等
	 // 第一个字符串 > 第二个字符串 返回 1
	 // 第一个字符串 < 第二个字符串 返回 -1
	 res := strings.Compare(str1, str2)
	 fmt.Println(res)

	bytes1 := []byte(str1)
	fmt.Println(bytes1)
	bytes2 := []byte(str2)
	fmt.Println(bytes2)
	*/

	 // EqualFold
	 /*
	 // EqualFold作用:
	 // 比较两个字符串是否相等, 相等返回true, 不相等返回false
	str1 := "123"
	str2 := "124"
	res := strings.EqualFold(str1, str2)
	fmt.Println(res)
	 */

}
