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
	next    *Monster
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
	Head  *Monster
	Tail  *Monster
	Type  MonsterType
	Skill Skill
	Next  *Monster
}

func CreateMonster(name string, monsterType MonsterType, level int) *Monster {
	return &Monster{
		Name:  name,
		Type:  monsterType,
		Level: level,
	}
}

func (list *MonsterList) Insert(monster *Monster) {
	// newNode := &Monster{Name: Name}

	if list.Head == nil {
		list.Head = monster
		list.Tail = monster
	} else {
		list.Tail.next = monster
		list.Tail = monster
	}
}

func (list *MonsterList) Display() {
	current := list.Head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%v ", current.Name)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := MonsterList{}

	globin := CreateMonster("Globin", Earth, 2)
	dragon := CreateMonster("Dragão", Fire, 5)

	list.Insert(globin)
	list.Insert(dragon)
	// list.Insert("Golem")

	list.Display()
}
