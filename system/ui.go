package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type ui struct {
	query        donburi.Query
	lastMessages []string
}

var UI = &ui{
	query: *donburi.NewQuery(filter.Contains(
		component.UserMessage,
	)),
	lastMessages: make([]string, 0, 5),
}

func (u *ui) Update(ecs *ecs.ECS) {
	// Get attack messages first
	u.query.Each(ecs.World, func(entry *donburi.Entry) {
		messages := component.UserMessage.Get(entry)
		if messages.AttackMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.AttackMessage)
			messages.AttackMessage = ""
		}
	})
	// Then process any deaths, including the player's
	u.query.Each(ecs.World, func(entry *donburi.Entry) {
		messages := component.UserMessage.Get(entry)
		if messages.DeadMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.DeadMessage)
			messages.DeadMessage = ""
			ecs.World.Remove(entry.Entity())
		}
		if messages.GameStateMessage != "" {
			u.lastMessages = append(u.lastMessages, messages.GameStateMessage)
			messages.GameStateMessage = ""
		}
	})

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
	screen.DrawImage(assets.UIPanel, options)

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
			text.Draw(screen, message, assets.HUDFont, textOptions)
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
	screen.DrawImage(assets.UIPanel, options)

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
		assets.HUDFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	// Armor
	message = fmt.Sprintf(
		"Armor Class: %d",
		playerHUD.Armor.ArmorClass,
	)
	text.Draw(
		screen,
		message,
		assets.HUDFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	message = fmt.Sprintf(
		"Defense: %d",
		playerHUD.Armor.Defense,
	)
	text.Draw(
		screen,
		message,
		assets.HUDFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	// Weapon
	message = fmt.Sprintf(
		"Damage: %d - %d",
		playerHUD.Weapon.MinimumDamage,
		playerHUD.Weapon.MaximumDamage,
	)
	text.Draw(
		screen,
		message,
		assets.HUDFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
	fontY += config.FontSize + 4

	message = fmt.Sprintf(
		"To Hit Bonus: %d",
		playerHUD.Weapon.ToHitBonus,
	)
	text.Draw(
		screen,
		message,
		assets.HUDFont,
		createTextDrawOptions(fontX, fontY, color.White),
	)
}
