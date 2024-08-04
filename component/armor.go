package component

import "github.com/yohamta/donburi"

type ArmorData struct {
	Name       string
	Defense    int
	ArmorClass int
}

var Armor = donburi.NewComponentType[ArmorData]()
