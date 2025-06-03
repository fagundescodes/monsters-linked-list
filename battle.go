package main

import "fmt"

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

type BattleResult struct {
	Attacker                *Monster
	Defender                *Monster
	Skill                   Skill
	Damage                  int
	DiedMonster             bool
	EffectivenessMultiplier float64
}

func CalculateDamage(attacker *Monster, defender *Monster, skill Skill) int {
	baseDamage := attacker.Attack + skill.Power - defender.Defense

	baseDamage = max(baseDamage, 1)

	multiplier := effectiveness[attacker.Type][defender.Type]

	finalDamage := int(float64(baseDamage) * multiplier)

	return finalDamage
}

func Attack(attacker *Monster, defender *Monster, skill Skill) BattleResult {
	damage := CalculateDamage(attacker, defender, skill)
	died := defender.TakeDamage(damage)
	multiplier := effectiveness[attacker.Type][defender.Type]

	return BattleResult{
		Attacker:                attacker,
		Defender:                defender,
		Skill:                   skill,
		Damage:                  damage,
		DiedMonster:             died,
		EffectivenessMultiplier: multiplier,
	}
}

func (br BattleResult) Display() {
	effectiveness := ""
	if br.EffectivenessMultiplier > 1.0 {
		effectiveness = "Super Efetivo!"
	} else if br.EffectivenessMultiplier < 1.0 {
		effectiveness = "Não Efetivo!"
	}

	fmt.Printf("%s usou %s em %s\n",
		br.Attacker.Name, br.Skill.Name, br.Defender.Name)

	fmt.Printf("Dano causado: %d %s\n", br.Damage, effectiveness)

	fmt.Printf("%s: %d/%d HP (%.1f%%)\n",
		br.Defender.Name, br.Defender.CurrentHP, br.Defender.MaxHP, br.Defender.HPercentage())

	if br.DiedMonster {
		fmt.Printf("%s foi derrotado\n", br.Defender.Name)
	}
}

func BattleByTurn(list *MonsterList, rounds int) {
	for round := 1; round <= rounds; round++ {
		aliveMonsters := list.GetAliveMonster()
		if len(aliveMonsters) < 2 {
			fmt.Println("Batalha encerrada")
			break
		}
		for _, attacker := range aliveMonsters {
			if !attacker.Alive() {
				continue
			}

			var targets []*Monster
			for _, monster := range aliveMonsters {
				if monster != attacker && monster.Alive() {
					targets = append(targets, monster)
				}
			}

			if len(targets) == 0 {
				continue
			}
			target := targets[0]
			skill := attacker.Skills[0]
			if len(attacker.Skills) > 1 {
				skill = attacker.Skills[1]
			}
			result := Attack(attacker, target, skill)
			result.Display()
		}

		removed := list.RemoveDeadMonster()
		if removed > 0 {
			fmt.Printf("%d monstro removido da batalha\n", removed)
		}
	}
	fmt.Println("Fim da Batalha!")
}
