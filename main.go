package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	name    string
	skill   int
	stamina int
	luck    int
}

func (c Character) calcAttack() int {
	attackPoint := c.skill + roleDice(2, 6)
	return attackPoint
}

func (c Character) calcDamaged() {
	c.stamina -= 2
}

type Monster struct {
	name    string
	skill   int
	stamina int
	attack  int
}

func main() {
	var c Character
	c.name = "Tomcha"
	c.skill = roleDice(1, 6) + 6
	c.stamina = roleDice(2, 6) + 12
	c.luck = roleDice(2, 6)

	fmt.Println(c)

	var m Monster
	m.name = "goblin"
	m.skill = 5
	m.stamina = 3
	m.attack = 1

	fmt.Println(m)
	battle(&c, &m)
	fmt.Println(c)
	fmt.Println(m)
}

func roleDice(times int, hedron int) int {
	rand.Seed(time.Now().UnixNano())
	total := 0
	for i := 0; i < times; i++ {
		total += rand.Intn(hedron) + 1
	}
	return total
}

func battle(c *Character, m *Monster) {
	var winner string
loop:
	for {
		if c.stamina <= 0 || m.stamina <= 0 {
			break loop
		}
		ca := c.calcAttack()
		ma := roleDice(2, 6) + m.skill
		if ca > ma {
			m.stamina -= 2
		} else if ma > ca {
			c.calcDamaged()
		}
	}

	if c.stamina > m.stamina {
		winner = c.name
	} else {
		winner = m.name
	}
	fmt.Println(winner, "is win")
}
