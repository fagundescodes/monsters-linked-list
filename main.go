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
	fmt.Println("--- Formação inicial ---")
	list.Display()

	fmt.Println("\n--- Iniciando combate automático ---")
	BattleByTurn(list, 2)

	fmt.Println("\n--- Estado final ---")
	list.Display()
}
