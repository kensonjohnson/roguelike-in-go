package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi/ecs"
	donburiDebug "github.com/yohamta/donburi/features/debug"
)

type debug struct {
	On bool
}

var Debug = &debug{
	On: false,
}

func (d *debug) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	width := config.ScreenWidth * config.TileWidth
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %v\nFPS: %v", int(ebiten.ActualTPS()), int(ebiten.ActualFPS())), width-60, 10)

	archetypes := donburiDebug.GetEntityCounts(ecs.World)
	allEntities := 0
	for _, system := range archetypes {
		allEntities += system.Count
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entities: %v", allEntities), 8, 6)

	for i, system := range archetypes {
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("%v Entities: %v", system.Archetype.Layout(), system.Count),
			8, 20+(i*14),
		)
	}

}
