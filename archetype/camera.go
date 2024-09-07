package archetype

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/setanarut/kamera/v2"
	"github.com/yohamta/donburi"
)

func CreateNewCamera(world donburi.World) {
	entry := world.Entry(world.Create(
		tags.CameraTag,
		component.Camera,
	))

	cameraData := &component.CameraData{
		MainCamera: kamera.NewCamera(
			0, 0,
			config.ScreenWidth*config.TileWidth,
			(config.ScreenHeight-config.UIHeight)*config.TileHeight,
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
	cameraData.MainCamera.ShakeOptions.Decay = 0.5

	component.Camera.Set(entry, cameraData)
}
