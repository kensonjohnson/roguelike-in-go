package component

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	Sprite   *ebiten.Image
}

type PlayerHUDData struct {
	Position PositionData
	FontX    int
	FontY    int
	Health   *HealthData
	Armor    *ArmorData
	Weapon   *WeaponData
	Sprite   *ebiten.Image
}

var UI = donburi.NewComponentType[UIData]()
