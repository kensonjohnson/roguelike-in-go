package system

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine/shapes"
	"github.com/yohamta/donburi/ecs"
)

type minimap struct {
	boxSprite *ebiten.Image
	blipSize  float32
}

var Minimap = &minimap{
	boxSprite: shapes.MakeBox(
		260, 170, 4,
		colors.Peru, color.Black,
		shapes.FancyItemCorner,
	),
	blipSize: 3.0,
}

func (m *minimap) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := tags.LevelTag.MustFirst(ecs.World)
	level := component.Level.Get(entry)

	startingXPixel := (config.ScreenWidth * config.TileWidth) - config.TileWidth - 260
	startingYPixel := config.TileHeight

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(startingXPixel), float64(startingYPixel))
	screen.DrawImage(m.boxSprite, options)

	startingXPixel += 10
	startingYPixel += 10

	// Draw the walls and floors
	for _, tile := range level.Tiles {
		x := float32(startingXPixel + (tile.TileX * int(m.blipSize)))
		y := float32(startingYPixel + (tile.TileY * int(m.blipSize)))
		if !tile.IsRevealed {
			continue
		}

		if tile.TileType == component.WALL {
			vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, colors.Peru, false)
		} else if tile.TileType == component.STAIR_DOWN {
			vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, colors.Lime, false)
		} else /* floor */ {
			vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, colors.LightGray, false)
		}
	}

	// Draw all discovered entities
	for entry = range component.Discoverable.Iter(ecs.World) {

		position := component.Position.Get(entry)
		if !level.InBounds(position.X, position.Y) {
			return
		}

		x := float32(startingXPixel + (position.X * int(m.blipSize)))
		y := float32(startingYPixel + (position.Y * int(m.blipSize)))

		if component.Discoverable.Get(entry).SeenByPlayer {
			if entry.HasComponent(tags.ItemTag) {
				vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, colors.DeepSkyBlue, false)
			} else {
				vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, colors.Red, false)
			}
		}
	}

	// Draw the player
	playerEntry := tags.PlayerTag.MustFirst(ecs.World)
	playerPosition := component.Position.Get(playerEntry)
	x := float32(startingXPixel + (playerPosition.X * int(m.blipSize)))
	y := float32(startingYPixel + (playerPosition.Y * int(m.blipSize)))
	// TODO: Pick a better color for the player
	vector.DrawFilledRect(screen, x, y, m.blipSize, m.blipSize, color.White, false)
}
