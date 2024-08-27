package items

import "github.com/kensonjohnson/roguelike-game-go/component"

type weapons struct {
	Fist       component.WeaponData
	ShortSword component.WeaponData
	Machete    component.WeaponData
	BattleAxe  component.WeaponData
}

var Weapons weapons

func init() {
	Weapons = weapons{
		Fist: component.WeaponData{
			Name:          "Fist",
			ActionText:    "punches",
			MinimumDamage: 1,
			MaximumDamage: 3,
			ToHitBonus:    2,
		},

		ShortSword: component.WeaponData{
			Name:          "Short Sword",
			ActionText:    "swings a short sword at",
			MinimumDamage: 2,
			MaximumDamage: 6,
			ToHitBonus:    0,
		},

		Machete: component.WeaponData{
			Name:          "Machete",
			ActionText:    "swings a machete at",
			MinimumDamage: 4,
			MaximumDamage: 8,
			ToHitBonus:    1,
		},

		BattleAxe: component.WeaponData{
			Name:          "Battle Axe",
			ActionText:    "cleaves a battle axe at",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		},
	}
}
