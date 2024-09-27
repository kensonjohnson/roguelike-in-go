package system

import (
	"fmt"
	"image/color"
	"log"
	"log/slog"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
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
const contextWindowHeight = 200

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
	inConfirmAction        bool
	confirmActionWindow    *ebiten.Image
	confirmAction          bool
	inInfoWindow           bool
	infoWindow             *ebiten.Image
	infoWindowText         *infoWindowText
}

var InventoryUI = inventoryUi{
	open:       false,
	background: buildInventorySprite(),
	posX:       15 * config.TileWidth,
	posY: (((config.ScreenHeight - config.UIHeight - 2) * config.TileHeight) -
		(inset + (totalBoxSpace * rows) - spacing + inset)),
	selector:      makeItemBox(color.White, color.Transparent),
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
	inConfirmAction:        false,
	confirmActionWindow: shapes.MakeBox(
		130, 40, 4,
		colors.Peru, colors.LightGray,
		shapes.BasicCorner,
	),
	confirmAction:  false,
	inInfoWindow:   false,
	infoWindow:     nil,
	infoWindowText: &infoWindowText{Name: "Info", Text: "No Description"},
}

type contextSelection int

const (
	discard contextSelection = iota
	info
	use
	back
)

type infoWindowText struct {
	Name string
	Text string
}

func (i *inventoryUi) Update(ecs *ecs.ECS) {
	if !i.open {
		return
	}
	if i.keyDelayCount > 0 {
		i.keyDelayCount--
		return
	}

	if i.inContextMenu {
		i.handleContextWindow(ecs)
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
	for index, entry := range playerInventory.Iter() {
		if entry == nil {
			continue
		}
		sprite := component.Sprite.Get(entry).Image
		options.GeoM.Reset()
		options.GeoM.Scale(3, 3)
		options.GeoM.Translate(
			float64(i.posX+((index%columns)*totalBoxSpace)+inset+9),
			float64(i.posY+((index/columns)*totalBoxSpace)+inset+9),
		)
		screen.DrawImage(sprite, options)
	}

	// Draw selector
	options.GeoM.Reset()
	options.GeoM.Translate(
		float64(i.posX+(i.selectorX*totalBoxSpace)+inset),
		float64(i.posY+(i.selectorY*totalBoxSpace)+inset),
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

	if inpututil.IsKeyJustPressed(ebiten.KeyI) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
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

func (i *inventoryUi) handleContextWindow(ecs *ecs.ECS) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		slog.Debug("Context window closed")
		i.contextWindowSelection = back
		i.inConfirmAction = false
		i.inInfoWindow = false
		i.inContextMenu = false
		i.infoWindow = nil
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		i.inConfirmAction || i.inInfoWindow {
		i.handleSelectionMade(ecs)
		return
	}

	// We use the `back` constant because it is the last in the enum. The magic
	// +2 in the 'down' keypress is because the enum is zero indexed. Instead of
	// writing `back - 1 + 1` and `back + 1 + 1`, we simplify.
	if inpututil.IsKeyJustPressed(ebiten.KeyW) ||
		inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		i.contextWindowSelection = (i.contextWindowSelection + back) % (back + 1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		i.contextWindowSelection = (i.contextWindowSelection + back + 2) % (back + 1)
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

	itemBox := makeItemBox(colors.Gray, colors.Smudgy)
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

func makeItemBox(border, fill color.Color) *ebiten.Image {
	return shapes.MakeBox(
		boxSize, boxSize, 3,
		border, fill,
		shapes.SimpleCorner,
	)
}

func (i *inventoryUi) drawContextWindowOptions(screen *ebiten.Image) {
	x := i.posX + (i.selectorX * totalBoxSpace) + inset
	y := i.posY + (i.selectorY * totalBoxSpace) + inset

	lineHeight := i.contextWindow.Bounds().Dy() / 4

	const inset = 10

	options := &text.DrawOptions{}
	options.GeoM.Translate(
		float64(x+inset),
		float64(y-i.contextWindow.Bounds().Dy()),
	)

	i.drawContextOption(screen, "Discard", discard, options)

	options.GeoM.Translate(0, float64(lineHeight))
	i.drawContextOption(screen, "Info", info, options)

	options.GeoM.Translate(0, float64(lineHeight))
	i.drawContextOption(screen, "Equip", use, options)

	options.GeoM.Translate(0, float64(lineHeight))
	i.drawContextOption(screen, "Back", back, options)

	if i.inConfirmAction {
		i.drawConfirmWindow(
			screen,
			x+i.contextWindow.Bounds().Dx(),
			y-i.contextWindow.Bounds().Dy(),
		)
	}

	if i.inInfoWindow {
		i.drawInfoWindow(
			screen,
			x+i.contextWindow.Bounds().Dx(),
			y-i.contextWindow.Bounds().Dy(),
		)
	}
}

func (i *inventoryUi) drawContextOption(
	screen *ebiten.Image,
	label string,
	selection contextSelection,
	options *text.DrawOptions,
) {
	if i.contextWindowSelection == selection {
		options.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		options.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, label, i.contextFont, options)
	options.ColorScale.Reset()
}

func (i *inventoryUi) handleSelectionMade(ecs *ecs.ECS) {

	switch i.contextWindowSelection {
	case discard:
		if i.inConfirmAction {

			if inpututil.IsKeyJustPressed(ebiten.KeyA) ||
				inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				i.confirmAction = false
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyD) ||
				inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				i.confirmAction = true
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) &&
				i.confirmAction {

				slog.Debug("Discard item")
				playerEntry := tags.PlayerTag.MustFirst(ecs.World)
				playerInventory := component.Inventory.Get(playerEntry)
				index := i.selectorX + (i.selectorY * columns)
				playerInventory.RemoveItem(index)
				// TODO: Send event message to ui
				i.inConfirmAction = false
				i.inContextMenu = false
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) &&
				!i.confirmAction {
				i.inConfirmAction = false
			}
		} else {
			i.inConfirmAction = true
			i.confirmAction = false
			return
		}

	case info:
		if i.inInfoWindow {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				i.inInfoWindow = false
				i.infoWindow = nil
			}
		} else {
			slog.Debug("Item info")
			playerEntry := tags.PlayerTag.MustFirst(ecs.World)
			playerInventory := component.Inventory.Get(playerEntry)
			itemEntry, err := playerInventory.GetItem(i.selectorX + (i.selectorY * columns))
			if err != nil {
				log.Panic(err)
			}
			i.infoWindowText.Name = component.Name.Get(itemEntry).Value
			var description = "No description"
			if archetype.IsConsumable(itemEntry) {
				value := component.Heal.Get(itemEntry).HealAmount
				description = fmt.Sprintf("Heals for %v", value)
			}
			if archetype.IsValuable(itemEntry) {
				itemDescription := component.Description.Get(itemEntry)
				value := component.Value.Get(itemEntry).Amount
				description = fmt.Sprintf("%v Worth %v gold", itemDescription.Value, value)
			}
			i.infoWindowText.Text = description
			i.inInfoWindow = true
		}

	case use:
		slog.Debug("Item action")

	case back:
		slog.Debug("Close context window")
		i.inContextMenu = false
	}
}

func (i *inventoryUi) drawConfirmWindow(screen *ebiten.Image, x, y int) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(i.confirmActionWindow, options)

	const inset = 10

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(float64(x+inset), float64(y))
	if !i.confirmAction {
		textOptions.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		textOptions.ColorScale.ScaleWithColor(color.Black)
	}
	text.Draw(screen, "No", i.contextFont, textOptions)
	textOptions.ColorScale.Reset()

	textOptions.GeoM.Translate(50.0, 0)
	if i.confirmAction {
		textOptions.ColorScale.ScaleWithColor(colors.DarkGray)
	} else {
		textOptions.ColorScale.ScaleWithColor(color.Black)
	}

	text.Draw(screen, "Yes", i.contextFont, textOptions)

}

func (i *inventoryUi) drawInfoWindow(screen *ebiten.Image, x, y int) {
	if i.infoWindow == nil {
		// find minimum width needed for title
		titleWidth := text.Advance(i.infoWindowText.Name, i.contextFont)
		// find proper width for item description text
		squaredCharCount := int(math.Sqrt(float64(len(i.infoWindowText.Text))))
		desiredLineWidth := text.Advance(strings.Repeat("m", squaredCharCount*2), assets.KenneyMiniSquaredFont)
		// set width according to which length is longer
		width := math.Max(titleWidth, desiredLineWidth)
		// find proper height based on found width
		fields := strings.Fields(i.infoWindowText.Text)
		lines := make([]string, 0)
		var currentLine string

		for index, string := range fields {
			if index == 0 {
				currentLine = string
				continue
			}
			if text.Advance(currentLine+" "+string, assets.KenneyMiniSquaredFont) > float64(width) {
				lines = append(lines, currentLine)
				currentLine = string
			} else {
				currentLine += " " + string
			}
		}
		lines = append(lines, currentLine)
		formattedDescription := strings.Join(lines, "\n")
		_, descriptionHeight := text.Measure(formattedDescription, assets.KenneyMiniSquaredFont, 25.0)
		// create text box, with insets for title and back button
		const inset = 10
		window := shapes.MakeBox(
			inset+int(width)+inset,
			inset+40+int(descriptionHeight)+20+inset,
			4,
			colors.Peru, colors.LightGray,
			shapes.BasicCorner,
		)
		// draw title
		textOptions := &text.DrawOptions{}
		textOptions.GeoM.Translate(float64(+inset), 0)
		textOptions.ColorScale.ScaleWithColor(color.Black)
		text.Draw(window, i.infoWindowText.Name, i.contextFont, textOptions)
		// draw description
		textOptions.GeoM.Translate(0, 40.0)
		textOptions.LineSpacing = 25
		text.Draw(window, formattedDescription, assets.KenneyMiniSquaredFont, textOptions)
		// draw back button
		textOptions.GeoM.Translate(0, descriptionHeight)
		textOptions.ColorScale.Reset()
		textOptions.ColorScale.ScaleWithColor(colors.DarkGray)
		text.Draw(window, "Back", i.contextFont, textOptions)

		i.infoWindow = window
	}

	// draw box
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(i.infoWindow, options)
}
