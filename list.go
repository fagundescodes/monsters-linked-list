package main

import (
	"fmt"
	"strings"
)

type MonsterList struct {
	Head   *Monster
	Tail   *Monster
	Length int
}

func NewMonsterList() *MonsterList {
	return &MonsterList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (list *MonsterList) InsertBySpeed(monster *Monster) {
	if list.Head == nil {
		list.Head = monster
		list.Tail = monster
		monster.Next = nil
	} else if monster.Speed > list.Head.Speed {
		monster.Next = list.Head
		list.Head = monster
	} else {
		current := list.Head
		for current.Next != nil && current.Next.Speed > monster.Speed {
			current = current.Next
		}

		monster.Next = current.Next
		current.Next = monster
	}

	list.Length++
}

func (list *MonsterList) Insert(monster *Monster) {
	if list.Head == nil {
		list.Head = monster
		list.Tail = monster
	} else {
		list.Tail.Next = monster
		list.Tail = monster
	}
	list.Length++
}

func (list *MonsterList) RemoveMonster(name string) bool {
	if list.Head == nil {
		return false
	}

	if list.Head.Name == name {
		list.Head = list.Head.Next
		if list.Head == nil {
			list.Tail = nil
		}
		list.Length--
		return true
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Name == name {
			removedMonster := current.Next
			current.Next = removedMonster.Next

			if removedMonster == list.Tail {
				list.Tail = current
			}
			list.Length--
			return true
		}
		current = current.Next
	}
	return false
}

func (list *MonsterList) InsertAfter(target *Monster, newMonster *Monster) bool {
	if target == nil {
		return false
	}

	nextMonster := target.Next

	target.Next = newMonster
	newMonster.Next = nextMonster

	if target == list.Tail {
		list.Tail = newMonster
	}

	list.Length++
	return true
}

func (list *MonsterList) FindMonster(name string) *Monster {
	current := list.Head
	for current != nil {
		if current.Name == name {
			return current
		}
		current = current.Next
	}

	return nil
}

func (list *MonsterList) Display() {
	if list.Head == nil {
		fmt.Println("A lista está vazia")
		return
	}

	current := list.Head

	index := 1

	for current != nil {
		status := "Vivo"
		if !current.Alive() {
			status = "Morto"
		}

		createHp := HPBar(current.CurrentHP, current.MaxHP, 10)

		fmt.Printf("%d. [%s] %s (%s)\n", index, status, current.Name, current.Type)
		fmt.Printf("   HP: %s %d/%d (%.1f%%)\n",
			createHp, current.CurrentHP, current.MaxHP, current.HPercentage())
		fmt.Printf("   Atk: %d | Def: %d | Spd: %d | Lvl: %d\n",
			current.Attack, current.Defense, current.Speed, current.Level)
		fmt.Println("   " + strings.Repeat("-", 50))

		current = current.Next
		index++
	}
}

func HPBar(curent, max, barLength int) string {
	if max == 0 {
		return strings.Repeat(" ", barLength)
	}

	filledUP := int(float64(curent) / float64(max) * float64(barLength))
	if filledUP < 0 {
		filledUP = 0
	}
	if filledUP > barLength {
		filledUP = barLength
	}

	bar := strings.Repeat("█", filledUP)
	bar += strings.Repeat("░", barLength-filledUP)

	return bar
}

func (list *MonsterList) GetAliveMonster() []*Monster {
	var alive []*Monster
	current := list.Head
	for current != nil {
		if current.Alive() {
			alive = append(alive, current)
		}
		current = current.Next
	}
	return alive
}
