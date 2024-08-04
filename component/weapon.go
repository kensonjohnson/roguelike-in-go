package component

import "github.com/yohamta/donburi"

type WeaponData struct {
	Name          string
	ActionText    string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

var Weapon = donburi.NewComponentType[WeaponData]()
