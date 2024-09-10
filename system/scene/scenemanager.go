package scene

import (
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
	Setup(world donburi.World)
	Teardown()
	Ready() bool
}

const transitionMaxCount = 30

type SceneManagerData struct {
	current                Scene
	next                   Scene
	fadeOut                bool
	transitionCount        int
	transitionTo           *ebiten.Image
	transitionFrom         *ebiten.Image
	transitionImagesCached bool
	world                  donburi.World
}

var SceneManager = &SceneManagerData{
	transitionTo:   ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight),
	transitionFrom: ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight),
}

func (sm *SceneManagerData) Setup() {
	slog.Debug("SceneManager Setup")
	sm.world = donburi.NewWorld()
	archetype.CreateNewPlayer(sm.world, items.Weapons.BattleAxe, items.Armor.PlateArmor)
	archetype.CreateNewCamera(sm.world)
}

func (sm *SceneManagerData) Update() {
	if sm.next == nil && sm.transitionCount == 0 {
		sm.current.Update()
		return
	}

	sm.transitionCount--
	if sm.transitionCount > 0 {
		return
	}

	if sm.fadeOut {
		slog.Debug("Running Teardown on current scene")
		sm.fadeOut = false
		sm.current.Teardown()
		return
	}

	if sm.transitionCount != 0 && sm.next != nil && sm.current.Ready() {
		slog.Debug("Running Setup on next scene")
		sm.current = sm.next
		sm.next = nil
		sm.current.Setup(sm.world)
		return
	}

	if sm.transitionCount != 0 && sm.next == nil && sm.current.Ready() {
		sm.transitionCount = transitionMaxCount
		sm.transitionImagesCached = false
	}

}

func (sm *SceneManagerData) Draw(screen *ebiten.Image) {

	if sm.next == nil && sm.transitionCount == 0 {
		sm.current.Draw(screen)
		return
	}

	if sm.fadeOut && !sm.transitionImagesCached {
		sm.current.Draw(sm.transitionFrom)
		sm.transitionTo.Fill(color.Black)
		sm.transitionImagesCached = true
	}

	if !sm.fadeOut && !sm.transitionImagesCached {
		sm.current.Draw(sm.transitionTo)
		sm.transitionFrom.Fill(color.Black)
		sm.transitionImagesCached = true
	}

	screen.DrawImage(sm.transitionFrom, nil)
	alpha := 1 - float32(sm.transitionCount)/float32(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)
	screen.DrawImage(sm.transitionTo, op)
}

func (sm *SceneManagerData) GoTo(scene Scene) {
	if sm.current == nil {
		sm.current = scene
	} else {
		sm.next = scene
		sm.fadeOut = true
		sm.transitionCount = transitionMaxCount
		sm.transitionImagesCached = false
	}
}
