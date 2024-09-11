package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/kensonjohnson/roguelike-game-go/system/scene"
)

type Game struct {
	sceneManager *scene.SceneManagerData
}

func (g *Game) configure() {
	g.sceneManager = scene.SceneManager
	g.sceneManager.Setup()
	g.sceneManager.GoTo(&scene.TitleScene{
		ImageBackground: assets.Floor,
		PixelWidth:      config.ScreenWidth * config.TileWidth,
		PixelHeight:     config.ScreenHeight * config.TileHeight,
	})

}

func (g *Game) Update() error {
	g.sceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

// Returns the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return config.TileWidth * config.ScreenWidth, config.TileHeight * config.ScreenHeight
}

func main() {
	DebugOn := flag.Bool("debug", false, "Enable debug screens and prints")
	flag.Parse()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Roguelike")
	if DebugOn != nil && *DebugOn {
		ebiten.SetVsyncEnabled(false)
		system.Debug.On = true
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	log.SetFlags(log.Lshortfile)

	g := &Game{}
	g.configure()

	slog.Debug("Starting Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Panic(err)
	}
}
