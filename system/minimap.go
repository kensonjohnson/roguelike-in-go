package system

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi/ecs"
)

const blipSize = 4

func DrawMinimap(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := tags.LevelTag.MustFirst(ecs.World)
	level := component.Level.Get(entry)

	// The values of 330 and 210 are based on the size of the minimap image.
	// That image is 340x220 pixels, with a 10 pixel border, and is placed
	// in the bottom right corner of the screen.
	startingXPixel := (config.ScreenWidth * config.TileWidth) - 330
	startingYPixel := (config.ScreenHeight * config.TileWidth) - 210

	// Draw the walls and floors
	for _, tile := range level.Tiles {
		x := startingXPixel + (tile.TileX * blipSize)
		y := startingYPixel + (tile.TileY * blipSize)
		if !tile.IsRevealed {
			continue
		}

		if tile.TileType == component.WALL {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, colors.Peru, false)
		} else if tile.TileType == component.STAIR_DOWN {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, colors.Lime, false)
		} else /* floor */ {
			vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, colors.LightGray, false)
		}
	}

	// Draw all discovered entities
	for entry = range component.Discoverable.Iter(ecs.World) {

		position := component.Position.Get(entry)
		if !level.InBounds(position.X, position.Y) {
			return
		}

		x := startingXPixel + (position.X * blipSize)
		y := startingYPixel + (position.Y * blipSize)

		if component.Discoverable.Get(entry).SeenByPlayer {
			if entry.HasComponent(tags.ItemTag) {
				vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, colors.DeepSkyBlue, false)
			} else {
				vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, colors.Red, false)
			}
		}
	}

	// Draw the player
	playerEntry := tags.PlayerTag.MustFirst(ecs.World)
	playerPosition := component.Position.Get(playerEntry)
	x := startingXPixel + (playerPosition.X * blipSize)
	y := startingYPixel + (playerPosition.Y * blipSize)
	vector.DrawFilledRect(screen, float32(x), float32(y), blipSize, blipSize, color.White, false)
}
