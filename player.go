package main

import (
	"math/rand"
)

type Player struct {
	name    string
	skill   int
	stamina int
	luck    int
}

func (p *Player) createNewCharactor(n string) {
	p.name = n
	p.skill = rand.Intn(6) + 7
	p.stamina = rand.Intn(6) + rand.Intn(6) + 14
	p.luck = rand.Intn(6) + 7
}
