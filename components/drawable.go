package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type DrawableData struct {
	Image *ebiten.Image
}

var Drawable = donburi.NewComponentType[DrawableData]()
