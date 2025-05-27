package main

import "fmt"

type MonsterType string

const (
	Fire  MonsterType = "Fogo"
	Water MonsterType = "Água"
	Earth MonsterType = "Terra"
	Air   MonsterType = "Ar"
)

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
