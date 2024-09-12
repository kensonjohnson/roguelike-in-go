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
	hudBox       *ebiten.Image
	messageBox   *ebiten.Image
	divider      *ebiten.Image
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
	hudBox: shapes.MakeBox(
		200, config.UIHeight*config.TileHeight, 4,
		colors.Peru, color.Black,
	),
	messageBox: shapes.MakeBox(
		50*config.TileWidth, config.UIHeight*config.TileHeight, 4,
		colors.Peru, color.Black,
	),
	divider: shapes.MakeDivider(
		config.UIHeight*config.TileHeight-8, 3,
		colors.Peru, color.Transparent,
		false,
	),
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

	spacing := config.TileWidth

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(spacing), float64(spacing),
	)
	screen.DrawImage(u.hudBox, options)

	playerEntry := tags.PlayerTag.MustFirst(world)
	health := component.Health.Get(playerEntry)
	attack := component.Attack.Get(playerEntry)
	defense := component.Defense.Get(playerEntry)

	// Draw the player's info
	fontX := spacing + 28
	fontY := spacing + config.FontSize

	// Health
	message := fmt.Sprintf(
		"Health: %d / %d",
		health.CurrentHealth,
		health.MaxHealth,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	// Armor
	message = fmt.Sprintf(
		"Armor Class: %d",
		defense.ArmorClass,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	message = fmt.Sprintf(
		"Defense: %d",
		defense.Defense,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	// Weapon
	message = fmt.Sprintf(
		"Damage: %d - %d",
		attack.MinimumDamage,
		attack.MaximumDamage,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	message = fmt.Sprintf(
		"To Hit Bonus: %d",
		attack.ToHitBonus,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
}

func (u *ui) drawUserMessages(screen *ebiten.Image, lastMessages []string) {
	spacing := 15 * config.TileWidth
	top := (config.ScreenHeight - config.UIHeight) * config.TileHeight
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(spacing), float64(top),
	)
	screen.DrawImage(u.messageBox, options)

	// Draw the user messages
	fontX := spacing + 28
	fontY := top + 10
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
