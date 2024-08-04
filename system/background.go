package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi/ecs"
)

func DrawBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := archetype.MustFindDungeon(ecs.World)
	level := component.Dungeon.Get(entry).CurrentLevel
	entry = archetype.PlayerTag.MustFirst(ecs.World)
	playerVision := component.Fov.Get(entry).VisibleTiles

	maxTiles := config.ScreenWidth * (config.ScreenHeight - config.UIHeight)
	for i := 0; i < maxTiles; i++ {
		tile := level.Tiles[i]
		isVisible := playerVision.IsVisible(tile.TileX, tile.TileY)
		if isVisible {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
			tile.IsRevealed = true
		} else if tile.IsRevealed {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			options.ColorScale.ScaleAlpha(0.35)
			screen.DrawImage(tile.Image, options)
		}
	}
}
