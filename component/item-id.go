package component

import "github.com/yohamta/donburi"

type ItemIdData struct {
	Id int
}

var ItemId = donburi.NewComponentType[ItemIdData]()
