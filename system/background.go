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
	entry = archetype.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)

	maxTiles := config.ScreenWidth * (config.ScreenHeight - config.UIHeight)
	for i := 0; i < maxTiles; i++ {
		tile := level.Tiles[i]
		isVisible := playerVision.IsVisible(tile.TileX, tile.TileY)
		camera.CamImageOptions.GeoM.Reset()
		camera.CamImageOptions.ColorScale.Reset()
		if isVisible {
			camera.CamImageOptions.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			camera.MainCamera.Draw(tile.Image, camera.CamImageOptions, screen)
			tile.IsRevealed = true
		} else if tile.IsRevealed {
			camera.CamImageOptions.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			camera.CamImageOptions.ColorScale.ScaleAlpha(0.35)
			camera.MainCamera.Draw(tile.Image, camera.CamImageOptions, screen)
		}
	}
	camera.CamImageOptions.GeoM.Reset()
	camera.CamImageOptions.ColorScale.Reset()
}
