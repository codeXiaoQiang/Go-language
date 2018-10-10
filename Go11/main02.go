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
	/*
	mp := Game.GameMap{}
	mp.SetRow(7)
	mp.SetCol(12)
	temp := [][]byte{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	mp.SetDetail(temp)
	*/
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
	/*
	// 3.1初始化蛇头的位置
	headPo := Game.Position{mp.GetCol() / 2, mp.GetRow() / 2}
	// 3.2初始化蛇身体的位置
	bodyPo := Game.Position{mp.GetCol() / 2 - 1, mp.GetRow() / 2}
	// 3.3将所有蛇的位置放入切片
	pos := []Game.Position{headPo, bodyPo}
	sk := Game.CreateGameSnake(pos, len(pos), "d")
	*/
	sk := Game.DefaultGameSnake(mp)

	// 4.创建食物
	/*
	// 4.1设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 4.2生成随机数
	fdRow := rand.Intn(mp.GetRow() - 2) + 1
	fdCol := rand.Intn(mp.GetCol() - 2) + 1
	fd := Game.CreateGameFood(fdCol, fdRow)
	*/
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
