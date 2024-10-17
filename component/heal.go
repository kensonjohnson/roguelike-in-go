package component

import "github.com/yohamta/donburi"

type HealData struct {
	Amount int
}

var Heal = donburi.NewComponentType[HealData]()
