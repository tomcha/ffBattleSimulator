package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Game Start")
	p := Player{}
	p.createNewCharactor("tomcha")
	p.outputStatus()

	ml := monsterList{}
	ml.createMonster()

	enemy := ml.selectMonster()
	fmt.Printf("Enemy name: %s\n", enemy.Name)
	round := 1
	for !(p.isDead()) {
		fmt.Printf("Round %d fight!\n", round)
		hap := p.calcAttackPoint()
		eap := enemy.SkillPoint + rand.Intn(6) + 1 + rand.Intn(6) + 1
		if eap < hap {
			fmt.Printf("%s 's Attack is hit!\n", p.name)
			enemy.StaminaPoint -= 2
		} else if eap > hap {
			fmt.Printf("%s 's Attack is hit!\n", enemy.Name)
			p.stamina -= enemy.AttackPoint
		} else {
			fmt.Println("Attack is miss")
		}

		if enemy.StaminaPoint <= 0 {
			fmt.Printf("%s is dead! You win!\n", enemy.Name)
			break
		}
		round += 1
	}
	if p.isDead() {
		fmt.Printf("%s is dead! You lose!\n")
	}
	fmt.Println("Game End")
}
