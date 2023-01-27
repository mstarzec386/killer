package killer

import (
	"fmt"

	"mstarzec.pw/killer/internal/position"
)

type Killer struct {
	name string
	health int
	position *position.Position
}

func (k *Killer) SetName(name string) {
	k.name = name
}

func (k *Killer) SetHealth(n int) {
	k.health = n
}

func (k *Killer) SetPosition(p position.Position) {
	k.position = &p
	fmt.Printf("(%d %d) -> (%d %d)\n",p.GetX(), p.GetY(), k.position.GetX(), k.position.GetY())
}

func (k *Killer) Hit() {
	if k.health > 0 {
	    k.health -= 1
	}
}

func (k Killer) GetPosition() position.Position {
	return *k.position
}


func (k Killer) GetHealth() int {
	return k.health
}

func (k Killer) IsAlive() bool {
	return k.health > 0
}

func (k Killer) ToString() string {
	return fmt.Sprintf("%s - %d (%d, %d)", k.name, k.health, k.position.GetX(), k.position.GetY())
}