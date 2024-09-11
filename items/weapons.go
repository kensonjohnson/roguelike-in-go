package items

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type WeaponData struct {
	ItemData
	ActionText    string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

type weapons struct {
	ShortSword WeaponData
	BattleAxe  WeaponData
}

var Weapons = weapons{
	ShortSword: WeaponData{
		ItemData: ItemData{
			Name:   "Short Sword",
			Sprite: assets.MustBeValidImage(assets.ShortSword, "ShortSword"),
		},
		ActionText:    "swings a short sword at",
		MinimumDamage: 2,
		MaximumDamage: 6,
		ToHitBonus:    0,
	},
	BattleAxe: WeaponData{
		ItemData: ItemData{
			Name:   "Battle Axe",
			Sprite: assets.MustBeValidImage(assets.BattleAxe, "BattleAxe"),
		},
		ActionText:    "cleaves a battle axe at",
		MinimumDamage: 10,
		MaximumDamage: 20,
		ToHitBonus:    3,
	},
}
