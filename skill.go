package main

type Skill struct {
	Name        string
	Power       int
	Description string
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
