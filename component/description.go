package component

import "github.com/yohamta/donburi"

type DescriptionData struct {
	Value string
}

var Description = donburi.NewComponentType[DescriptionData]()
