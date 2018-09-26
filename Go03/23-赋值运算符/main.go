package main

import "fmt"

func main() {
	/*
	1.C语言中的赋值运算符有
	= += *= /= %= &= |= ^= <<= >>=
	2.Go语言中的赋值运算符和C语言一样
	= += *= /= %= &= |= ^= <<= >>=
	只不过新增了一个   &^=


	 */
	 var num int
	 num = 10
	 num += 5 // 相当于 num = num + 5
	 num -= 5 // 相当于 num = num - 5
	 num *= 5 // 相当于 num = num * 5
	 num /= 5 // 相当于 num = num / 5
	 num %= 5 // 相当于 num = num % 5
	 num &= 5 // 相当于 num = num & 5
	 num |= 5 // 相当于 num = num | 5
	 num <<= 5 // 相当于 num = num << 5
	 num >>= 5 // 相当于 num = num >> 5
	 num &^= 5 // 相当于 num = num &^ 5
	 fmt.Println(num)
}
