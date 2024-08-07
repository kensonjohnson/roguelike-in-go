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
	Floor *ebiten.Image
	Wall  *ebiten.Image

	// UI
	UIPanel               *ebiten.Image
	UIPanelWithMinimap    *ebiten.Image
	HUDFont               *text.GoTextFace
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
	Floor = mustLoadImage("floor.png")
	Wall = mustLoadImage("wall.png")
	UIPanel = mustLoadImage("UIPanel.png")
	UIPanelWithMinimap = mustLoadImage("UIPanelWithMinimap.png")
	Player = mustLoadImage("player.png")
	Skelly = mustLoadImage("skelly.png")
	Orc = mustLoadImage("orc.png")

	HUDFont = mustLoadFont(MPlus1pRegular_ttf)
	kenneyMiniFontBytes, err := assetsFS.ReadFile("KenneyMini.ttf")
	if err != nil {
		log.Fatal(err)
	}
	KenneyMiniFont = mustLoadFont(kenneyMiniFontBytes)
	kenneyMiniSquaredFontBytes, err := assetsFS.ReadFile("KenneyMiniSquared.ttf")
	if err != nil {
		log.Fatal(err)
	}
	KenneyMiniSquaredFont = mustLoadFont(kenneyMiniSquaredFontBytes)
	kenneyPixelFontBytes, err := assetsFS.ReadFile("KenneyPixel.ttf")
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
