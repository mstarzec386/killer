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
    g.board.PrintBoard()

    endGame := false

    for !endGame {
        endGame = g.nextRound()
    }

}

func (g GameController) nextRound() bool{
    killers := g.board.GetKillers()
    numberOfAlivedKillers := 0

    for _, currentKiller := range killers {
        g.board.PrintBoard()

        if currentKiller.IsAlive() {
            numberOfAlivedKillers += 1
            killerPosition := currentKiller.GetPosition()

            fmt.Println("current", currentKiller.ToString())
            opponents := g.findOpponents(killerPosition)
            if len(opponents) > 0 {
                weakest := findWeakestOpponent(opponents)
                weakest.Hit()
                if !weakest.IsAlive() {
                    numberOfAlivedKillers -= 1
                }
            } else {
                moveToPosition := g.findNewPosition(currentKiller)
                fmt.Printf("MOVE kurwa (%d %d) -> (%d %d) :: ", currentKiller.GetPosition().GetX(), currentKiller.GetPosition().GetY(), moveToPosition.GetX(), moveToPosition.GetY())
                g.board.MoveKiller(currentKiller, *moveToPosition)
                fmt.Printf("MOVE kurwa to samo powinno byc (%d %d) -> (%d %d)\n", currentKiller.GetPosition().GetX(), currentKiller.GetPosition().GetY(), moveToPosition.GetX(), moveToPosition.GetY())
            }
        }

        g.board.PrintBoard()
    } 


    return numberOfAlivedKillers < 2
}

func (g GameController) findOpponents(p position.Position) []*killer.Killer {
    nearPossitions := p.GetNearPositions()
    var opponents []*killer.Killer

    for _, possition := range nearPossitions {
        if possition.GetX() < MaxSize && possition.GetY() < MaxSize {
            opponent := g.board.GetPosition(possition)
            
            if opponent != nil && opponent.IsAlive() && (opponent.GetPosition().GetX() != possition.GetX() || opponent.GetPosition().GetY() != possition.GetY()) {
                fmt.Printf("KURWA NO POJEBANE %v %v", opponent.GetPosition(), possition);
            }

            if opponent != nil && opponent.IsAlive() {
                fmt.Printf("opponent (%s) (%d %d)\n", opponent.ToString(), possition.GetX(), possition.GetY())
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

    //  TODO possible endless for
    for !found && len(possiblePossitions) > 0 {
        randomPossiblePosition := possiblePossitions[helpers.GetRandomInt(len(possiblePossitions))]
        if randomPossiblePosition.GetX() < 10 && randomPossiblePosition.GetY() < 10 && g.board.GetPosition(randomPossiblePosition) == nil {
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
