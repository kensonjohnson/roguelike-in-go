package main

import (
	"fmt"

	"github.com/bytearena/ecs"
	"github.com/kensonjohnson/roguelike-game-go/components"
)

func AttackSystem(g *Game, attackerPosition, defenderPosition *components.Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	// Get the attacker and defender if either is a player
	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[position].(*components.Position)
		if pos.IsEqual(attackerPosition) {
			attacker = playerCombatant
		} else if pos.IsEqual(defenderPosition) {
			defender = playerCombatant
		}
	}

	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[position].(*components.Position)
		if pos.IsEqual(attackerPosition) {
			attacker = cbt
		} else if pos.IsEqual(defenderPosition) {
			defender = cbt
		}
	}

	if attacker == nil || defender == nil {
		return
	}

	defenderArmor := defender.Components[armor].(*components.Armor)
	defenderHealth := defender.Components[health].(*components.Health)
	defenderName := defender.Components[name].(*components.Name)
	defenderMessage := defender.Components[userMessage].(*components.UserMessage)

	attackerWeapon := attacker.Components[meleeWeapon].(*components.MeleeWeapon)
	attackerName := attacker.Components[name].(*components.Name)
	attackerMessage := attacker.Components[userMessage].(*components.UserMessage)

	if attacker.Components[health].(*components.Health).CurrentHealth <= 0 {
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
