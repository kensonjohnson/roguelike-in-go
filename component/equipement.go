package component

import "github.com/yohamta/donburi"

type EquipmentData struct {
	Weapon *donburi.Entry
	Sheild *donburi.Entry
	Gloves *donburi.Entry
	Armor  *donburi.Entry
	Boots  *donburi.Entry
}

var Equipment = donburi.NewComponentType[EquipmentData]()
