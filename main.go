package main

import "fmt"

type Monster struct {
	Name    string
	Type    MonsterType
	Level   int
	HP      int
	Attack  int
	Defense int
	Skills  []string
	Next    *Monster
}

type MonsterType string

type Skill struct {
	Name  string
	Power int
}

const (
	Fire  MonsterType = "Fogo"
	Water MonsterType = "Água"
	Earth MonsterType = "Terra"
	Air   MonsterType = "Ar"
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

func CreateMonster(name string, monsterType MonsterType, level int) *Monster {
	return &Monster{
		Name:   name,
		Type:   monsterType,
		Level:  level,
		HP:     50 + level*10,
		Attack: 10 + level*2,
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

func (list *MonsterList) Display() {
	if list.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	current := list.Head

	index := 1

	for current != nil {
		fmt.Printf("%d. Nome: %v, Tipo: %s, Nível: %d, HP: %d\n",
			index, current.Name, current.Type, current.Level, current.HP)
		current = current.Next
		index++
	}
}

func main() {
	list := MonsterList{}

	globin := CreateMonster("Globin", Earth, 2)
	dragon := CreateMonster("Dragão", Fire, 5)
	golem := CreateMonster("Golem", Earth, 3)
	hydra := CreateMonster("Hydra", Water, 7)

	list.Insert(globin)
	list.Insert(dragon)
	list.Insert(golem)
	list.Insert(hydra)

	list.Display()

	list.RemoveMonster("Golem")
	fmt.Println()

	list.Display()
}
