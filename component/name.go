package component

import "github.com/yohamta/donburi"

type NameData struct {
	Value string
}

var Name = donburi.NewComponentType[NameData]()
