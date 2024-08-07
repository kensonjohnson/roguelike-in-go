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
	attackerWeapon := component.Weapon.Get(attacker)
	attackerHealth := component.Health.Get(attacker)
	attackerMessages := component.UserMessage.Get(attacker)
	defenderName := component.Name.Get(defender)
	defenderArmor := component.Armor.Get(defender)
	defenderHealth := component.Health.Get(defender)
	defenderMessages := component.UserMessage.Get(defender)

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
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and deals %d dmg.\n", attackerName.Label, attackerWeapon.ActionText, defenderName.Label, damageDone)
		if defenderHealth.CurrentHealth <= 0 {
			defenderMessages.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName.Label)
		}
		entry := archetype.CameraTag.MustFirst(world)
		camera := component.Camera.Get(entry)
		camera.MainCamera.AddTrauma(0.2)
	} else {
		attackerMessages.AttackMessage = fmt.Sprintf("%s %s %s and misses.\n", attackerName.Label, attackerWeapon.ActionText, defenderName.Label)
	}
}
