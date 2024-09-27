package shapes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine"
)

type dividerVariant int
type dividerShape [][]int8

const (
	SimpleDivider dividerVariant = iota
)

func MakeDivider(h, scale int, border, fill color.Color, fade bool, variant ...dividerVariant) *ebiten.Image {
	var shape dividerShape
	if len(variant) <= 0 {
		shape = simpleDivider
	} else {
		switch variant[0] {
		case 0:
			shape = simpleDivider
		default:
			shape = simpleDivider
		}
	}

	width := len(shape[0]) * scale

	image := ebiten.NewImage(width, h)

	// Draw the fancy bit at the top
	block := ebiten.NewImage(scale, scale)
	options := &ebiten.DrawImageOptions{}
	for y, row := range shape {
		for x, value := range row {
			options.GeoM.Reset()
			options.GeoM.Translate(float64(x*scale), float64(y*scale))
			switch value {
			case -1:
				continue
			case 0:
				block.Fill(border)
			case 1:
				block.Fill(fill)
			}
			image.DrawImage(block, options)
		}
	}

	// Create the slice that will fill out the rest of the divider
	filler := ebiten.NewImage(width, 1)
	line := ebiten.NewImage(scale, 1)
	for x, value := range shape[len(shape)-1] {
		options.GeoM.Reset()
		options.GeoM.Translate(float64(x*scale), 0)
		switch value {
		case -1:
			continue
		case 0:
			line.Fill(border)
		case 1:
			line.Fill(fill)
		}
		filler.DrawImage(line, options)
	}

	if fade {
		// Fill in the divider, up to the two thirds point
		twoThirds := (h / 3) * 2
		for i := len(shape) * scale; i < twoThirds; i++ {
			options.GeoM.Reset()
			options.GeoM.Translate(0, float64(i))
			image.DrawImage(filler, options)
		}

		// Fill in the divider, fading to full transparency
		for i := twoThirds; i < h; i++ {
			options.GeoM.Reset()
			options.GeoM.Translate(0, float64(i))
			options.ColorScale.ScaleAlpha(engine.Normalize(i%twoThirds, h-twoThirds))
			image.DrawImage(filler, options)
			options.ColorScale.Reset()
		}
	} else {
		// Fill in the divider for remaining height
		for i := len(shape) * scale; i < h; i++ {
			options.GeoM.Reset()
			options.GeoM.Translate(0, float64(i))
			image.DrawImage(filler, options)
		}
	}

	return image
}

/*
	For any given dividerShape, the final row must be the repeating pattern
	that completes the part "under" the design. This last row will be repeated
	in MakeDivider to satisfy the height requirement.
*/

var simpleDivider dividerShape = [][]int8{
	{-1, -1, 0, 1, 0, -1, -1},
	{-1, -1, 0, 1, 0, -1, -1},
	{0, 0, 0, 1, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 1, 0, 0, 0},
	{-1, -1, 0, 1, 0, -1, -1},
	{-1, 0, 0, 1, 0, 0, -1},
	{-1, 0, 1, 1, 1, 0, -1},
	{-1, 0, 1, 1, 1, 0, -1},
	{-1, 0, 0, 1, 0, 0, -1},
	{-1, -1, 0, 1, 0, -1, -1},
}
