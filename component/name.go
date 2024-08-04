package component

import "github.com/yohamta/donburi"

type NameData struct {
	Label string
}

var Name = donburi.NewComponentType[NameData]()
