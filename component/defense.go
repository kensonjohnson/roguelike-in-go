package component

import "github.com/yohamta/donburi"

type DefenseData struct {
	Defense    int
	ArmorClass int
}

var Defense = donburi.NewComponentType[DefenseData]()
