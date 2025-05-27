package main

import "fmt"

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

	fairy := CreateMonster("Fada", Water, 3)
	list.InsertAfter(monsters[2], fairy)

	// list.Display()

	list.RemoveMonster("Golem")
	fmt.Println()

	// list.Display()

	fmt.Println()

	dragao := list.FindMonster("Dragão")
	hydra := list.FindMonster("Hydra")
	globin := list.FindMonster("Globin")
	fada := list.FindMonster("Fada")
	orc := list.FindMonster("Orc")

	// monstroEncontrado := list.FindMonster("Dragão")
	// fmt.Println(monstroEncontrado)

	dano1 := CalculateDamage(globin, dragao, globin.Skills[1])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		globin.Name, globin.Skills[1].Name, dragao.Name, dano1)

	dano2 := CalculateDamage(globin, dragao, globin.Skills[0])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		globin.Name, globin.Skills[0].Name, dragao.Name, dano2)

	dano3 := CalculateDamage(dragao, hydra, dragao.Skills[1])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		dragao.Name, dragao.Skills[1].Name, hydra.Name, dano3)

	dano4 := CalculateDamage(hydra, dragao, hydra.Skills[1])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		hydra.Name, hydra.Skills[1].Name, dragao.Name, dano4)

	dano5 := CalculateDamage(fada, globin, fada.Skills[0])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		fada.Name, fada.Skills[0].Name, globin.Name, dano5)

	dano6 := CalculateDamage(globin, fada, globin.Skills[1])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		globin.Name, globin.Skills[1].Name, fada.Name, dano6)

	dano7 := CalculateDamage(orc, globin, orc.Skills[0])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		orc.Name, orc.Skills[1].Name, globin.Name, dano7)

	dano8 := CalculateDamage(orc, globin, orc.Skills[1])
	fmt.Printf("%s usou %s em %s e causou %d de dano!\n",
		orc.Name, orc.Skills[1].Name, globin.Name, dano8)

	fmt.Println("Danos causados")
	fmt.Printf("Dano 1 (Globin Tornado → Dragão): %d\n", dano1)
	fmt.Printf("Dano 2 (Globin Básico → Dragão): %d\n", dano2)
	fmt.Printf("Dano 3 (Dragão Fogo → Hydra): %d\n", dano3)
	fmt.Printf("Dano 4 (Hydra Água → Dragão): %d\n", dano4)
	fmt.Printf("Dano 5 (Fada Básico → Globin): %d\n", dano5)
	fmt.Printf("Dano 6 (Globin Tornado → Fada): %d\n", dano6)
	fmt.Printf("Dano 7 (Orc skill → Globin): %d\n", dano7)
	fmt.Printf("Dano 8 (Orc skill → Globin): %d\n", dano8)
}
