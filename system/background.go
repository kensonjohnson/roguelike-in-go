package system

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi/ecs"
)

func (r *render) DrawBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := archetype.MustFindDungeon(ecs.World)
	level := component.Dungeon.Get(entry).CurrentLevel
	entry = archetype.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)
	if !level.Redraw {
		camera.MainCamera.Draw(r.backgroundImage, camera.CamImageOptions, screen)
		return
	}
	log.Println("Redrawing background")
	entry = archetype.PlayerTag.MustFirst(ecs.World)
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
