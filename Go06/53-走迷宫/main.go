package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 定义全局变量
var(
	// 定义变量保存小人当前的位置
	currentRow int = 1
	currentCol int = 1
	// 定义变量保存迷宫出口的位置(索引)
	endRow int = 1
	endCol int = 5
)

func main() {
	// 1.定义二维数组保存迷宫地图
	sce := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'R', ' ', '#', ' ', ' '},
		{'#', ' ', '#', '#', ' ', '#'},
		{'#', ' ', ' ', '#', ' ', '#'},
		{'#', '#', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	// 2.定义一个函数打印地图
	printMap(sce)
	// 3.定义一个函数用于接收用户输入的方向
	for{
		// 2.接收用户输入的数据
		ch := input()
		// 3.定义函数让小人根据用户输入行走
		move(sce, ch)
		// 4.判断是否已经走出迷宫
		if currentRow == endRow && currentCol == endCol{
			break
		}
		// 4.再次打印地图
		printMap(sce)
	}
	fmt.Println("恭喜你, 已经通关了")
}
func move(m [][]byte, ch string)  {
	switch ch {
	case "w", "W":
		//fmt.Println("向上走")
		if m[currentRow - 1][currentCol] != '#' {
			m[currentRow][currentCol] = ' '
			m[currentRow - 1][currentCol] = 'R'
			currentRow--
		}
	case "s", "S":
		//fmt.Println("向下走")
		if m[currentRow + 1][currentCol] != '#' {
			m[currentRow][currentCol] = ' '
			m[currentRow + 1][currentCol] = 'R'
			currentRow++
		}
	case "a", "A":
		//fmt.Println("向左走")
		if m[currentRow][currentCol - 1] != '#' {
			m[currentRow][currentCol] = ' '
			m[currentRow][currentCol - 1] = 'R'
			currentCol--
		}
	case "d", "D":
		//fmt.Println("向右走")
		if m[currentRow][currentCol + 1] != '#' {
			m[currentRow][currentCol] = ' '
			m[currentRow][currentCol + 1] = 'R'
			currentCol++
		}
	default:
		fmt.Println("先实现自定义方向移动, 请购买付费版本")
	}
}
//func input() byte  {
func input() (ch string)  {
	/*
	//var ch byte
	//fmt.Scanf("%c", &ch)
	//return ch
	*/
	/*
	//var ch string
	//fmt.Scanln(&ch)
	// return ch
	 */
	// 1.提示用户如何输入
	fmt.Println("请输入w a s d, 以回车结束")
	fmt.Println("w代表向上走, a代表向左走, s代表向下走, d代表向右走")

	// 1.接收用户输入的数据
	// 如果是获取字符串, 那么缓冲区中所有的内容都会被取出来包括\n
	fmt.Scanln(&ch)
	// 2.将接收到的数据返回给调用者
	return
}

func printMap(m [][]byte)  {
	/*
	// 1.计算切片的行数
	//row := len(m)
	// 2.计算切片的列数
	//col := len(m[0])
	// 3.通过循环打印地图
	//for i:=0; i<row; i++{
	//	for j:=0; j<col; j++ {
	//		fmt.Printf("%c", m[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	*/

	// 1.清空控制台
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// 2.打印地图
	for _, v1 := range m{
		for _, v2 := range v1{
			fmt.Printf("%c", v2)
		}
		fmt.Printf("\n")
	}
}
