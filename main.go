package main

import "fmt"

type Monster struct {
	Name    string
	Type    MonsterType
	Level   int
	HP      int
	Attack  int
	Defense int
	Skills  []Skill
	Next    *Monster
}

func (m *Monster) String() string {
	return fmt.Sprintf("Nome: %s, Tipo: %s, Nível: %d, HP: %d, Atk: %d, Def: %d",
		m.Name,
		m.Type,
		m.Level,
		m.HP,
		m.Attack,
		m.Defense)
}

type MonsterType string

type Skill struct {
	Name        string
	Power       int
	Description string
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
	baseAtk := 10 + level*2
	baseHp := 50 + level*10
	baseDef := 3 + level*2

	skills := []Skill{
		{Name: "Ataque Básico", Power: baseAtk, Description: "Um ataque básico"},
	}

	switch monsterType {
	case Fire:
		skills = append(
			skills,
			Skill{
				Name: "Bola de fogo",
			},
		)
	case Water:
		skills = append(skills,

			Skill{
				Name: "Jato de água",
			},
		)
	}

	return &Monster{
		Name:    name,
		Type:    monsterType,
		Level:   level,
		HP:      baseHp,
		Attack:  baseAtk,
		Defense: baseDef,
		Skills:  skills,
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
			"%d. Nome: %v, Tipo: %s, Nível: %d, HP: %d, Atk: %d, Def: %d\n",
			index,
			current.Name,
			current.Type,
			current.Level,
			current.HP,
			current.Attack,
			current.Defense,
		)
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

	fmt.Println()

	monstroEncontrado := list.FindMonster("Dragão")
	fmt.Println(monstroEncontrado)
}
