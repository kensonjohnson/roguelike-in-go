package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/fonts"
)

var hudFont *text.GoTextFace = nil

func ProcessHUD(g *Game, screen *ebiten.Image) {

	if hudFont == nil {
		source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
		if err != nil {
			log.Fatal(err)
		}
		hudFont = &text.GoTextFace{
			Source: source,
			Size:   FONT_SIZE,
		}
	}

	gd := NewGameData()

	uiX := (gd.ScreenWidth * gd.TileWidth) / 2
	uiY := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	var fontX = uiX + FONT_SIZE
	var fontY = uiY + 24

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(uiX), float64(uiY))
	screen.DrawImage(assets.UIPanel, options)

	for _, p := range g.World.Query(g.WorldTags["players"]) {
		h := p.Components[health].(*Health)
		healthText := fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth)
		text.Draw(screen, healthText, hudFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		ac := p.Components[armor].(*Armor)
		acText := fmt.Sprintf("Armor Class: %d", ac.ArmorClass)
		text.Draw(screen, acText, hudFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		defText := fmt.Sprintf("Defense: %d", ac.Defense)
		text.Draw(screen, defText, hudFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		wpn := p.Components[meleeWeapon].(*MeleeWeapon)
		dmg := fmt.Sprintf("Damage: %d - %d", wpn.MinimumDamage, wpn.MaximumDamage)
		text.Draw(screen, dmg, hudFont, createTextDrawOptions(fontX, fontY, color.White))
		fontY += FONT_SIZE
		bonus := fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus)
		text.Draw(screen, bonus, hudFont, createTextDrawOptions(fontX, fontY, color.White))
	}
}

func createTextDrawOptions(x, y int, color color.Color) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	return options
}
