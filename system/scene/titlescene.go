package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/config"
)

type TitleScene struct {
	count           int
	ImageBackground *ebiten.Image
	PixelWidth      int
	PixelHeight     int
}

func (s *TitleScene) Update() {
	s.count++
	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	// 	// state.SceneManager.GoTo(NewGameScene())
	// 	return nil
	// }

}

const scale = 4

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.drawTitleBackground(screen, s.count)
	x := float64(s.PixelWidth / 2)
	y := 32.0
	drawLogo(screen, "ROGUELIKE DEMO", x, y)

	message := "PRESS SPACE TO START"
	y = float64(s.PixelHeight - 200)
	drawTextWithShadow(
		screen,
		message,
		x, y,
		color.RGBA{R: 178, G: 182, B: 194, A: 255},
		text.AlignCenter, text.AlignStart,
	)

	x = float64((s.PixelWidth / 2) - 200)
	y = float64(s.PixelHeight / 2)
	drawCharacter(screen, x, y)

	x = float64((s.PixelWidth / 2) + 120)
	y = float64((s.PixelHeight / 2) - 150)
	drawSkelly(screen, x, y)

	x = float64((s.PixelWidth / 2) + 200)
	y = float64((s.PixelHeight / 2) + 130)
	drawOrc(screen, x, y)
}

func (s *TitleScene) drawTitleBackground(screen *ebiten.Image, count int) {

	options := &ebiten.DrawImageOptions{}
	offset := float64(count % (config.TileWidth * 4))
	for i := 0; i < (config.ScreenWidth)*(config.ScreenHeight); i++ {
		options.GeoM.Reset()
		x := float64(i%(config.ScreenWidth)) * config.TileWidth * scale
		y := float64(i/(config.ScreenWidth)-1) * config.TileHeight * scale
		dstX := x - offset
		dstY := y + offset
		options.GeoM.Scale(scale, scale)
		options.GeoM.Translate(dstX, dstY)
		screen.DrawImage(s.ImageBackground, options)
	}
}

func drawLogo(screen *ebiten.Image, str string, x, y float64) {
	drawTextWithShadow(screen, str, x, y, color.RGBA{R: 202, G: 146, B: 74, A: 255}, text.AlignCenter, text.AlignStart)
}

func drawCharacter(screen *ebiten.Image, x, y float64) {
	tileOffset := float64(config.TileWidth * scale * scale / 2)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scale*scale, scale*scale)
	options.GeoM.Translate(x-tileOffset, y-tileOffset)
	screen.DrawImage(assets.Player, options)
}

func drawSkelly(screen *ebiten.Image, x, y float64) {
	tileOffset := float64(config.TileWidth * scale * scale / 2)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scale*scale, scale*scale)
	options.GeoM.Translate(x-tileOffset, y-tileOffset)
	screen.DrawImage(assets.Skelly, options)
}

func drawOrc(screen *ebiten.Image, x, y float64) {
	tileOffset := float64(config.TileWidth * scale * scale / 2)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scale*scale, scale*scale)
	options.GeoM.Translate(x-tileOffset, y-tileOffset)
	screen.DrawImage(assets.Orc, options)
}

func drawTextWithShadow(screen *ebiten.Image, message string, x, y float64, clr color.Color, primaryAlign, secondaryAlign text.Align) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x)+1, float64(y)+1)
	options.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 0x80})
	options.LineSpacing = config.FontSize * float64(scale)
	options.PrimaryAlign = primaryAlign
	options.SecondaryAlign = secondaryAlign
	text.Draw(screen, message, &text.GoTextFace{
		Source: assets.KenneyMiniSquaredFont.Source,
		Size:   config.FontSize * float64(scale),
	}, options)

	options.GeoM.Reset()
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.Reset()
	options.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, message, &text.GoTextFace{
		Source: assets.KenneyMiniSquaredFont.Source,
		Size:   config.FontSize * float64(scale),
	}, options)
}
