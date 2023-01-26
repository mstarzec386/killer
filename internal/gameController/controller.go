package gameController

import (
	"fmt"

	"mstarzec.pw/killer/internal/board"
	"mstarzec.pw/killer/internal/killer"
	"mstarzec.pw/killer/internal/position"
)

type GameController struct {
	playersNumber int
    board board.Board
}


func (g GameController) Run() {
    g.board.PlaceKillers(g.playersNumber)

    endGame := false

    for !endGame {
        endGame = g.nextRound()
    }

}

func (g GameController) nextRound() bool{
    killers := g.board.GetKillers()

    for _, killer := range killers {
        killerPosition := killer.GetPosition()

        opponent := g.findOpponent(killerPosition)
        if opponent == nil {
            moveToPosition := g.findNewPosition(killer)
            g.board.MoveKiller(killerPosition, moveToPosition)
        }
        
        fmt.Println(*killer)
    } 

    return true
}

func (g GameController) findOpponent(p position.Position) *killer.Killer {

    return nil
}

func (g GameController) findNewPosition(*killer.Killer) position.Position {

    return position.New(0, 0)
}

func New(players int) GameController {
    if (players > 100) {
        players = 100
    }

    fmt.Println("Number of players:", players)

    return GameController{playersNumber: players}
}
