package assets

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/fonts"
	"github.com/kensonjohnson/roguelike-game-go/globals"
)

var (
	//go:embed "*"
	assetsFS embed.FS

	// Tiles
	Floor *ebiten.Image
	Wall  *ebiten.Image

	// UI
	UIPanel *ebiten.Image
	HUDFont *text.GoTextFace

	// Characters
	Player *ebiten.Image
	Skelly *ebiten.Image
	Orc    *ebiten.Image
)

func MustLoadAssets() {
	Floor = MustLoadImage("floor.png")
	Wall = MustLoadImage("wall.png")
	UIPanel = MustLoadImage("UIPanel.png")
	Player = MustLoadImage("player.png")
	Skelly = MustLoadImage("skelly.png")
	Orc = MustLoadImage("orc.png")

	HUDFont = MustLoadFont(fonts.MPlus1pRegular_ttf)
}

func MustLoadImage(filePath string) *ebiten.Image {
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

func MustLoadFont(font []byte) *text.GoTextFace {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.Fatal(err)
	}
	return &text.GoTextFace{
		Source: source,
		Size:   globals.FONT_SIZE,
	}
}
