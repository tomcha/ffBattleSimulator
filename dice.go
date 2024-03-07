package main

import (
	"math/rand"
)

type Dice struct {
	diceValue int
}

func (d *Dice) rollDice(piece int, number int) {
	v := 0
	for i := 0; i < piece; i++ {
		v += rand.Intn(number) + 1
	}
	d.diceValue = v
}