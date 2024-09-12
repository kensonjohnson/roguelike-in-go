package system

import (
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine/shapes"
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
	image := shapes.MakeBox(
		(45)*config.TileWidth,
		(40)*config.TileHeight,
		4,
		colors.SlateGray,
		color.Black,
		// shapes.PointedCornerTransparent,
	)
	// divider := shapes.MakeDivider(
	// 	(35)*config.TileWidth-8,
	// 	4,
	// 	colors.SlateGray,
	// 	color.Transparent,
	// 	true,
	// )

	options := &ebiten.DrawImageOptions{}

	// // Rotate divider and draw it in the box
	// options.GeoM.Translate(-float64(divider.Bounds().Dx()/2), -float64(divider.Bounds().Dy()/2))
	// options.GeoM.Rotate(engine.DegreesToRadians(-90))
	// options.GeoM.Translate(float64(image.Bounds().Dx()/2), float64(image.Bounds().Dy()/3))
	// image.DrawImage(divider, options)

	// Draw the box
	options.GeoM.Reset()
	options.GeoM.Translate(float64((config.ScreenWidth*config.TileWidth-image.Bounds().Dx())/2), float64(5*config.TileHeight))
	screen.DrawImage(image, options)
}

func (i *inventoryUi) Open() {
	i.open = true
}
func (i *inventoryUi) Close() {
	i.open = false
}
