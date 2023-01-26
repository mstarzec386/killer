package killerApp

import (
	"fmt"
	"os"
	"mstarzec.pw/killer/internal/gameController"
)

func KillerRun() {
	fmt.Println("starting", os.Getpid())
	game := gameController.New(10)
	game.Run()
}