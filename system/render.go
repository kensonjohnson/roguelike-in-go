package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type render struct {
	query           *donburi.Query
	backgroundImage *ebiten.Image
}

var Render = &render{
	query: donburi.NewQuery(
		filter.Contains(
			component.Position,
			component.Sprite,
		)),
	backgroundImage: ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight),
}

func (r *render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := archetype.PlayerTag.MustFirst(ecs.World)
	playerVision := component.Fov.Get(entry).VisibleTiles
	entry = archetype.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)

	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		sprite := component.Sprite.Get(entry)

		if playerVision.IsVisible(position.X, position.Y) {
			camera.CamImageOptions.GeoM.Reset()
			if sprite.Animating {
				offsetX, offsetY := sprite.GetAnimationStep()
				camera.CamImageOptions.GeoM.Translate(float64(position.X*config.TileWidth)+offsetX, float64(position.Y*config.TileHeight)+offsetY)
			} else {
				camera.CamImageOptions.GeoM.Translate(float64(position.X*config.TileWidth), float64(position.Y*config.TileHeight))
			}
			camera.MainCamera.Draw(sprite.Image, camera.CamImageOptions, screen)
		}
	})

	camera.CamImageOptions.GeoM.Reset()
	camera.CamImageOptions.GeoM.Translate(0, 0)
	camera.MainCamera.Draw(camera.CamScreen, camera.CamImageOptions, screen)
}
