package component

import "github.com/yohamta/donburi"

type ValueData struct {
	Amount int
}

var Value = donburi.NewComponentType[ValueData]()
