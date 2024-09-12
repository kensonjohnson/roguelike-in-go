package component

import (
	"github.com/yohamta/donburi"
)

type UIData struct {
	MessageBox UserMessageBoxData
	PlayerHUD  PlayerHUDData
}

type UserMessageBoxData struct {
	Position PositionData
	FontX    int
	FontY    int
}

type PlayerHUDData struct {
	Position PositionData
	FontX    int
	FontY    int
	Health   *HealthData
	Attack   *AttackData
	Defense  *DefenseData
}

var UI = donburi.NewComponentType[UIData]()
