package Game

import (
	"math/rand"
	"time"
)

// 1.定义一个食物类
type GameFood struct {
	Position
}

func init()  {
	// 4.1设置随机种子
	rand.Seed(time.Now().UnixNano())
}
// 2.对外提供一些公有的接口, 便于外界使用
func (fd *GameFood)SetX(x int)  {
	fd.X = x
}
func (fd *GameFood)SetY(y int)  {
	fd.Y = y
}
func (fd *GameFood)GetX() int {
	return fd.X
}
func (fd *GameFood)GetY() int {
	return fd.Y
}
// 刷新食物的位置
func (fd *GameFood)refreshPosition(mp GameMap)  {
	fd.Y = rand.Intn(mp.GetRow() - 2) + 1
	fd.X = rand.Intn(mp.GetCol() - 2) + 1
}
// 工厂方法: 作用, 快速创建一个对象
func CreateGameFood(x int, y int) GameFood {
	return GameFood{Position{x, y}}
}
func DefaultGameFood(mp GameMap) GameFood {
	// 4.2生成随机数
	fdRow := rand.Intn(mp.GetRow() - 2) + 1
	fdCol := rand.Intn(mp.GetCol() - 2) + 1
	return GameFood{Position{fdCol, fdRow}}
}