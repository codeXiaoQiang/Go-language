package Game

// 蛇方向
const(
	DirectionTop = iota
	DirectionBottom = iota
	DirectionLeft = iota
	DirectionRight = iota
)
// 定义一个结构体, 用于保存位置
type Position struct {
	X int // 多少列
	Y int // 多行行
}
// 1.定义一个蛇类
type GameSnake struct {
	pos []Position // 用于保存蛇所有位置
	length int // 用于保存蛇长度
	direction int // 保存蛇当前运动的方向
	offset Position // 保存蛇移动的偏移位
}

// 2.对外提供一些公有的接口, 便于外界使用
func (sk *GameSnake)SetPos(position Position)  {
	sk.pos = append(sk.pos, position)
}
func (sk *GameSnake)SetLength(len int)  {
	sk.length = len
}
func (sk *GameSnake)SetDirection(det int)  {
	sk.direction = det
}
func (sk *GameSnake)SetOffset(x , y int){
	sk.offset = Position{x, y}
}
func (sk *GameSnake)GetPos() []Position{
	return sk.pos
}
func (sk *GameSnake)GetLength() int  {
	return sk.length
}
func (sk *GameSnake)GetDirection() int  {
	return sk.direction
}
func (sk *GameSnake)GetOffset() Position{
	return sk.offset
}

// 新增蛇长度的方法
func (sk *GameSnake)appendSnakeBody()  {
	// 1.创建一节空的身体
	newPos := Position{}
	lastPos := sk.pos[sk.length - 1]
	// 2.根据蛇运动的方向设置新身体的位置
	switch sk.direction {
	case DirectionTop:
		newPos.X = lastPos.X
		newPos.Y = lastPos.Y + 1
	case DirectionBottom:
		newPos.X = lastPos.X
		newPos.Y = lastPos.Y - 1
	case DirectionLeft:
		newPos.X = lastPos.X + 1
		newPos.Y = lastPos.Y
	case DirectionRight:
		newPos.X = lastPos.X - 1
		newPos.Y = lastPos.Y
	}
	// 3.将新的一节添加到蛇中
	sk.pos = append(sk.pos, newPos)
	sk.length += 1
}

// 工厂方法: 作用, 快速创建一个对象
func CreateGameSnake(pos []Position, len int, det int) GameSnake {
	sk := GameSnake{pos, len, det, Position{}}
	return sk
}
func DefaultGameSnake(mp GameMap) GameSnake {
	// 3.1初始化蛇头的位置
	headPo := Position{mp.GetCol() / 2, mp.GetRow() / 2}
	// 3.2初始化蛇身体的位置
	bodyPo := Position{mp.GetCol() / 2 - 1, mp.GetRow() / 2}
	// 3.3将所有蛇的位置放入切片
	pos := []Position{headPo, bodyPo}
	sk := GameSnake{pos, len(pos), DirectionRight, Position{}}
	return sk
}