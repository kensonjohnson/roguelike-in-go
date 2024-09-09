package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/event"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/kensonjohnson/roguelike-game-go/system/layer"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type LevelScene struct {
	ecs   ecs.ECS
	ready bool
}

func (ls *LevelScene) Update() {
	ls.ecs.Update()
	event.ProgressLevelEvent.ProcessEvents(ls.ecs.World)
}

func (ls *LevelScene) Draw(screen *ebiten.Image) {
	ls.ecs.Draw(screen)
}

func (ls *LevelScene) Ready() bool {
	return ls.ready
}

func (ls *LevelScene) Setup(world donburi.World) {
	ls.ready = false

	if logger.DebugOn {
		logger.DebugLogger.Println("LevelScene setup")
	}

	go func() {

		levelData := archetype.GenerateLevel(world)

		if _, ok := tags.UITag.First(world); !ok {
			archetype.CreateNewUI(world)
		}

		playerEntry := tags.PlayerTag.MustFirst(world)
		playerPosition := component.Position.Get(playerEntry)
		startingRoom := levelData.Rooms[0]
		playerPosition.X, playerPosition.Y = startingRoom.Center()

		playerSprite := component.Sprite.Get(playerEntry)
		playerSprite.OffestX = 0
		playerSprite.OffestY = 0

		component.Fov.Get(playerEntry).
			VisibleTiles.Compute(levelData, playerPosition.X, playerPosition.Y, 8)

		// FIX: This is a workaround to the kamera camera keeping a 'memory' of
		// previous location, even after lerp is turned off.
		archetype.ReplaceCamera(
			world,
			float64((playerPosition.X*config.TileWidth)+config.TileWidth/2),
			float64((playerPosition.Y*config.TileHeight)+config.TileHeight/2),
		)

		ls.configureECS(world)

		ls.ready = true
	}()
}

func (ls *LevelScene) Teardown() {
	ls.ready = false

	if logger.DebugOn {
		logger.DebugLogger.Println("LevelScene teardown")
	}

	go func() {
		tags.LevelTag.MustFirst(ls.ecs.World).Remove()

		for entry := range tags.MonsterTag.Iter(ls.ecs.World) {
			if logger.DebugOn {
				logger.DebugLogger.Println("Removing entry: ", entry.String())
			}
			entry.Remove()
		}

		for entry := range tags.PickupTag.Iter(ls.ecs.World) {
			if logger.DebugOn {
				logger.DebugLogger.Println("Removing entry: ", entry.String())
			}
			entry.Remove()
		}

		ls.ready = true
	}()
}

func progressLevel(world donburi.World, eventData event.ProgressLevel) {
	newLevelScene := &LevelScene{}
	SceneManager.GoTo(newLevelScene)
}

func (ls *LevelScene) configureECS(world donburi.World) {
	ls.ecs = *ecs.NewECS(world)
	// Add systems
	ls.ecs.AddSystem(system.Camera.Update)
	ls.ecs.AddSystem(system.Turn.Update)
	ls.ecs.AddSystem(system.UI.Update)

	// Add renderers
	ls.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	ls.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	ls.ecs.AddRenderer(layer.UI, system.UI.Draw)
	ls.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	if system.Debug.On {
		ls.ecs.AddRenderer(layer.UI, system.Debug.Draw)
	}

	// Add event listeners
	event.ProgressLevelEvent.Subscribe(ls.ecs.World, progressLevel)
}
