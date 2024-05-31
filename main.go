package main

import "fmt"

func main() {
	fmt.Println("Game Start")
	p := Player{}
	p.createNewCharactor("tomcha")
	p.outputStatus()

	ml := monsterList{}
	ml.createMonster()

	enemy := ml.selectMonster()
	fmt.Printf("Enemy name: %s\n", enemy.Name)
	fmt.Println("Game End")
}
