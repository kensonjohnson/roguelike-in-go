package assets

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	//go:embed "*"
	assetsFS embed.FS

	// Tiles
	Floor *ebiten.Image
	Wall  *ebiten.Image

	// UI
	UIPanel *ebiten.Image

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
