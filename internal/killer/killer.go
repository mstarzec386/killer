package killer

import (
	"mstarzec.pw/killer/internal/position"
)

type Killer struct {
	name string
	health int
	position position.Position
}

func (k *Killer) SetName(name string) {
	k.name = name
}

func (k *Killer) SetHealth(n int) {
	k.health = n
}

func (k *Killer) SetPosition(p position.Position) {
	k.position = p
}

func (k Killer) GetPosition() position.Position {
	return k.position
}