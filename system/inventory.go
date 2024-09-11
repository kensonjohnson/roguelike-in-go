package system

import (
	"image/color"
	"log/slog"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi/ecs"
)

type inventoryUi struct {
	open bool
}

var InventoryUI = inventoryUi{open: false}

func (i *inventoryUi) Update(ecs *ecs.ECS) {
	if !i.open {
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyI) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		slog.Debug("Close Inventory")
		Turn.TurnState = PlayerTurn
		i.open = false
	}
}

func (i *inventoryUi) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if !i.open {
		return
	}
	image := makeBox(
		(15)*config.TileWidth,
		(8)*config.TileHeight,
		colors.SlateGray,
		colors.Black,
	)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Reset()
	options.GeoM.Translate(float64(5*config.TileWidth), float64(5*config.TileHeight))
	screen.DrawImage(image, options)

	image = makeBox(
		(15)*config.TileWidth,
		(5)*config.TileHeight,
		colors.White,
		colors.Transparent,
	)
	options.GeoM.Reset()
	options.GeoM.Translate(float64(5*config.TileWidth), float64(15*config.TileHeight))
	screen.DrawImage(image, options)

	image = makeBox(
		(15)*config.TileWidth,
		(8)*config.TileHeight,
		colors.Peru,
		colors.SlateGray,
	)
	options.GeoM.Reset()
	options.GeoM.Translate(float64(5*config.TileWidth), float64(25*config.TileHeight))
	screen.DrawImage(image, options)
}

func (i *inventoryUi) Open() {
	i.open = true
}
func (i *inventoryUi) Close() {
	i.open = false
}

// TODO: Refactor to a proper place, probably internal
// TODO: Allow for corner variants
// w and h are pixel values
func makeBox(w, h int, border, fill color.RGBA) *ebiten.Image {

	image := ebiten.NewImage(w, h)
	corner := makeCornerImage(border, fill)
	size := corner.Bounds().Size()
	options := &ebiten.DrawImageOptions{}

	// NW
	image.DrawImage(corner, options)

	// NE
	options.GeoM.Rotate(degreesToRadians(90))
	options.GeoM.Translate(float64(w), 0)
	image.DrawImage(corner, options)

	// SE
	options.GeoM.Reset()
	options.GeoM.Rotate(degreesToRadians(180))
	options.GeoM.Translate(float64(w), float64(h))
	image.DrawImage(corner, options)

	// SW
	options.GeoM.Reset()
	options.GeoM.Rotate(degreesToRadians(270))
	options.GeoM.Translate(0, float64(h))
	image.DrawImage(corner, options)

	// Draw top and bottom lines, plus fill
	line := ebiten.NewImage(w-(size.X*2), 1)
	line.Fill(border)
	for i := 0; i < size.Y; i++ {
		if i == 4 {
			line.Fill(fill)
		}
		options.GeoM.Reset()
		options.GeoM.Translate(float64(size.X), float64(i))
		image.DrawImage(line, options)
		options.GeoM.Translate(0, float64(h-(i*2)-1))
		image.DrawImage(line, options)
	}

	// Draw vertical lines and fill
	line = ebiten.NewImage(w, h-(size.Y*2))
	line.Fill(fill)
	ends := ebiten.NewImage(4, h-(size.Y*2))
	ends.Fill(border)
	options.GeoM.Reset()
	options.GeoM.Translate(float64(w-4), 0)
	line.DrawImage(ends, options)
	options.GeoM.Reset()
	line.DrawImage(ends, options)
	options.GeoM.Translate(0, float64(size.Y))
	image.DrawImage(line, options)

	return image
}

func makeCornerImage(border, fill color.RGBA) *ebiten.Image {
	image := ebiten.NewImage(32, 32)
	block := ebiten.NewImage(4, 4)
	options := &ebiten.DrawImageOptions{}
	for y, row := range cornerShape {
		for x, value := range row {
			options.GeoM.Reset()
			options.GeoM.Translate(float64(x*4), float64(y*4))
			switch value {
			case -1:
				continue
			case 0:
				block.Fill(border)
			case 1:
				block.Fill(fill)
			}
			image.DrawImage(block, options)
		}
	}

	return image
}

func degreesToRadians(degrees int) float64 {
	return float64(degrees) * math.Pi / 180
}

var cornerShape = [][]int8{
	{0, 0, 0, 0, 0, -1, 0, 0},
	{0, 1, 1, 1, 0, -1, 0, 1},
	{0, 1, 1, 1, 0, 0, 0, 1},
	{0, 1, 1, 1, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 1, 1, 1},
	{-1, -1, 0, 1, 1, 1, 1, 1},
	{0, 0, 0, 1, 1, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 1, 1},
}
