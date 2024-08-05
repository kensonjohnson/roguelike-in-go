package archetype

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/setanarut/kamera/v2"
	"github.com/yohamta/donburi"
)

var CameraTag = donburi.NewTag("camera")

func CreateNewCamera(world donburi.World) {
	camera := world.Entry(world.Create(
		CameraTag,
		component.Camera,
	))

	cameraData := &component.CameraData{
		MainCamera: kamera.NewCamera(
			config.ScreenWidth*config.TileWidth/2,
			config.ScreenHeight*config.TileHeight/2,
			config.ScreenWidth*config.TileWidth,
			config.ScreenHeight*config.TileHeight,
		),
		CamSpeed: 5,
		CamScreen: ebiten.NewImage(
			config.ScreenWidth*config.TileWidth,
			config.ScreenHeight*config.TileHeight,
		),
		CamImageOptions: &ebiten.DrawImageOptions{},
	}
	cameraData.MainCamera.Lerp = true
	cameraData.MainCamera.ZoomFactor = 100
	cameraData.MainCamera.ShakeOptions.MaxShakeAngle = 0

	component.Camera.Set(camera, cameraData)
}
