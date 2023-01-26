package board

import (
	"fmt"

	"mstarzec.pw/killer/internal/position"
	"mstarzec.pw/killer/internal/helpers"
	"mstarzec.pw/killer/internal/killer"
)

type Board struct {
    board [10][10]*killer.Killer
	killers []*killer.Killer
}

func (b *Board) PlaceKillers(number int) {
	for i := 0; i < number; i++ {
		assigned := false

		for !assigned {
		  position := position.New(helpers.GetRandomInt(10), helpers.GetRandomInt(10))
		  fmt.Println(position)

		  if b.board[position.GetX()][position.GetY()] == nil {
			k := killer.Killer{}
			k.SetName(helpers.GenerateName(i))
			k.SetPosition(position)

			b.board[position.GetX()][position.GetY()] = &k
			b.killers = append(b.killers, &k)
		    assigned = true
		  }
		}
	}

	fmt.Println(b.board)
	fmt.Println(b.killers)

	fmt.Println(b.board)
	fmt.Println(b.killers)
	fmt.Println(*b.killers[0])
}

func (b *Board) MoveKiller(from position.Position, to position.Position) {
	k := b.board[from.GetX()][from.GetY()] 

	b.board[from.GetX()][from.GetY()] = nil
	b.board[to.GetX()][to.GetY()] = k
	k.SetPosition(to)
}

func (b Board) GetKillers() []*killer.Killer {
	return b.killers
}