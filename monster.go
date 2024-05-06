package main

import (
	"encoding/json"
	"log"
	"os"
)

type monsterList struct {
	Name     string
	Monsters []monster
}

type monster struct {
	Name         string
	Id           int
	SkillPoint   int
	StaminaPoint int
	AttackPoint  int
}

func (ml *monsterList) createMonster() {
	f, err := os.Open("monsterData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&ml)
}

func (ml *monsterList) selectMonster() monster {
	d := Dice{}
	d.rollDice(1, 6)
	return ml.Monsters[d.diceValue-1]
}
