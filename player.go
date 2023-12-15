package main

import (
	"fmt"
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

func(p *Player) outputStatus(){
	fmt.Printf("%s のステータス\n", p.name)
	fmt.Printf("  skill   : %d\n", p.skill)
	fmt.Printf("  stamina : %d\n", p.stamina)
	fmt.Printf("  luck    : %d\n", p.luck)
}