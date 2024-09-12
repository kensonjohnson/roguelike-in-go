package shapes

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type cornerVariant int
type cornerShape [][]int8

const (
	PointedCorner cornerVariant = iota
	PointedCornerTransparent
	SimpleCorner
	SimpleCornerTransparent
)

func MakeBox(w, h, scale int, border, fill color.Color, variant ...cornerVariant) *ebiten.Image {

	image := ebiten.NewImage(w, h)
	corner := makeCornerImage(scale, border, fill, variant...)
	size := corner.Bounds().Size()
	options := &ebiten.DrawImageOptions{}

	// NW
	image.DrawImage(corner, options)

	// NE
	options.GeoM.Rotate(degreesToRadians(90))
	options.GeoM.Translate(float64(w), 0)
	image.DrawImage(corner, options)

	// SE
	options.GeoM.Reset()
	options.GeoM.Rotate(degreesToRadians(180))
	options.GeoM.Translate(float64(w), float64(h))
	image.DrawImage(corner, options)

	// SW
	options.GeoM.Reset()
	options.GeoM.Rotate(degreesToRadians(270))
	options.GeoM.Translate(0, float64(h))
	image.DrawImage(corner, options)

	// Draw top and bottom lines, plus fill
	line := ebiten.NewImage(w-(size.X*2), 1)
	line.Fill(border)
	for i := 0; i < size.Y; i++ {
		if i == scale {
			line.Fill(fill)
		}
		options.GeoM.Reset()
		options.GeoM.Translate(float64(size.X), float64(i))
		image.DrawImage(line, options)
		options.GeoM.Translate(0, float64(h-(i*2)-1))
		image.DrawImage(line, options)
	}

	// Draw vertical lines and fill
	line = ebiten.NewImage(w, h-(size.Y*2))
	line.Fill(fill)
	ends := ebiten.NewImage(scale, h-(size.Y*2))
	ends.Fill(border)
	options.GeoM.Reset()
	options.GeoM.Translate(float64(w-scale), 0)
	line.DrawImage(ends, options)
	options.GeoM.Reset()
	line.DrawImage(ends, options)
	options.GeoM.Translate(0, float64(size.Y))
	image.DrawImage(line, options)

	return image
}

func makeCornerImage(scale int, border, fill color.Color, variant ...cornerVariant) *ebiten.Image {
	var shape cornerShape
	if len(variant) <= 0 {
		shape = pointedCorner
	} else {
		switch variant[0] {
		case PointedCorner:
			shape = pointedCorner
		case PointedCornerTransparent:
			shape = pointedCornerTransparent
		case SimpleCorner:
			shape = simpleCorner
		case SimpleCornerTransparent:
			shape = simpleCornerTransparent
		default:
			shape = pointedCorner
		}
	}

	width := len(shape[0]) * scale
	height := len(shape) * scale

	image := ebiten.NewImage(width, height)
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

	return image
}

func degreesToRadians(degrees int) float64 {
	return float64(degrees) * math.Pi / 180
}

var pointedCorner cornerShape = [][]int8{
	{0, 0, 0, 0, 0, -1, 0, 0},
	{0, 1, 1, 1, 0, -1, 0, 1},
	{0, 1, 1, 1, 0, 0, 0, 1},
	{0, 1, 1, 1, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 1, 1, 1},
	{-1, -1, 0, 1, 1, 1, 1, 1},
	{0, 0, 0, 1, 1, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 1, 1},
}

var pointedCornerTransparent cornerShape = [][]int8{
	{0, 0, 0, 0, 0, -1, 0, 0},
	{0, -1, -1, -1, 0, -1, 0, 1},
	{0, -1, -1, -1, 0, 0, 0, 1},
	{0, -1, -1, -1, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 1, 1, 1},
	{-1, -1, 0, 1, 1, 1, 1, 1},
	{0, 0, 0, 1, 1, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 1, 1},
}

var simpleCorner cornerShape = [][]int8{
	{0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 1, 0, 1, 1, 1},
	{0, 1, 0, 1, 1, 1},
	{0, 0, 0, 1, 1, 1},
}

var simpleCornerTransparent cornerShape = [][]int8{
	{0, 0, 0, 0, 0, 0},
	{0, -1, -1, -1, -1, 0},
	{0, -1, 0, 0, 0, 0},
	{0, -1, 0, 1, 1, 1},
	{0, -1, 0, 1, 1, 1},
	{0, 0, 0, 1, 1, 1},
}
