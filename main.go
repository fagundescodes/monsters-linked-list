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
		list.Insert(monster)
	}

	fada := CreateMonster("Fada", Fire, 3)
	list.Insert(fada)

	dragao := list.FindMonster("Dragão")
	hydra := list.FindMonster("Hydra")
	globin := list.FindMonster("Globin")

	list.Display()
	fmt.Println()

	dragao.TakeDamage(30)
	hydra.TakeDamage(100)
	globin.TakeDamage(47)
	orc := list.FindMonster("Orc")
	orc.TakeDamage(100)

	fmt.Println()

	list.Display()
}
