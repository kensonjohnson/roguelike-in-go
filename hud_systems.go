package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/components"
	"github.com/kensonjohnson/roguelike-game-go/config"
)

func ProcessHUD(g *Game, screen *ebiten.Image) {

	uiX := (config.Config.ScreenWidth * config.Config.TileWidth) / 2
	uiY := (config.Config.ScreenHeight - config.Config.UIHeight) * config.Config.TileHeight
	var fontX = uiX + FONT_SIZE
	var fontY = uiY + 24

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(uiX), float64(uiY))
	screen.DrawImage(assets.UIPanel, options)

	for _, p := range g.World.Query(g.WorldTags["players"]) {
		h := p.Components[health].(*components.Health)
		healthText := fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth)
		text.Draw(screen, healthText, assets.HUDFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		ac := p.Components[armor].(*components.Armor)
		acText := fmt.Sprintf("Armor Class: %d", ac.ArmorClass)
		text.Draw(screen, acText, assets.HUDFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		defText := fmt.Sprintf("Defense: %d", ac.Defense)
		text.Draw(screen, defText, assets.HUDFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		wpn := p.Components[meleeWeapon].(*components.MeleeWeapon)
		dmg := fmt.Sprintf("Damage: %d - %d", wpn.MinimumDamage, wpn.MaximumDamage)
		text.Draw(screen, dmg, assets.HUDFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		bonus := fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus)
		text.Draw(screen, bonus, assets.HUDFont, createTextDrawOptions(fontX, fontY, color.White))
	}
}

func createTextDrawOptions(x, y int, color color.Color) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	return options
}
