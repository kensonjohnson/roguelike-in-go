package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/system/scene"
)

type Game struct {
	sceneManager *scene.SceneManagerData
}

func (g *Game) configure() {
	g.sceneManager = scene.SceneManager
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
	assets.MustLoadAssets()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Roguelike")

	g := &Game{}
	g.configure()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
