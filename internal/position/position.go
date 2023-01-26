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

func (p Position) GetNearPositions() []*Position {
	var possitions []*Position

	x := p.x
	y := p.y
	possitions = append(possitions, &Position{x: x + 0, y: y + 1})
	possitions = append(possitions, &Position{x: x + 1, y: y})
	possitions = append(possitions, &Position{x: x + 1, y: y + 1})

	if x > 0 {
		possitions = append(possitions, &Position{x: x - 1, y: y + 1})
		possitions = append(possitions, &Position{x: x - 1, y: y})
	} 	

	if y > 0 {
		possitions = append(possitions, &Position{x: x + 1, y: y - 1})
		possitions = append(possitions, &Position{x: x, y: y - 1})
	} 	

	if x > 0 && y > 0 {
		possitions = append(possitions, &Position{x: x - 1, y: y - 1})
	} 	

	return possitions
}

func New(x int, y int) Position {
	return Position{x: x, y: y}
}
