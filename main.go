package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/layer"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Game struct {
	ecs ecs.ECS
}

func (g *Game) configure() {
	g.ecs = *ecs.NewECS(createWorld())
	g.ecs.AddSystem(system.Camera.Update)
	g.ecs.AddSystem(system.Turn.Update)
	g.ecs.AddSystem(system.UI.Update)
	g.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	g.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	g.ecs.AddRenderer(layer.UI, system.UI.Draw)
	g.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	g.ecs.AddRenderer(layer.UI, system.Debug.Draw)
}

func (g *Game) Update() error {
	g.ecs.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.ecs.DrawLayer(layer.Background, screen)
	g.ecs.DrawLayer(layer.Foreground, screen)
	g.ecs.DrawLayer(layer.UI, screen)
}

// Returns the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return config.TileWidth * config.ScreenWidth, config.TileHeight * config.ScreenHeight
}

func createWorld() donburi.World {
	world := donburi.NewWorld()

	// Create dungeon component
	dungeon := archetype.GenerateDungeon(world)

	for index, room := range dungeon.CurrentLevel.Rooms {
		if index == 0 {
			archetype.CreateNewPlayer(world)
		} else {
			archetype.CreateMonster(world, dungeon.CurrentLevel, room)
		}
	}

	// Create the UI
	archetype.CreateNewUI(world)

	// Create the camera
	archetype.CreateNewCamera(world)

	return world
}

func main() {
	assets.MustLoadAssets()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Roguelike")
	// ebiten.SetVsyncEnabled(false)

	g := &Game{}
	g.configure()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
