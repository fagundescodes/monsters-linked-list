package main

import "fmt"

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
		fmt.Println("Linked list is empty")
		return
	}

	current := list.Head

	index := 1

	for current != nil {
		fmt.Printf(
			"%d. Nome: %v, Tipo: %s, Nível: %d, HP: %d, Atk: %d, Def: %d, Spd: %d\n",
			index,
			current.Name,
			current.Type,
			current.Level,
			current.MaxHP,
			current.CurrentHP,
			current.Attack,
			current.Defense,
			current.Speed,
		)
		current = current.Next
		index++
	}
}
