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
	query *donburi.Query
}

var Render = &render{
	query: donburi.NewQuery(
		filter.Contains(
			component.Position,
			component.Sprite,
		)),
}

func (r *render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	entry := archetype.PlayerTag.MustFirst(ecs.World)
	playerVision := component.Fov.Get(entry).VisibleTiles

	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		sprite := component.Sprite.Get(entry)

		if playerVision.IsVisible(position.X, position.Y) {
			drawOptions := &ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(position.X*config.TileWidth), float64(position.Y*config.TileHeight))
			screen.DrawImage(sprite.Image, drawOptions)
		}
	})
}
