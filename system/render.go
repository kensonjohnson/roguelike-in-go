package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type render struct {
	enemyQuery      *donburi.Query
	pickupsQuery    *donburi.Query
	backgroundImage *ebiten.Image
}

var Render = &render{
	enemyQuery: donburi.NewQuery(
		filter.Contains(
			tags.MonsterTag,
			component.Sprite,
			component.Position,
		)),
	pickupsQuery: donburi.NewQuery(
		filter.Contains(
			tags.ItemTag,
			component.Position,
			component.Sprite,
		)),
	backgroundImage: ebiten.NewImage(config.ScreenWidth*config.TileWidth, config.ScreenHeight*config.TileHeight),
}

func (r *render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	playerEntry := tags.PlayerTag.MustFirst(ecs.World)
	playerVision := component.Fov.Get(playerEntry).VisibleTiles
	entry := tags.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)

	for entry = range r.pickupsQuery.Iter(ecs.World) {
		renderEntity(screen, camera, entry, playerVision)
	}

	for entry = range r.enemyQuery.Iter(ecs.World) {
		renderEntity(screen, camera, entry, playerVision)
	}

	renderEntity(screen, camera, playerEntry, playerVision)

	camera.CamImageOptions.GeoM.Reset()
	camera.CamImageOptions.GeoM.Translate(0, 0)
	camera.MainCamera.Draw(camera.CamScreen, camera.CamImageOptions, screen)
}

func renderEntity(
	screen *ebiten.Image,
	camera *component.CameraData,
	entry *donburi.Entry,
	playerVision *fov.View,
) {
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
}
