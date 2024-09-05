package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type ui struct {
	query        donburi.Query
	lastMessages []string
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
	entry := archetype.UITag.MustFirst(ecs.World)
	ui := component.UI.Get(entry)

	// Draw the user message box
	drawUserMessages(screen, &ui.MessageBox, u.lastMessages)

	// Draw the player HUD
	drawPlayerHud(screen, &ui.PlayerHUD)

}

func createTextDrawOptions(x, y int, color color.Color) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	return options
}

func drawUserMessages(screen *ebiten.Image, messageBox *component.UserMessageBoxData, lastMessages []string) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(messageBox.Position.X),
		float64(messageBox.Position.Y),
	)
	screen.DrawImage(messageBox.Sprite, options)

	// Draw the user messages
	fontX := messageBox.FontX
	fontY := messageBox.FontY
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

func drawPlayerHud(screen *ebiten.Image, playerHUD *component.PlayerHUDData) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(playerHUD.Position.X),
		float64(playerHUD.Position.Y),
	)
	screen.DrawImage(playerHUD.Sprite, options)

	// Draw the player's info
	fontX := playerHUD.FontX
	fontY := playerHUD.FontY

	// Health
	message := fmt.Sprintf(
		"Health: %d / %d",
		playerHUD.Health.CurrentHealth,
		playerHUD.Health.MaxHealth,
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
		playerHUD.Defense.ArmorClass,
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
		playerHUD.Defense.Defense,
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
		playerHUD.Attack.MinimumDamage,
		playerHUD.Attack.MaximumDamage,
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
		playerHUD.Attack.ToHitBonus,
	)
	text.Draw(
		screen,
		message,
		assets.KenneyPixelFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
}
