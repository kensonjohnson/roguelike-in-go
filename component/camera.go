package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/kamera/v2"
	"github.com/yohamta/donburi"
)

type CameraData struct {
	MainCamera      *kamera.Camera
	CamSpeed        float64
	CamScreen       *ebiten.Image
	CamImageOptions *ebiten.DrawImageOptions
}

var Camera = donburi.NewComponentType[CameraData]()
