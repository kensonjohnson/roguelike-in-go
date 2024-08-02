package main

import (
	_ "image/png"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/scenes"
)

// Holds all data the entire game will need.
type Game struct {
	Map         scenes.GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        scenes.TurnState
	TurnCounter int
}

// Creates a new Game Object and initializes the data.
func NewGame() *Game {
	g := &Game{}
	g.Map = scenes.NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)

	g.WorldTags = tags
	g.World = world
	g.Turn = scenes.PlayerTurn
	g.TurnCounter = 0

	return g
}

// Called each tick (game loop).
func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == scenes.PlayerTurn && g.TurnCounter > 8 {
		TakePlayerAction(g)
	}
	if g.Turn == scenes.MonsterTurn {
		UpdateMonster(g)
	}

	return nil
}

// Called each draw cycle in the game loop.
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
	ProcessUserLog(g, screen)
	ProcessHUD(g, screen)
}

// Returns the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return config.Config.TileWidth * config.Config.ScreenWidth, config.Config.TileHeight * config.Config.ScreenHeight
}

func main() {
	assets.MustLoadAssets()

	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Roguelike")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
