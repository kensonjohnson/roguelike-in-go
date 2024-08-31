package component

import "github.com/yohamta/donburi"

type ActionTextData struct {
	Value string
}

var ActionText = donburi.NewComponentType[ActionTextData]()
