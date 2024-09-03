package archetype

import (
	"errors"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

var ItemTag = donburi.NewTag("item")

func CreateNewItem(world donburi.World, itemData *items.ItemData) *donburi.Entry {
	item := world.Entry(world.Create(
		ItemTag,
		component.Name,
		component.Sprite,
	))

	name := component.NameData{
		Value: itemData.Name,
	}
	component.Name.SetValue(item, name)

	sprite := component.SpriteData{
		Image: itemData.Sprite,
	}
	component.Sprite.SetValue(item, sprite)

	return item
}

func isItem(entry *donburi.Entry) bool {
	return entry.HasComponent(ItemTag)
}

var PickupTag = donburi.NewTag("pickup")

func PlaceItemInWorld(entry *donburi.Entry, x, y int, discoverable bool) error {
	if !isItem(entry) {
		return errors.New("entry is not an Item Entity")
	}

	entry.AddComponent(PickupTag)

	entry.AddComponent(component.Position)
	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(entry, position)

	if discoverable {
		entry.AddComponent(component.Discoverable)
		discovery := component.DiscoverableData{}
		component.Discoverable.SetValue(entry, discovery)
	}

	return nil
}

func RemoveItemFromWorld(entry *donburi.Entry) {
	entry.RemoveComponent(PickupTag)
	entry.RemoveComponent(component.Position)
	if entry.HasComponent(component.Discoverable) {
		entry.RemoveComponent(component.Discoverable)
	}
}
