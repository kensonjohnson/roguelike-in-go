package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/config"
)

var (
	transitionFrom = ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	transitionTo   = ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

const transitionMaxCount = 20

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
}

func (s *SceneManager) Update() {
	if s.transitionCount == 0 {
		s.current.Update()
		return
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return
	}

	s.current = s.next
	s.next = nil

}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(screen)
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	screen.DrawImage(transitionFrom, nil)

	alpha := 1 - float32(s.transitionCount)/float32(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)
	screen.DrawImage(transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
