package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/colors"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine/shapes"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type ui struct {
	query        donburi.Query
	lastMessages []string
	healthBox    *ebiten.Image
	coinBox      *ebiten.Image
	messageBox   *ebiten.Image
	posX, posY   int
}

var defaultMessages = []string{
	"Welcome to the game!",
	"Use the arrow keys to move.",
	"Press 'Q' to pass your turn.",
	"Good luck!",
}

var UI = &ui{
	query: *donburi.NewQuery(filter.Contains(
		component.UserMessage,
	)),
	lastMessages: defaultMessages,
	healthBox:    createHealthBox(),
	coinBox:      createCoinBox(),
	messageBox: shapes.MakeBox(
		50*config.TileWidth, config.UIHeight*config.TileHeight, 4,
		colors.Peru, color.Black,
		shapes.SimpleCorner,
	),
	posX: 15 * config.TileWidth,
	posY: (config.ScreenHeight - config.UIHeight) * config.TileHeight,
}

func (u *ui) Update(ecs *ecs.ECS) {
	// Get attack messages first
	for entry := range u.query.Iter(ecs.World) {
		messages := component.UserMessage.Get(entry)
		if messages.AttackMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.AttackMessage)
			messages.AttackMessage = ""
		}
		if messages.WorldInteractionMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.WorldInteractionMessage)
			messages.WorldInteractionMessage = ""
		}
	}

	// Then process any deaths, including the player's
	for entry := range u.query.Iter(ecs.World) {
		messages := component.UserMessage.Get(entry)
		if messages.DeadMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.DeadMessage)
			messages.DeadMessage = ""
		}
		if messages.GameStateMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.GameStateMessage)
			messages.GameStateMessage = ""
		}
	}

	if len(u.lastMessages) > 6 {
		// Save just the last 6 messages
		u.lastMessages = u.lastMessages[len(u.lastMessages)-6:]
	}

}

func (u *ui) Draw(ecs *ecs.ECS, screen *ebiten.Image) {

	// Draw the player HUD
	u.drawPlayerHud(screen, ecs.World)

	// Draw the user message box
	u.drawUserMessages(screen, u.lastMessages)

}

func createTextDrawOptions(x, y int, color color.Color) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	return options
}

func (u *ui) drawPlayerHud(screen *ebiten.Image, world donburi.World) {

	// spacing := config.TileWidth

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(u.posX), float64(u.posY-28),
	)
	screen.DrawImage(u.healthBox, options)

	options.GeoM.Translate(float64(u.healthBox.Bounds().Dx()-4), 0)
	screen.DrawImage(u.coinBox, options)

	playerEntry := tags.PlayerTag.MustFirst(world)
	health := component.Health.Get(playerEntry)
	wallet := component.Wallet.Get(playerEntry)

	// Draw the player's info
	fontX := u.posX + 36
	fontY := u.posY - 26

	textOptions := createTextDrawOptions(fontX, fontY, color.White)

	// Health
	// TODO: Color text based on current health
	message := fmt.Sprintf(
		"%d / %d",
		health.CurrentHealth,
		health.MaxHealth,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyMiniSquaredFont,
		textOptions,
	)

	// Move cursor to next box
	textOptions.GeoM.Translate(float64(u.healthBox.Bounds().Dx()), 0)
	message = fmt.Sprintf("%d", wallet.Amount)
	text.Draw(
		screen,
		message,
		assets.KenneyMiniSquaredFont,
		textOptions,
	)
}

func (u *ui) drawUserMessages(screen *ebiten.Image, lastMessages []string) {

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(u.posX), float64(u.posY),
	)
	screen.DrawImage(u.messageBox, options)

	// Draw the user messages
	fontX := u.posX + 28
	fontY := u.posY + 10
	for _, message := range lastMessages {
		if message != "" {
			textOptions := &text.DrawOptions{}
			textOptions.GeoM.Translate(
				float64(fontX),
				float64(fontY),
			)
			textOptions.ColorScale.ScaleWithColor(color.White)
			text.Draw(screen, message, assets.KenneyPixelFont, textOptions)
			fontY += config.FontSize + 2
		}
	}
}

func createHealthBox() *ebiten.Image {
	image := shapes.MakeBox(
		90+assets.Heart.Bounds().Dx()*2,
		assets.Heart.Bounds().Dx()*2,
		4,
		colors.Peru, color.Black,
		shapes.BasicCorner,
	)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(2, 2)
	options.GeoM.Translate(2, 0)
	image.DrawImage(assets.Heart, options)

	return image
}

func createCoinBox() *ebiten.Image {
	image := shapes.MakeBox(
		90+assets.WorldSmallCoin.Bounds().Dx()*2,
		assets.WorldSmallCoin.Bounds().Dx()*2,
		4,
		colors.Peru, color.Black,
		shapes.BasicCorner,
	)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(2, 2)
	options.GeoM.Translate(2, 0)
	image.DrawImage(assets.WorldSmallCoin, options)

	return image
}
