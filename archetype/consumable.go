package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/consumables"
	"github.com/yohamta/donburi"
)

var ConsumableTag = donburi.NewTag("consumable")

func CreateNewConsumable(world donburi.World, consumablesId consumables.ConsumablesId) *donburi.Entry {
	consumableData := consumables.Data[consumablesId]

	entry := CreateNewItem(world, int(consumablesId), consumableData.Name, consumableData.Sprite)

	// Mark as a consumable
	entry.AddComponent(ConsumableTag)

	// Add heal data
	entry.AddComponent(component.Heal)
	heal := component.HealData{
		HealAmount: consumableData.AmountHeal,
	}
	component.Heal.SetValue(entry, heal)

	return entry
}

func IsConsumable(entry *donburi.Entry) bool {
	return entry.HasComponent(ConsumableTag)
}
