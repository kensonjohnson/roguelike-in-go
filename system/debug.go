package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi/ecs"
)

const Controls = `
| Key     | Action         |
| ------- | -------------- |
| WASD    | Move character |
| ← ↑ → ↓ | Move character |
| Q       | Pass Turn      |
`

type debug struct{}

var Debug = &debug{}

func (d *debug) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	width := config.ScreenWidth * config.TileWidth
	ebitenutil.DebugPrintAt(screen, Controls, 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %v\nFPS: %v", int(ebiten.ActualTPS()), int(ebiten.ActualFPS())), width-60, 10)
}
