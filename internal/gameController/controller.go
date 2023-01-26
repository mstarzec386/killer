package gameController

import (
	"fmt"

	"mstarzec.pw/killer/internal/board"
	"mstarzec.pw/killer/internal/helpers"
	"mstarzec.pw/killer/internal/killer"
	"mstarzec.pw/killer/internal/position"
)

const MaxSize = 10

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
    numberOfAlivedKillers := 0

    for _, currentKiller := range killers {
        if currentKiller.IsAlive() {
            numberOfAlivedKillers += 1
            killerPosition := currentKiller.GetPosition()

            opponents := g.findOpponents(killerPosition)
            if len(opponents) > 0 {
                weakest := findWeakestOpponent(opponents)
                weakest.Hit()
            } else {
                moveToPosition := g.findNewPosition(currentKiller)
                g.board.MoveKiller(killerPosition, *moveToPosition)
            }
        }
    } 

    g.board.PrintBoard()

    return numberOfAlivedKillers < 2
}

func (g GameController) findOpponents(p position.Position) []*killer.Killer {
    nearPossitions := p.GetNearPositions()
    var opponents []*killer.Killer

    for _, possition := range nearPossitions {
        if possition.GetX() < MaxSize && possition.GetY() < MaxSize {
            opponent := g.board.GetPosition(*possition)
            if opponent != nil && opponent.IsAlive() {
                opponents = append(opponents, opponent)
            }
        }
    }

    return opponents
}

func (g GameController) findNewPosition(k *killer.Killer) *position.Position {
    currentKillerPosition := k.GetPosition()
    x := currentKillerPosition.GetX()
    y := currentKillerPosition.GetY()

    possiblePossitions := currentKillerPosition.GetNearPositions()
    found := false

    for !found && len(possiblePossitions) > 0 {
        randomPossiblePosition := possiblePossitions[helpers.GetRandomInt(len(possiblePossitions))]
        if randomPossiblePosition.GetX() < 10 && randomPossiblePosition.GetY() < 10 && g.board.GetPosition(*randomPossiblePosition) == nil {
            found = true
            x = randomPossiblePosition.GetX()
            y = randomPossiblePosition.GetY()
        }
    }

    newPosition := position.New(x, y)

    return &newPosition
}

func findWeakestOpponent(opponents []*killer.Killer) *killer.Killer {
    weakest := opponents[0];

    for _, opponent := range opponents[1:] {
        if (opponent.GetHealth() < weakest.GetHealth()) {
            weakest = opponent
            break
        }
    }

    return weakest
}

func New(players int) GameController {
    if (players > 100) {
        players = 100
    }

    fmt.Println("Number of players:", players)

    return GameController{playersNumber: players}
}
