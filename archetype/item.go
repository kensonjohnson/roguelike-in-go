package archetype

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/yohamta/donburi"
)

func CreateNewItem(world donburi.World, itemId int, itemName string, itemImage *ebiten.Image) *donburi.Entry {
	item := world.Entry(world.Create(
		component.ItemId,
		component.Name,
		component.Sprite,
	))

	id := component.ItemIdData{
		Id: itemId,
	}
	component.ItemId.SetValue(item, id)

	name := component.NameData{
		Value: itemName,
	}
	component.Name.SetValue(item, name)

	sprite := component.SpriteData{
		Image: itemImage,
	}
	component.Sprite.SetValue(item, sprite)

	return item
}

func isItem(entry *donburi.Entry) bool {
	return entry.HasComponent(component.ItemId)
}

func PlaceItemInWorld(entry *donburi.Entry, x, y int, discoverable bool) error {
	if !isItem(entry) {
		return errors.New("entry is not an Item Entity")
	}

	entry.AddComponent(component.Position)
	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(entry, position)

	if discoverable {
		discovery := component.DiscoverableData{}
		component.Discoverable.SetValue(entry, discovery)
	}

	return nil
}

func RemoveItemFromWorld(entry *donburi.Entry) {
	entry.RemoveComponent(component.Position)
}
