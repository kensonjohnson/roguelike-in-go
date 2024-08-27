package archetype

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/yohamta/donburi"
)

var PickupTag = donburi.NewTag("pickup")

func CreateNewPickup(world donburi.World, level *component.LevelData, name string, x, y int, spriteImage *ebiten.Image) {
	pickup := world.Entry(world.Create(
		PickupTag,
		component.Name,
		component.Position,
		component.Sprite,
	))

	itemName := component.NameData{
		Label: name,
	}
	component.Name.SetValue(pickup, itemName)

	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(pickup, position)

	sprite := component.SpriteData{
		Image: spriteImage,
	}
	component.Sprite.SetValue(pickup, sprite)
}
