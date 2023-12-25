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

func (p *Player) outputStatus() {
	fmt.Printf("%s のステータス\n", p.name)
	fmt.Printf("  skill   : %d\n", p.skill)
	fmt.Printf("  stamina : %d\n", p.stamina)
	fmt.Printf("  luck    : %d\n", p.luck)
}

func (p *Player) calcAttackPoint() (ap int) {
	ap = rand.Intn(6) + rand.Intn(6) + 2 + p.skill
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
	judgePoint := rand.Intn(6) + rand.Intn(6) + 2
	if judgePoint > p.luck {
		isLucky = false
	} else {
		isLucky = true
	}
	p.luck -= 1
	err = nil
	return isLucky, err
}
