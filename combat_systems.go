package main

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerPosition, defenderPosition *Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	// Get the attacker and defender if either is a player
	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[position].(*Position)
		if pos.IsEqual(attackerPosition) {
			attacker = playerCombatant
		} else if pos.IsEqual(defenderPosition) {
			defender = playerCombatant
		}
	}

	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[position].(*Position)
		if pos.IsEqual(attackerPosition) {
			attacker = cbt
		} else if pos.IsEqual(defenderPosition) {
			defender = cbt
		}
	}

	if attacker == nil || defender == nil {
		return
	}

	defenderArmor := defender.Components[armor].(*Armor)
	defenderHealth := defender.Components[health].(*Health)
	defenderName := defender.Components[name].(*Name)
	defenderMessage := defender.Components[userMessage].(*UserMessage)

	attackerWeapon := attacker.Components[meleeWeapon].(*MeleeWeapon)
	attackerName := attacker.Components[name].(*Name)
	attackerMessage := attacker.Components[userMessage].(*UserMessage)

	if attacker.Components[health].(*Health).CurrentHealth <= 0 {
		return
	}

	// Roll a d10 to hit
	toHitRoll := GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		// It's a hit
		damageRoll := GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)

		damageDone := damageRoll - defenderArmor.Defense
		// Prevent healing the defender
		if damageDone < 0 {
			damageDone = 0
		}

		defenderHealth.CurrentHealth -= damageDone
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and hits for %d health.\n", attackerName.Label, attackerWeapon.Name, defenderName.Label, damageDone)

		if defenderHealth.CurrentHealth <= 0 {
			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName.Label)
			if defenderName.Label == "Player" {
				defenderMessage.GameStateMessage = "Game Over!\n"
				g.Turn = GameOver
			}
		}
	} else {
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses.\n", attackerName.Label, attackerWeapon.Name, defenderName.Label)
	}
}
