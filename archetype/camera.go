package archetype

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/setanarut/kamera/v2"
	"github.com/yohamta/donburi"
)

var CameraTag = donburi.NewTag("camera")

func CreateNewCamera(world donburi.World) {
	var entry *donburi.Entry
	var ok bool
	if entry, ok = PlayerTag.First(world); !ok {
		log.Fatal("CreateNewCamera failed: Player not found")
	}
	playerPosition := component.Position.Get(entry)

	entry = world.Entry(world.Create(
		CameraTag,
		component.Camera,
	))

	cameraData := &component.CameraData{
		MainCamera: kamera.NewCamera(
			float64((playerPosition.X*config.TileWidth)+config.TileWidth/2),
			float64((playerPosition.Y*config.TileHeight)+config.TileHeight/2),
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
