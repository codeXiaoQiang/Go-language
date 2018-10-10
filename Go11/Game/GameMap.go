package Game

import (
	"Clib"
	"fmt"
	"time"
)

// 1.定义地图类
type GameMap struct {
	// 行数
	row int
	// 列数
	col int
	// 地图信息
	detail string
}
// 2.对外提供一些公有的接口, 便于外界使用
func (mp *GameMap)SetRow(row int)  {
	mp.row = row
}
func (mp *GameMap)SetCol(col int)  {
	mp.col = col
}
func (mp *GameMap)SetDetail(detail string)  {
	mp.detail = detail
}
func (mp *GameMap)GetRow() int  {
	return mp.row
}
func (mp *GameMap)GetCol() int  {
	return mp.col
}
func (mp *GameMap)GetDetail() string  {
	return mp.detail
}

// 打印自己的
func (mp *GameMap)PrintMap(){
	fmt.Println(mp.detail)
}

// 绘制蛇方法
func (mp *GameMap)DrawSnake(sk GameSnake)  {
	// 绘制蛇身体
	for i:=sk.length-1; i>0; i--{
		Clib.GotoPostion(sk.pos[i].X, sk.pos[i].Y)
		fmt.Print("*")
	}

	// 绘制蛇头
	Clib.GotoPostion(sk.pos[0].X, sk.pos[0].Y)
	fmt.Print("@")
}
// 绘制食物方法
func (mp *GameMap)DrawFood(fd GameFood)  {
	// 将光标移动到指定的位置
	Clib.GotoPostion(fd.X, fd.Y)
	// 在指定位置绘制食物
	fmt.Print("O")
}
// 让蛇移动
func (mp *GameMap)MoveSnake(sk *GameSnake, fd *GameFood)  {
	for {
		// 每移动一次停留一会
		time.Sleep(time.Second / 3)
		// 2.根据当前蛇的方向计算偏移位
		switch sk.direction {
		case DirectionTop:
			sk.SetOffset(0, -1)
		case DirectionBottom:
			sk.SetOffset(0, 1)
		case DirectionLeft:
			sk.SetOffset(-1, 0)
		case DirectionRight:
			sk.SetOffset(1, 0)
		}
		// 3.清除最后一节蛇身体
		lastPos := sk.pos[sk.length - 1]
		Clib.GotoPostion(lastPos.X, lastPos.Y)
		fmt.Print(" ")

		// 4.修改蛇在地图中的位置
		// 4.1修改蛇身体在地图中的位置
		//  i:= 1
		for i:=sk.length-1; i>0; i--{
			// 让后一个蛇身体移动到前一个蛇身体的位置
			// sk.pos[1] = sk.pos[0]
			sk.pos[i] = sk.pos[i-1]
		}
		// 4.2修改蛇头在地图中的位置
		sk.pos[0].X += sk.GetOffset().X
		sk.pos[0].Y += sk.GetOffset().Y

		// 5.将修改好的蛇绘制到地图上
		mp.DrawSnake(*sk)

		// 6.做其它相关检测
		res := mp.Controll(sk, fd)

		// 7.如果碰到了身体或者食物, 那么结束游戏
		if res == true{
			return
		}
	}
}
// 蛇移动过程中的相关检测
func (mp *GameMap)Controll(sk *GameSnake, fd *GameFood) bool{
	// 1.检测有没有碰到墙壁
	headPos := sk.pos[0]
	if headPos.X <= 0 || headPos.X >= mp.GetCol()-1 ||
		headPos.Y <=0 || headPos.Y >= mp.GetRow()-1{
		return true
	}
	// 2.检测有没有碰到食物
	if headPos.X == fd.X && headPos.Y == fd.Y{
		// 2.1重新生成食物的位置
		fd.refreshPosition(*mp)
		// 2.2重新绘制食物
		mp.DrawFood(*fd)
		// 2.3新增一节蛇
		sk.appendSnakeBody()
		// 2.4重新绘制蛇
		mp.DrawSnake(*sk)
	}
	// 3.检测有没有碰到身体
	return  false
}
// 接收键盘输入
func (mp *GameMap)GetDet(sk *GameSnake){
	for{
		// 1.调用无显示接收函数, 接收用户输入
		switch Clib.Direction() {
		case 119, 87:
			//fmt.Println("向上")
			sk.direction = DirectionTop
		case 115, 83:
			//fmt.Println("向下")
			sk.direction = DirectionBottom
		case 97, 65:
			//fmt.Println("向左")
			sk.direction = DirectionLeft
		case 100, 68:
			//fmt.Println("向右")
			sk.direction = DirectionRight
		}
	}
}

// 工厂方法: 作用, 快速创建一个对象
func CreateGameMap(row , col int, detail string)(mp GameMap) {
	 mp = GameMap{row, col, detail}
	 return
}