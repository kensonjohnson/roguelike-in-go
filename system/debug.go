package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi/ecs"
	donburiDebug "github.com/yohamta/donburi/features/debug"
)

type debug struct {
	On            bool
	frameCount    int
	totalEntities int
	monsterCount  int
	itemCount     int
	pickupCount   int
	miscCount     int
}

var Debug = &debug{
	On:            false,
	frameCount:    0,
	totalEntities: 0,
	monsterCount:  0,
	itemCount:     0,
	pickupCount:   0,
	miscCount:     0,
}

func (d *debug) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	width := config.ScreenWidth * config.TileWidth
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %v\nFPS: %v", int(ebiten.ActualTPS()), int(ebiten.ActualFPS())), width-60, 250)

	// Draw with known information
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entities: %v", d.totalEntities), 8, 180)
	offset := 0
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v %v", "Monster: ", d.monsterCount),
		8, 194+(offset*14),
	)
	offset++

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v %v", "Total Items: ", d.itemCount),
		8, 194+(offset*14),
	)
	offset++

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v %v", "Pickups: ", d.pickupCount),
		8, 194+(offset*14),
	)
	offset++

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v %v", "Uncategorized: ", d.miscCount),
		8, 194+(offset*14),
	)

	d.frameCount++
	if d.frameCount < 60 {
		return
	}
	d.frameCount = 0

	// Recalculate all numbers
	d.totalEntities = ecs.World.Len()
	d.monsterCount = 0
	d.itemCount = 0
	d.pickupCount = 0
	d.miscCount = 0
	archetypes := donburiDebug.GetEntityCounts(ecs.World)
	for _, arch := range archetypes {
		// List out the entities that you care to record
		if arch.Archetype.Layout().HasComponent(tags.MonsterTag) {
			d.monsterCount += arch.Count
		} else if arch.Archetype.Layout().HasComponent(tags.PickupTag) {
			d.pickupCount += arch.Count
			d.itemCount += arch.Count
		} else if arch.Archetype.Layout().HasComponent(tags.ItemTag) {
			d.itemCount += arch.Count
		} else {
			d.miscCount += arch.Count
		}
	}
}
