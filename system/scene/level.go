package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/layer"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Level struct {
	ecs ecs.ECS
}

func (level *Level) Update() {
	level.ecs.Update()
}

func (level *Level) Draw(screen *ebiten.Image) {
	level.ecs.Draw(screen)
}

func NewLevel() *Level {
	level := &Level{}
	level.ecs = *ecs.NewECS(createWorld())

	level.ecs.AddSystem(system.Camera.Update)
	level.ecs.AddSystem(system.Turn.Update)
	level.ecs.AddSystem(system.UI.Update)
	level.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	level.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	level.ecs.AddRenderer(layer.UI, system.UI.Draw)
	level.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	// level.ecs.AddRenderer(layer.UI, system.Debug.Draw)

	return level
}

func createWorld() donburi.World {
	world := donburi.NewWorld()

	// Create current level
	level := archetype.GenerateLevel(world)

	for index, room := range level.Rooms {
		if index == 0 {
			archetype.CreateNewPlayer(world)
		} else {
			archetype.CreateMonster(world, level, room)
		}
	}

	// Create the UI
	archetype.CreateNewUI(world)

	// Create the camera
	archetype.CreateNewCamera(world)

	return world
}
