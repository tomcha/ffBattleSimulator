package main

import (
	"fmt"
)

type Player struct {
	name    string
	skill   int
	stamina int
	luck    int
}

func (p *Player) createNewCharactor(n string) {
	p.name = n
	d := Dice{}
	d.rollDice(1, 6)
	p.skill = d.diceValue + 6
	d.rollDice(2, 6)
	p.stamina = d.diceValue + 12
	d.rollDice(1, 6)
	p.luck = d.diceValue + 6
}

func (p *Player) outputStatus() {
	fmt.Printf("%s のステータス\n", p.name)
	fmt.Printf("  skill   : %d\n", p.skill)
	fmt.Printf("  stamina : %d\n", p.stamina)
	fmt.Printf("  luck    : %d\n", p.luck)
}

func (p *Player) calcAttackPoint() (ap int) {
	d := Dice{}
	d.rollDice(2, 6)
	ap = d.diceValue + p.skill
	return
}

func (p *Player) calcDamage(dp int) (isDead bool) {
	isDead = false
	p.stamina -= dp
	if p.stamina <= 0 {
		isDead = true
	}
	return
}

func (p *Player) lucktest() (bool, error) {
	var isLucky bool
	var err error
	if p.luck <= 0 {
		isLucky = false
		err = fmt.Errorf("error")
		return isLucky, err
	}
	d := Dice{}
	d.rollDice(2, 6)
	judgePoint := d.diceValue
	if judgePoint > p.luck {
		isLucky = false
	} else {
		isLucky = true
	}
	p.luck -= 1
	err = nil
	return isLucky, err
}

func (p *Player) isDead() bool {
	if p.stamina <= 0 {
		return true
	}
	return false
}
