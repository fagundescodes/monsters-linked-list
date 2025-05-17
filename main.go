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

func CreateSkill(name string, power int, description string) Skill {
	return Skill{
		Name:        name,
		Power:       power,
		Description: description,
	}
}

func CreateMonsterSkills(monsterType MonsterType, baseAtk int) []Skill {
	skills := []Skill{
		CreateSkill("Ataque Básico", baseAtk, "Ataque básico"),
	}

	switch monsterType {
	case Fire:
		skills = append(
			skills,
			CreateSkill("Bola de Fogo", baseAtk*2, "Dispara uma bola de fogo no inimigo"),
		)
	case Water:
		skills = append(
			skills,
			CreateSkill("Jato d'Água", int(float64(baseAtk)*1.5), "Poderoso jato de água"),
		)
	case Earth:
		skills = append(
			skills,
			CreateSkill("Terremoto", int(float64(baseAtk)*1.2), "Causa dano em àrea"),
		)
	case Air:
		skills = append(
			skills,
			CreateSkill("Tornado", int(float64(baseAtk)*1.7), "Ataque de ventos cortantes"),
		)
	}

	return skills
}

func CreateMonster(name string, monsterType MonsterType, level int) *Monster {
	baseAtk := 10 + level*2
	baseHp := 50 + level*10
	baseDef := 3 + level*2

	skills := CreateMonsterSkills(monsterType, baseAtk)

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
	list := NewMonsterList()

	monsters := []*Monster{
		CreateMonster("Globin", Earth, 2),
		CreateMonster("Dragão", Fire, 5),
		CreateMonster("Golem", Earth, 3),
		CreateMonster("Hydra", Water, 7),
	}

	for _, monster := range monsters {
		list.Insert(monster)
	}

	fairy := CreateMonster("Fada", Water, 3)
	list.InsertAfter(monsters[2], fairy)

	list.Display()

	list.RemoveMonster("Golem")
	fmt.Println()

	list.Display()

	fmt.Println()

	monstroEncontrado := list.FindMonster("Dragão")
	fmt.Println(monstroEncontrado)
}
