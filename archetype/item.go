package archetype

import (
	"errors"

	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

func CreateNewItem(world donburi.World, itemData *items.ItemData) *donburi.Entry {
	entry := world.Entry(world.Create(
		tags.ItemTag,
		component.Name,
		component.Sprite,
		component.Description,
	))

	name := component.NameData{
		Value: itemData.Name,
	}
	component.Name.SetValue(entry, name)

	sprite := component.SpriteData{
		Image: itemData.Sprite,
	}
	component.Sprite.SetValue(entry, sprite)

	description := component.DescriptionData{
		Value: itemData.Description,
	}
	component.Description.SetValue(entry, description)

	return entry
}

func isItem(entry *donburi.Entry) bool {
	return entry.HasComponent(tags.ItemTag)
}

func PlaceItemInWorld(entry *donburi.Entry, x, y int, discoverable bool) error {
	if !isItem(entry) {
		return errors.New("entry is not an Item Entity")
	}

	entry.AddComponent(tags.PickupTag)

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
	entry.RemoveComponent(tags.PickupTag)
	entry.RemoveComponent(component.Position)
	if entry.HasComponent(component.Discoverable) {
		entry.RemoveComponent(component.Discoverable)
	}
}
