package combat

import (
	"fmt"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/yohamta/donburi"
)

func AttackSystem(attacker, defender *donburi.Entry) {
	attackerName := component.Name.Get(attacker)
	attackerWeapon := component.Weapon.Get(attacker)
	attackerHealth := component.Health.Get(attacker)
	attackerMessages := component.UserMessage.Get(attacker)
	defenderName := component.Name.Get(defender)
	defenderArmor := component.Armor.Get(defender)
	defenderHealth := component.Health.Get(defender)
	// defenderMessages := component.UserMessage.Get(defender)

	if attackerHealth.CurrentHealth <= 0 {
		return
	}

	toHitRoll := engine.GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		// Its a hit!
		damageRoll := engine.GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)
		damageDone := damageRoll - defenderArmor.Defense
		// Prevent healing the defender
		if damageDone < 0 {
			damageDone = 0
		}

		defenderHealth.CurrentHealth -= damageDone
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and hits for %d health.\n", attackerName.Label, attackerWeapon.ActionText, defenderName.Label, damageDone)
	} else {
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and misses.\n", attackerName.Label, attackerWeapon.ActionText, defenderName.Label)
	}
}

// func attackSystem(   ) {

// 	defenderArmor := defender.Components[armor].(*components.Armor)
// 	defenderHealth := defender.Components[health].(*components.Health)
// 	defenderName := defender.Components[name].(*components.Name)
// 	defenderMessage := defender.Components[userMessage].(*components.UserMessage)

// 	attackerWeapon := attacker.Components[meleeWeapon].(*components.MeleeWeapon)
// 	attackerName := attacker.Components[name].(*components.Name)
// 	attackerMessage := attacker.Components[userMessage].(*components.UserMessage)

// 	if attacker.Components[health].(*components.Health).CurrentHealth <= 0 {
// 		return
// 	}

// 	// Roll a d10 to hit
// 	toHitRoll := engine.GetDiceRoll(10)

// 	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
// 		// It's a hit
// 		damageRoll := engine.GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)

// 		damageDone := damageRoll - defenderArmor.Defense
// 		// Prevent healing the defender
// 		if damageDone < 0 {
// 			damageDone = 0
// 		}

// 		defenderHealth.CurrentHealth -= damageDone
// 		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and hits for %d health.\n", attackerName.Label, attackerWeapon.Name, defenderName.Label, damageDone)

// 		if defenderHealth.CurrentHealth <= 0 {
// 			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName.Label)
// 			if defenderName.Label == "Player" {
// 				defenderMessage.GameStateMessage = "Game Over!\n"
// 				g.Turn = scenes.GameOver
// 			}
// 		}
// 	} else {
// 		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses.\n", attackerName.Label, attackerWeapon.Name, defenderName.Label)
// 	}
// }
