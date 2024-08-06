package component

import "github.com/yohamta/donburi"

type DiscoverableData struct {
	SeenByPlayer bool
}

var Discoverable = donburi.NewComponentType[DiscoverableData]()
