package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/components"
	"github.com/kensonjohnson/roguelike-game-go/scenes"
)

func ProcessRenderables(g *Game, level scenes.Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[position].(*components.Position)
		img := result.Components[renderable].(*components.Renderable).Image

		if level.PlayerVisible.IsVisible(pos.X, pos.Y) {
			index := level.GetIndexFromXY(pos.X, pos.Y)
			tile := level.Tiles[index]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(img, options)
		}
	}
}
