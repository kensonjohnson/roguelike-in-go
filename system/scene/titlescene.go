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

	y = float64(s.PixelHeight / 2)
	drawCharacter(screen, x, y)

	message := "PRESS SPACE TO START"
	y = float64(s.PixelHeight - 200)
	drawTextWithShadow(screen, message, x, y, color.RGBA{R: 178, G: 182, B: 194, A: 255}, text.AlignCenter, text.AlignStart)
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image, count int) {

	op := &ebiten.DrawImageOptions{}
	offset := float64(count % (config.TileWidth * 4))
	for i := 0; i < (config.ScreenWidth)*(config.ScreenHeight); i++ {
		op.GeoM.Reset()
		x := float64(i%(config.ScreenWidth)) * config.TileWidth * scale
		y := float64(i/(config.ScreenWidth)-1) * config.TileHeight * scale
		dstX := x - offset
		dstY := y + offset
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(dstX, dstY)
		r.DrawImage(s.ImageBackground, op)
	}
}

func drawLogo(r *ebiten.Image, str string, x, y float64) {
	drawTextWithShadow(r, str, x, y, color.RGBA{R: 202, G: 146, B: 74, A: 255}, text.AlignCenter, text.AlignStart)
}

func drawCharacter(r *ebiten.Image, x, y float64) {
	tileOffset := float64(config.TileWidth * scale * scale / 2)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale*scale, scale*scale)
	op.GeoM.Translate(x-tileOffset, y-tileOffset)
	r.DrawImage(assets.Player, op)
}

func drawTextWithShadow(rt *ebiten.Image, str string, x, y float64, clr color.Color, primaryAlign, secondaryAlign text.Align) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x)+1, float64(y)+1)
	op.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 0x80})
	op.LineSpacing = config.FontSize * float64(scale)
	op.PrimaryAlign = primaryAlign
	op.SecondaryAlign = secondaryAlign
	text.Draw(rt, str, &text.GoTextFace{
		Source: assets.KenneyMiniSquaredFont.Source,
		Size:   config.FontSize * float64(scale),
	}, op)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.Reset()
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(rt, str, &text.GoTextFace{
		Source: assets.KenneyMiniSquaredFont.Source,
		Size:   config.FontSize * float64(scale),
	}, op)
}
