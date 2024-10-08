package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi/ecs"
)

func (r *render) DrawBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := tags.LevelTag.MustFirst(ecs.World)
	level := component.Level.Get(entry)
	entry = tags.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)
	if !level.Redraw {
		camera.MainCamera.Draw(r.backgroundImage, camera.CamImageOptions, screen)
		return
	}
	r.backgroundImage.Clear()
	entry = tags.PlayerTag.MustFirst(ecs.World)
	playerVision := component.Fov.Get(entry).VisibleTiles

	maxTiles := config.ScreenWidth * (config.ScreenHeight - config.UIHeight)
	for i := 0; i < maxTiles; i++ {
		tile := level.Tiles[i]
		isVisible := playerVision.IsVisible(tile.TileX, tile.TileY)
		options := ebiten.DrawImageOptions{}
		if isVisible {
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			r.backgroundImage.DrawImage(tile.Image, &options)
			tile.IsRevealed = true
		} else if tile.IsRevealed {
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			options.ColorScale.ScaleAlpha(0.35)
			r.backgroundImage.DrawImage(tile.Image, &options)
		}
	}
	camera.CamImageOptions.GeoM.Reset()
	camera.CamImageOptions.ColorScale.Reset()
	camera.MainCamera.Draw(r.backgroundImage, camera.CamImageOptions, screen)
	level.Redraw = false
}
