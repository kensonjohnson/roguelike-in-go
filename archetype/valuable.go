package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

func CreateNewValuable(world donburi.World, valuableData *items.ValuableData) *donburi.Entry {
	entry := CreateNewItem(world, &valuableData.ItemData)

	entry.AddComponent(tags.ValuableTag)

	entry.AddComponent(component.Value)
	value := component.ValueData{
		Amount: valuableData.Value,
	}
	component.Value.SetValue(entry, value)

	return entry
}

func CreateCoins(world donburi.World, valuableData *items.ValuableData) *donburi.Entry {
	entry := CreateNewValuable(world, valuableData)

	entry.AddComponent(tags.CoinTag)

	return entry
}

func IsCoin(entry donburi.Entry) bool {
	return entry.HasComponent(tags.CoinTag)
}
