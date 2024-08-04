package component

import (
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
)

type FovData struct {
	VisibleTiles *fov.View
}

var Fov = donburi.NewComponentType[FovData]()
