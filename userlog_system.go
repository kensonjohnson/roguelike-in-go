package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/fonts"
)

var (
	mplusNormalFont *text.GoTextFace = nil
	lastText        []string         = make([]string, 0, 5)
)

const FONT_SIZE = 16

func ProcessUserLog(g *Game, screen *ebiten.Image) {

	if mplusNormalFont == nil {
		source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
		if err != nil {
			log.Fatal(err)
		}
		mplusNormalFont = &text.GoTextFace{
			Source: source,
			Size:   FONT_SIZE,
		}
	}
	gd := NewGameData()

	uiLocation := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	var fontX = FONT_SIZE
	var fontY = uiLocation + 24
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0.), float64(uiLocation))
	screen.DrawImage(assets.UIPanel, op)
	tmpMessages := make([]string, 0, 5)
	anyMessages := false

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.AttackMessage != "" {
			tmpMessages = append(tmpMessages, messages.AttackMessage)
			anyMessages = true
			messages.AttackMessage = ""
		}
	}
	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.DeadMessage != "" {
			tmpMessages = append(tmpMessages, messages.DeadMessage)
			anyMessages = true
			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}
		if messages.GameStateMessage != "" {
			tmpMessages = append(tmpMessages, messages.GameStateMessage)
			anyMessages = true
			//No need to clear, it's all over
		}

	}
	if anyMessages {
		lastText = tmpMessages
	}
	for _, msg := range lastText {
		if msg != "" {
			options := &text.DrawOptions{}
			options.GeoM.Translate(float64(fontX), float64(fontY))
			options.ColorScale.ScaleWithColor(color.White)
			text.Draw(screen, msg, mplusNormalFont, options)
			fontY += FONT_SIZE
		}
	}

}
