package components

import "github.com/yohamta/donburi"

type PlayerData struct{}

var Player = donburi.NewComponentType[PlayerData]()
