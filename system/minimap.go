package system

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

const blipSize = 4

func DrawMinimap(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := archetype.MustFindDungeon(ecs.World)
	level := component.Dungeon.Get(entry).CurrentLevel

	startingXPixel := (config.ScreenWidth * config.TileWidth) - 340
	startingYPixel := (config.ScreenHeight * config.TileWidth) - 210

	// Draw the walls and floors
	for _, tile := range level.Tiles {
		x := startingXPixel + (tile.TileX * blipSize)
		y := startingYPixel + (tile.TileY * blipSize)
		if !tile.IsRevealed {
			continue
		}

		if tile.TileType == component.WALL {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, color.RGBA{R: 202, G: 146, B: 74, A: 255}, false)
		} else {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, color.RGBA{R: 178, G: 182, B: 194, A: 255}, false)
		}
	}

	// Draw all discovered entities
	component.Discoverable.Each(ecs.World, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		if !level.InBounds(position.X, position.Y) {
			return
		}

		x := startingXPixel + (position.X * blipSize)
		y := startingYPixel + (position.Y * blipSize)

		if component.Discoverable.Get(entry).SeenByPlayer {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
		}
	})

	// Draw the player
	playerEntry := archetype.PlayerTag.MustFirst(ecs.World)
	playerPosition := component.Position.Get(playerEntry)
	x := startingXPixel + (playerPosition.X * blipSize)
	y := startingYPixel + (playerPosition.Y * blipSize)
	vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, color.White, false)
}
