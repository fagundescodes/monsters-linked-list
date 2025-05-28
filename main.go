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

	dragao := list.FindMonster("Dragão")
	// hydra := list.FindMonster("Hydra")
	// globin := list.FindMonster("Globin")
	// fada := list.FindMonster("Fada")
	orc := list.FindMonster("Orc")

	monstroEncontrado := list.FindMonster("Dragão")
	fmt.Println(monstroEncontrado)

	morreu := dragao.TakeDamage(30)
	fmt.Printf("Dragão após dano: %s\n", dragao)
	fmt.Printf("HP do Dragão depois de sofrer dano: %.1f%%\n", dragao.HPercentage())
	fmt.Printf("Morto? %t\n", morreu)

	dano := CalculateDamage(orc, dragao, orc.Skills[0])
	dragao.TakeDamage(dano)
	fmt.Printf("%s usou %s em %s e causou %v de dano!\n",
		orc.Name, orc.Skills[1].Name, dragao.Name, dano)
	fmt.Printf("Dragão após dano: %s\n", dragao)
	fmt.Printf("HP do Dragão depois de sofrer dano: %.1f%%\n", dragao.HPercentage())
	fmt.Printf("Morto? %t\n", morreu)

	dano2 := dragao.TakeDamage(70)
	fmt.Printf("Dragão após dano: %s\n", dragao)
	fmt.Printf("HP do Dragão depois de sofrer dano: %.1f%%\n", dragao.HPercentage())
	fmt.Printf("Morto? %t\n", dano2)
}
