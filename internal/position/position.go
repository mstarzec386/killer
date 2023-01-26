package position

type Position struct{
	x int
	y int
}

func (p Position) GetX() int {
	return p.x
}

func (p Position) GetY() int {
	return p.y
}

func New(x int, y int) Position {
	return Position{x: x, y: y}
}
