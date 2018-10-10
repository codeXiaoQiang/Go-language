package main

import (
	"Clib"
	"Game"
	"fmt"
	"time"
)

func init()  {
	// 隐藏光标
	Clib.HideCursor()
}

func main() {
	// 1.创建地图对象
	temp := "############\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"#          #\n"+
		"############\n"
	mp := Game.CreateGameMap(7,12, temp)
	// 2.让地图对象自己打印地图
	mp.PrintMap()

	// 3.创建蛇
	sk := Game.DefaultGameSnake(mp)

	// 4.创建食物
	fd := Game.DefaultGameFood(mp)

	// 5.将蛇绘制到地图中
	mp.DrawSnake(sk)

	// 6.将食物绘制到地图中
	mp.DrawFood(fd)

	// 7.让地图接收用户输入, 以便于控制方向
	// 注意点: go会开启一个协程
	go mp.GetDet(&sk)
	// 6.让蛇移动起来
	mp.MoveSnake(&sk, &fd)

	// 7.游戏结束
	// 7.1将关闭移动到最后
	Clib.GotoPostion(0, mp.GetRow())
	fmt.Println("游戏结束")
	time.Sleep(time.Minute)
}
