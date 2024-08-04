package component

import "github.com/yohamta/donburi"

type HealthData struct {
	MaxHealth     int
	CurrentHealth int
}

var Health = donburi.NewComponentType[HealthData]()
