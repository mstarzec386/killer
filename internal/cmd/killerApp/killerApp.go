package killerApp

import (
	"fmt"
	"mstarzec.pw/killer/internal/gameController"
)

func KillerRun() {
	fmt.Println("Let's play")
	game := gameController.New(10)
	game.Run()
}