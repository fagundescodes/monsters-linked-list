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
	Name      string
	Type      MonsterType
	Level     int
	CurrentHP int
	MaxHP     int
	Attack    int
	Defense   int
	Speed     int
	Skills    []Skill
	Next      *Monster
}

func (m *Monster) String() string {
	return fmt.Sprintf("Nome: %s, Tipo: %s, Nível: %d, HP: %d, Atk: %d, Def: %d, Speed: %v",
		m.Name,
		m.Type,
		m.Level,
		m.CurrentHP,
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
		Name:      name,
		Type:      monsterType,
		Level:     level,
		CurrentHP: baseHp,
		MaxHP:     baseHp,
		Attack:    baseAtk,
		Defense:   baseDef,
		Speed:     baseSpd,
		Skills:    skills,
	}
}

func (m *Monster) Alive() bool {
	return m.CurrentHP > 0
}

func (m *Monster) TakeDamage(damage int) bool {
	m.CurrentHP -= damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
	return !m.Alive()
}

func (m *Monster) HPercentage() float64 {
	if m.MaxHP == 0 {
		return 0
	}

	return float64(m.CurrentHP) / float64(m.MaxHP) * 100
}
