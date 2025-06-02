package main

import (
	"fmt"
)

func main() {
	list := NewMonsterList()

	monsters := []*Monster{
		CreateMonster("Globin", Air, 2),
		CreateMonster("Dragão", Fire, 5),
		CreateMonster("Golem", Earth, 3),
		CreateMonster("Hydra", Water, 7),
		CreateMonster("Orc", Earth, 4),
	}

	for _, monster := range monsters {
		fmt.Printf("%s speed: %d\n", monster.Name, monster.Speed)
		list.InsertBySpeed(monster)
	}

	fmt.Println()
	dragao := list.FindMonster("Dragão")
	hydra := list.FindMonster("Hydra")

	fmt.Printf("Dragão: %s\n", list.FindMonster("Dragão"))
	fmt.Printf("Hydra: %s\n", list.FindMonster("Hydra"))

	result := Attack(dragao, hydra, dragao.Skills[1])
	result.Display()

	result2 := Attack(hydra, dragao, hydra.Skills[1])
	result2.Display()
}
