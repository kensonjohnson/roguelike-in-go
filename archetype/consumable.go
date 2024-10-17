package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

func CreateNewConsumable(world donburi.World, consumableData items.ConsumableData) *donburi.Entry {

	entry := CreateNewItem(world, &consumableData.ItemData)

	// Mark as a consumable
	entry.AddComponent(tags.ConsumableTag)

	// Add heal data
	entry.AddComponent(component.Heal)
	heal := component.HealData{
		Amount: consumableData.AmountHeal,
	}
	component.Heal.SetValue(entry, heal)

	return entry
}

func IsConsumable(entry *donburi.Entry) bool {
	return entry.HasComponent(tags.ConsumableTag)
}
