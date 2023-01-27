package board

import (
	"fmt"
	"time"

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

		  if b.board[position.GetX()][position.GetY()] == nil {
			k := killer.Killer{}
			k.SetName(helpers.GenerateName(i))
			k.SetPosition(position)
			k.SetHealth(9)

			b.board[position.GetX()][position.GetY()] = &k
			b.killers = append(b.killers, &k)
		    assigned = true
		  }
		}
	}
}

func (b *Board) MoveKiller(killer *killer.Killer, to position.Position) {
	from := killer.GetPosition()

	fmt.Println("move from", from.GetX(), from.GetY())
	fmt.Println("move to", to.GetX(), to.GetY())

	killer.SetPosition(to)

	b.board[to.GetX()][to.GetY()] = killer
	b.board[from.GetX()][from.GetY()] = nil

	b.PrintBoard()
}

func (b Board) GetPosition(p *position.Position) *killer.Killer {
	fmt.Printf("get position (%d, %d) -> ", p.GetX(), p.GetY())
	k := b.board[p.GetX()][p.GetY()]
	if (k != nil) {
		fmt.Printf("(%d, %d)", k.GetPosition().GetX(), k.GetPosition().GetY())
	}
	fmt.Println(".")
	return b.board[p.GetX()][p.GetY()]
}

func (b Board) GetKillers() []*killer.Killer {
	return b.killers
}

func (b Board) PrintBoard() {
	fmt.Println(" ---------------------------------------")
	for _, row := range b.board {
			fmt.Printf("|")
        for _, killer := range row {
			if killer != nil {
				fmt.Printf(" %d |", killer.GetHealth())
			} else {
				fmt.Printf(" - |")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(" ---------------------------------------")
	time.Sleep(50 * time.Millisecond)
}