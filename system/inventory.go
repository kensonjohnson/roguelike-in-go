package system

import (
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine/shapes"
	"github.com/yohamta/donburi/ecs"
)

// How long until a pressed key registers a new event
const keyDelay = 10

// UI setting
const inset = 20
const boxSize = 48 + 18 // (item sprite * scale) + (border size * 2)
const spacing = 10
const totalBoxSpace = boxSize + spacing
const rows = 4
const columns = 6
const contextWindowWidth = 140
const contextWindowHeight = 150

type inventoryUi struct {
	open                   bool
	background             *ebiten.Image
	posX, posY             int
	selector               *ebiten.Image
	selectorX, selectorY   int
	keyDelayCount          int
	inContextMenu          bool
	contextWindow          *ebiten.Image
	contextFont            *text.GoTextFace
	contextWindowSelection contextSelection
}

var InventoryUI = inventoryUi{
	open:          false,
	background:    buildInventorySprite(),
	posX:          15 * config.TileWidth,
	posY:          (((config.ScreenHeight - config.UIHeight - 2) * config.TileHeight) - (inset + (totalBoxSpace * rows) - spacing + inset)),
	selector:      makeItemBox(color.White),
	selectorX:     0,
	selectorY:     0,
	keyDelayCount: 0,
	inContextMenu: false,
	contextWindow: shapes.MakeBox(
		contextWindowWidth, contextWindowHeight, 4,
		colors.Peru, colors.LightGray,
		shapes.BasicCorner,
	),
	contextFont: &text.GoTextFace{
		Source: assets.KenneyMiniSquaredFont.Source,
		Size:   assets.KenneyMiniSquaredFont.Size * 1.5,
	},
	contextWindowSelection: back,
}

type contextSelection int

const (
	discard contextSelection = iota
	info
	use
	back
)

func (i *inventoryUi) Update(ecs *ecs.ECS) {
	if !i.open {
		return
	}
	if i.keyDelayCount > 0 {
		i.keyDelayCount--
		return
	}

	if i.inContextMenu {
		i.handleContextWindow()
	} else {
		i.handleSelectionWindow()
	}

}

func (i *inventoryUi) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if !i.open {
		return
	}

	options := &ebiten.DrawImageOptions{}

	// Draw the box
	options.GeoM.Translate(
		float64(i.posX),
		float64(i.posY),
	)
	screen.DrawImage(i.background, options)

	// Draw each item
	playerEntry := tags.PlayerTag.MustFirst(ecs.World)
	playerInventory := component.Inventory.Get(playerEntry)
	for index, entry := range playerInventory.Items {
		if entry == nil {
			continue
		}
		sprite := component.Sprite.Get(entry).Image
		column := index % columns
		row := index / columns
		options.GeoM.Reset()
		options.GeoM.Scale(3, 3)
		options.GeoM.Translate(
			float64(i.posX+(column*totalBoxSpace)+inset+9),
			float64(i.posY+(row*totalBoxSpace)+inset+9),
		)
		screen.DrawImage(sprite, options)
	}

	// Draw selector
	selectorPosX := i.posX + (i.selectorX * totalBoxSpace) + inset
	selectorPosY := i.posY + (i.selectorX * totalBoxSpace) + inset
	options.GeoM.Reset()
	options.GeoM.Translate(
		float64(selectorPosX),
		float64(selectorPosY),
	)
	screen.DrawImage(i.selector, options)

	// Draw context window
	if !i.inContextMenu {
		return
	}

	// The position is already on the top left corner of the selection box;
	// We just need to move up by the height of the context window.
	options.GeoM.Translate(0, -float64(i.contextWindow.Bounds().Dy()))
	screen.DrawImage(i.contextWindow, options)

	i.drawContextWindowOptions(screen)
}

func (i *inventoryUi) handleSelectionWindow() {

	if inpututil.IsKeyJustPressed(ebiten.KeyI) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		slog.Debug("Close Inventory")
		Turn.TurnState = PlayerTurn
		i.open = false
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		i.contextWindowSelection = back
		i.inContextMenu = true
		slog.Debug("Open context")
		return
	}

	moveX := 0
	moveY := 0

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		slog.Debug("Pressing Up in Inventory!")
		moveY = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		slog.Debug("Pressing Down in Inventory!")
		moveY = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		slog.Debug("Pressing Left in Inventory!")
		moveX = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		slog.Debug("Pressing Right in Inventory!")
		moveX = 1
	}

	if moveX != 0 || moveY != 0 {
		i.keyDelayCount = keyDelay
	}

	i.selectorX = (i.selectorX + moveX + columns) % columns
	i.selectorY = (i.selectorY + moveY + rows) % rows
}

func (i *inventoryUi) handleContextWindow() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		slog.Debug("Selection made!", "Selection: ", i.contextWindowSelection)
		// Do some work on selection
		if i.contextWindowSelection == back {
			i.inContextMenu = false
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		slog.Debug("Context window closed")
		i.contextWindowSelection = back
		i.inContextMenu = false
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		i.contextWindowSelection--
		if i.contextWindowSelection < discard {
			i.contextWindowSelection = discard
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		i.contextWindowSelection++
		if i.contextWindowSelection > back {
			i.contextWindowSelection = back
		}
	}
}

func (i *inventoryUi) Open() {
	i.open = true
}
func (i *inventoryUi) Close() {
	i.open = false
}

func buildInventorySprite() *ebiten.Image {

	image := shapes.MakeBox(
		inset+(totalBoxSpace*columns)-spacing+inset,
		inset+(totalBoxSpace*rows)-spacing+inset,
		4,
		colors.SlateGray,
		color.Black,
		shapes.SmallPointedCorner,
	)

	itemBox := makeItemBox(colors.CornflowerBlue)
	options := &ebiten.DrawImageOptions{}
	for y := 0; y < rows; y++ {
		options.GeoM.Translate(float64(inset), float64(inset+(y*totalBoxSpace)))
		for x := 0; x < columns; x++ {
			image.DrawImage(itemBox, options)
			options.GeoM.Translate(totalBoxSpace, 0)
		}
		options.GeoM.Reset()
	}

	return image
}

func makeItemBox(border color.Color) *ebiten.Image {
	return shapes.MakeBox(
		boxSize, boxSize, 3,
		border, color.Black,
		shapes.SimpleCorner,
	)
}

func (i *inventoryUi) drawContextWindowOptions(screen *ebiten.Image) {
	x := i.posX + (i.selectorX * totalBoxSpace) + inset
	y := i.posY + (i.selectorY * totalBoxSpace) + inset
	lineHeight := i.contextWindow.Bounds().Dy() / 4
	const inset = 10
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x+inset), float64(y-i.contextWindow.Bounds().Dy()))

	if i.contextWindowSelection == discard {
		options.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		options.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, "Discard", i.contextFont, options)

	options.GeoM.Translate(0, float64(lineHeight))
	options.ColorScale.Reset()
	if i.contextWindowSelection == info {
		options.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		options.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, "Info", i.contextFont, options)

	options.GeoM.Translate(0, float64(lineHeight))
	options.ColorScale.Reset()
	if i.contextWindowSelection == use {
		options.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		options.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, "Equip", i.contextFont, options)

	options.GeoM.Translate(0, float64(lineHeight))
	options.ColorScale.Reset()
	if i.contextWindowSelection == back {
		options.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		options.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, "Back", i.contextFont, options)
}
