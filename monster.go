package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type list struct {
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

func (m *monster) createMonster() {
	l := list{}
	f, err := os.Open("monsterData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&l)
	fmt.Println(l)
}
