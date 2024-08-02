package components

import (
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
)

type FovData struct {
	fov *fov.View
}

var Fov = donburi.NewComponentType[FovData]()
