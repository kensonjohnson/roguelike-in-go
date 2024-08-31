package component

import "github.com/yohamta/donburi"

type HealData struct {
	HealAmount int
}

var Heal = donburi.NewComponentType[HealData]()
