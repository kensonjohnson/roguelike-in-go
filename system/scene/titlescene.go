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

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.drawTitleBackground(screen, s.count)
	drawLogo(screen, "ROGUELIKE DEMO", s.PixelWidth)

	message := "PRESS SPACE TO START"
	x := s.PixelWidth / 2
	y := s.PixelHeight - 200
	drawTextWithShadow(screen, message, x, y, 3, color.RGBA{R: 178, G: 182, B: 194, A: 255}, text.AlignCenter, text.AlignStart)
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image, count int) {
	width, height := s.ImageBackground.Bounds().Dx(), s.ImageBackground.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < ((s.PixelWidth)/width+1)*(s.PixelHeight/height+2); i++ {
		op.GeoM.Reset()
		dx := -(count / 4) % width * 4
		dy := (count / 4) % height * 4
		dstX := (i%(s.PixelWidth/width+1))*width*4 + dx
		dstY := (i/(s.PixelWidth/width+1)-1)*height*4 + dy
		op.GeoM.Scale(4, 4)
		op.GeoM.Translate(float64(dstX), float64(dstY))
		r.DrawImage(s.ImageBackground, op)
	}
}

func drawLogo(r *ebiten.Image, str string, pixelWidth int) {
	const scale = 4
	x := pixelWidth / 2
	y := 32
	drawTextWithShadow(r, str, x, y, scale, color.RGBA{R: 202, G: 146, B: 74, A: 255}, text.AlignCenter, text.AlignStart)
}

var (
	shadowColor = color.RGBA{0, 0, 0, 0x80}
)

func drawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, primaryAlign, secondaryAlign text.Align) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x)+1, float64(y)+1)
	op.ColorScale.ScaleWithColor(shadowColor)
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
