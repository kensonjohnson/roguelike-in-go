package combat

import (
	"fmt"

	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/yohamta/donburi"
)

func AttackSystem(world donburi.World, attacker, defender *donburi.Entry) {
	attackerName := component.Name.Get(attacker)
	attackerAttackValues := component.Attack.Get(attacker)
	attackerActionText := component.ActionText.Get(attacker)
	attackerHealth := component.Health.Get(attacker)
	attackerMessages := component.UserMessage.Get(attacker)
	defenderName := component.Name.Get(defender)
	defenderDefenseValues := component.Defense.Get(defender)
	defenderHealth := component.Health.Get(defender)
	defenderMessages := component.UserMessage.Get(defender)

	if attackerHealth.CurrentHealth <= 0 {
		return
	}

	toHitRoll := engine.GetDiceRoll(10)

	if toHitRoll+attackerAttackValues.ToHitBonus > defenderDefenseValues.ArmorClass {
		// Its a hit!
		damageRoll := engine.GetRandomBetween(attackerAttackValues.MinimumDamage, attackerAttackValues.MaximumDamage)
		damageDone := damageRoll - defenderDefenseValues.Defense
		// Prevent healing the defender
		if damageDone < 0 {
			damageDone = 0
		}

		defenderHealth.CurrentHealth -= damageDone
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and deals %d dmg.\n", attackerName.Value, attackerActionText.Value, defenderName.Value, damageDone)
		if defenderHealth.CurrentHealth <= 0 {
			defenderMessages.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName.Value)
		}
		entry := archetype.CameraTag.MustFirst(world)
		camera := component.Camera.Get(entry)
		camera.MainCamera.AddTrauma(0.2)
	} else {
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and misses.\n", attackerName.Value, attackerActionText.Value, defenderName.Value)
	}
}
