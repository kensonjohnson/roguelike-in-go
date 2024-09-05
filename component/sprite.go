package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image          *ebiten.Image
	OffestX        int // The direction the sprite is moving FROM, -1, 0, 1
	OffestY        int // The direction the sprite is moving FROM, -1, 0, 1
	AnimationFrame int
	progress       float64 // The progress of the current animation frame from 0 to 1
	Animating      bool
}

// Returns the X and Y offset for the current animation frame. TODO: Add
// support for animation frames.
func (sd *SpriteData) GetAnimationStep() (float64, float64) {
	offsetX := (1 - sd.progress) * config.TileWidth * float64(sd.OffestX)
	offsetY := (1 - sd.progress) * config.TileHeight * float64(sd.OffestY)
	return offsetX, offsetY
}

func (sd *SpriteData) SetProgress(progress float64) {
	sd.progress = progress
}

var Sprite = donburi.NewComponentType[SpriteData]()
