package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/config"
)

var (
	transitionFrom = ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight)
	transitionTo   = ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight)
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

const transitionMaxCount = 30

type SceneManagerData struct {
	current         Scene
	next            Scene
	transitionCount int
}

var SceneManager = &SceneManagerData{}

func (s *SceneManagerData) Update() {
	s.current.Update()
	if s.transitionCount == 0 {
		return
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return
	}

	s.current = s.next
	s.next = nil

}

func (s *SceneManagerData) Draw(screen *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(screen)
		return
	}

	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	screen.DrawImage(transitionFrom, nil)

	alpha := 1 - float32(s.transitionCount)/float32(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)
	screen.DrawImage(transitionTo, op)
}

func (s *SceneManagerData) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
