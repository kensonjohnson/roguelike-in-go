package assets

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/config"
)

var (
	//go:embed "*"
	assetsFS embed.FS

	// Tiles
	Floor      *ebiten.Image
	Wall       *ebiten.Image
	StairsUp   *ebiten.Image
	StairsDown *ebiten.Image

	// UI
	UIPanel               *ebiten.Image
	UIPanelWithMinimap    *ebiten.Image
	KenneyMiniFont        *text.GoTextFace
	KenneyMiniSquaredFont *text.GoTextFace
	KenneyPixelFont       *text.GoTextFace

	// Characters
	Player *ebiten.Image
	Skelly *ebiten.Image
	Orc    *ebiten.Image
)

// Loads all required assets, panics if any one fails.
func MustLoadAssets() {
	Floor = mustLoadImage("images/tiles/floor.png")
	Wall = mustLoadImage("images/tiles/wall.png")
	StairsUp = mustLoadImage("images/tiles/stairs-up.png")
	StairsDown = mustLoadImage("images/tiles/stairs-down.png")
	UIPanel = mustLoadImage("images/ui/UIPanel.png")
	UIPanelWithMinimap = mustLoadImage("images/ui/UIPanelWithMinimap.png")
	Player = mustLoadImage("images/characters/player.png")
	Skelly = mustLoadImage("images/enemies/skelly.png")
	Orc = mustLoadImage("images/enemies/orc.png")

	kenneyMiniFontBytes, err := assetsFS.ReadFile("fonts/KenneyMini.ttf")
	if err != nil {
		log.Fatal(err)
	}
	KenneyMiniFont = mustLoadFont(kenneyMiniFontBytes)
	kenneyMiniSquaredFontBytes, err := assetsFS.ReadFile("fonts/KenneyMiniSquared.ttf")
	if err != nil {
		log.Fatal(err)
	}
	KenneyMiniSquaredFont = mustLoadFont(kenneyMiniSquaredFontBytes)
	kenneyPixelFontBytes, err := assetsFS.ReadFile("fonts/KenneyPixel.ttf")
	if err != nil {
		log.Fatal(err)
	}
	KenneyPixelFont = mustLoadFont(kenneyPixelFontBytes)
	// For some reason, the KenneyPixel shows up as half the size of the other fonts.
	KenneyPixelFont.Size = float64(config.FontSize) * 1.5
}

// Loads image at specified path, panics if it fails.
func mustLoadImage(filePath string) *ebiten.Image {
	imgSource, err := assetsFS.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	image, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imgSource))
	if err != nil {
		log.Fatal(err)
	}
	return image
}

// Loads font at specified path, panics if it fails.
func mustLoadFont(font []byte) *text.GoTextFace {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.Fatal(err)
	}
	return &text.GoTextFace{
		Source: source,
		Size:   float64(config.FontSize),
	}
}
