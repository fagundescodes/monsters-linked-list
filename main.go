package main

import "fmt"

type Monster struct {
	Name    string
	Type    MonsterType
	Level   int
	HP      int
	Attack  int
	Defense int
	Speed   int
	Skills  []Skill
	Next    *Monster
}

func (m *Monster) String() string {
	return fmt.Sprintf("Nome: %s, Tipo: %s, Nível: %d, HP: %d, Atk: %d, Def: %d, Speed: %d",
		m.Name,
		m.Type,
		m.Level,
		m.HP,
		m.Attack,
		m.Defense,
		m.Speed,
	)
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
	baseAtk := 8 + int(2.5*float64(level))
	baseHp := 50 + level*12
	baseDef := 5 + int(1.8*float64(level))
	baseSpd := 10 + int(1.4*float64(level))

	skills := CreateMonsterSkills(monsterType, baseAtk)

	return &Monster{
		Name:    name,
		Type:    monsterType,
		Level:   level,
		HP:      baseHp,
		Attack:  baseAtk,
		Defense: baseDef,
		Speed:   baseSpd,
		Skills:  skills,
	}
}

var effectiveness = map[MonsterType]map[MonsterType]float64{
	Fire: {
		Fire:  1.0,
		Water: 0.5,
		Earth: 1.5,
		Air:   1.0,
	},
	Water: {
		Fire:  1.5,
		Water: 1.0,
		Earth: 1.0,
		Air:   0.5,
	},
	Earth: {
		Fire:  0.5,
		Water: 1.0,
		Earth: 1.0,
		Air:   1.5,
	},
	Air: {
		Fire:  1.0,
		Water: 1.5,
		Earth: 0.5,
		Air:   1.0,
	},
}

func CalculateDamage(attacker *Monster, defender *Monster, skill Skill) int {
	baseDamage := attacker.Attack + skill.Power - defender.Defense

	baseDamage = max(baseDamage, 1)

	multiplier := effectiveness[attacker.Type][defender.Type]

	finalDamage := int(float64(baseDamage) * multiplier)

	return finalDamage
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
			current.HP,
			current.Attack,
			current.Defense,
			current.Speed,
		)
		current = current.Next
		index++
	}
}

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
