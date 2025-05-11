package main

import "fmt"

type Monster struct {
	Name string
	Type MonsterType
	next *Monster
}

type MonsterType string

const (
	Fire MonsterType = "Fogo"
)

type MonsterList struct {
	Head *Monster
	Tail *Monster
}

func (list *MonsterList) Insert(Name string) {
	newNode := &Monster{Name: Name}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.next = newNode
		list.Tail = newNode
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

	list.Insert("Globin")
	list.Insert("Dragão")
	list.Insert("Golem")

	list.Display()
}
