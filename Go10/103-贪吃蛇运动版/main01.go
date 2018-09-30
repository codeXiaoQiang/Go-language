package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// 定义常量保存地图的行数和列数
const (
	MaxRow = 7
	MaxCol = 12
)

// 定义结构体保存蛇的偏移位
type Offset struct {
	ox int
	oy int
}

// 定义一个结构体, 用于保存位置
type Position struct {
	x int // 多少列
	y int // 多行行
}

// 定义一个结构体保存蛇相关的信息
type Snake struct {
	pos []Position // 用于保存蛇所有位置
	length int // 用于保存蛇长度
	direction string // 保存蛇当前运动的方向
}

// 定义一个结构体保存食物相关的信息
type Food struct {
	Position
}

func main() {
	// 1.定义一个二维数组, 保存地图
	snakeMap := [MaxRow][MaxCol]byte{
		{'#','#','#','#','#','#','#','#','#','#','#','#'},
		{'#',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ','#'},
		{'#',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ','#'},
		{'#',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ','#'},
		{'#',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ','#'},
		{'#',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ','#'},
		{'#','#','#','#','#','#','#','#','#','#','#','#'},
	}
	// 2.创建一条蛇
	sk := Snake{}
	sk.length = 2
	sk.pos = make([]Position, 2)
	// 设置蛇头的位置
	sk.pos[0].x = MaxCol / 2
	sk.pos[0].y = MaxRow / 2
	// 设置蛇身体的位置
	sk.pos[1].x = MaxCol / 2 - 1
	sk.pos[1].y = MaxRow / 2

	// 3.将蛇放入地图中
	snakeMap[sk.pos[0].y][sk.pos[0].x] = '@'
	snakeMap[sk.pos[1].y][sk.pos[1].x] = '*'

	// 4.创建食物
	fd := Food{}
	// 4.1设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 4.2设置食物的位置
	fd.x = rand.Intn(MaxCol - 2) + 1
	fd.y = rand.Intn(MaxRow - 2) + 1

	// 5.将食物放到地图中
	snakeMap[fd.y][fd.x] = 'O'

	// 6.打印地图
	printMap(snakeMap)

	// 7.接收用户的输入
	for{
		ch := getDirection()
		// 8.移动蛇
		move(&snakeMap, &sk, ch)

		// 9.进行逻辑判断(碰撞到墙壁, 碰撞到身体, 碰撞到食物)
		res := controller(&snakeMap, &sk, &fd)

		// 10.打印地图
		printMap(snakeMap)
		if res {
			break
		}
	}
	// 10.提示游戏结束
	fmt.Println("游戏结束")
}
// 蛇的逻辑控制相关的方法
func controller(m *[MaxRow][MaxCol]byte, sk *Snake, fd *Food) bool {
	// 0.获取蛇头的位置
	headPos := sk.pos[0]
	// 1.判断蛇有没有碰到墙壁
	if headPos.x <= 0 || headPos.x >= MaxCol-1 || headPos.y <=0 || headPos.y >= MaxRow-1{
		return true
	}
	// 2.判断有没有碰到食物
	if headPos.x == fd.x && headPos.y == fd.y{
		// 2.1获取蛇当前最后的一节身体
		lastPos := sk.pos[sk.length - 1]
		newPos := Position{}

		// 2.1根据蛇运动的方向创建一节行的蛇身体
		switch sk.direction {
		case "w":
			newPos.x = lastPos.x
			newPos.y = lastPos.y + 1
		case "s":
			newPos.x = lastPos.x
			newPos.y = lastPos.y - 1
		case "a":
			newPos.x = lastPos.x + 1
			newPos.y = lastPos.y
		case "d":
			newPos.x = lastPos.x - 1
			newPos.y = lastPos.y
		}
		// 2.2将新创建好的蛇身条件到sk当中
		sk.pos = append(sk.pos, newPos)
		// 2.3让蛇的身体长度+1
		sk.length += 1
		// 2.4将新创建的蛇绘制出来
		m[newPos.y][newPos.x] = '*'

		// 2.5重新随机生成食物的位置
		fd.x = rand.Intn(MaxCol - 2) + 1
		fd.y = rand.Intn(MaxRow - 2) + 1
		// 5.将食物放到地图中
		m[fd.y][fd.x] = 'O'

		return false
	}

	// 3.判断有没有碰到身体
	for i:=1; i<sk.length; i++{
		if sk.pos[0].x == sk.pos[i].x && sk.pos[0].y == sk.pos[i].y{
			return true
		}
	}
	return false
}
// 让蛇移动
func move(m *[MaxRow][MaxCol]byte, sk *Snake,ch string)  {
	// 1.创建一个偏移位
	offset := Offset{0,0}
	// 2.根据用户输入的方向, 获取偏移位
	switch ch {
	case "w", "W":
		offset.ox = 0
		offset.oy = -1
		sk.direction = "w"
	case "s", "S":
		offset.ox = 0
		offset.oy = 1
		sk.direction = "s"
	case "a", "A":
		offset.ox = -1
		offset.oy = 0
		sk.direction = "a"
	case "d", "D":
		offset.ox = 1
		offset.oy = 0
		sk.direction = "d"
	}
	// 3.从地图中移除最后一节蛇身体
	lastPos := sk.pos[sk.length - 1]
	m[lastPos.y][lastPos.x] = ' '

	// 3.修改蛇在地图中的位置
	// 3.1修改蛇身体在地图中的位置
	//  i:= 1
	for i:=sk.length-1; i>0; i--{
		// 让后一个蛇身体移动到前一个蛇身体的位置
		// sk.pos[1] = sk.pos[0]
		sk.pos[i] = sk.pos[i-1]
		m[sk.pos[i].y][sk.pos[i].x] = '*'
	}
	// 3.2修改蛇头在地图中的位置
	sk.pos[0].x += offset.ox
	sk.pos[0].y += offset.oy
	m[sk.pos[0].y][sk.pos[0].x] = '@'
}
// 接收用户输入
func getDirection() (det string) {
	// W上 S下 A左 D右
	fmt.Println("请输入W / S / A / D")
	fmt.Scanln(&det)
	return
}
// 打印地图
func printMap(m [MaxRow][MaxCol]byte)  {
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
