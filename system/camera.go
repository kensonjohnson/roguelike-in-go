package system

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi/ecs"
)

type camera struct{}

var Camera = &camera{}

func (c *camera) Update(ecs *ecs.ECS) {
	entry := tags.CameraTag.MustFirst(ecs.World)
	camera := component.Camera.Get(entry)
	entry = tags.PlayerTag.MustFirst(ecs.World)
	position := component.Position.Get(entry)

	camera.MainCamera.LookAt(
		float64((position.X)*config.TileWidth)+config.TileWidth/2,
		float64((position.Y)*config.TileHeight)+config.TileHeight/2,
	)
}
