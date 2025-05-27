package main

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
