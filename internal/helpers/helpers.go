package helpers

import (
	"fmt"
	"time"
	"math/rand"
)

func GenerateName(n int) string {
    return fmt.Sprintf("Player%d", n)
}

func GetRandomInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}