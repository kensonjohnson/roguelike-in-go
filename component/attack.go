package component

import "github.com/yohamta/donburi"

type AttackData struct {
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

var Attack = donburi.NewComponentType[AttackData]()
