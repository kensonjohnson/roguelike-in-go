package archetype

import (
	"errors"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/consumables"
	"github.com/yohamta/donburi"
)

var ConsumableTag = donburi.NewTag("consumable")

func CreateNewConsumable(world donburi.World, consumablesId consumables.ConsumablesId) {
	consumable := world.Entry(world.Create(
		ConsumableTag,
		component.Name,
		component.Heal,
		component.Sprite,
	))

	consumableData := consumables.Data[consumablesId]

	name := component.NameData{
		Value: consumableData.Name,
	}
	component.Name.SetValue(consumable, name)

	heal := component.HealData{
		HealAmount: consumableData.AmountHeal,
	}
	component.Heal.SetValue(consumable, heal)

	sprite := component.SpriteData{
		Image: consumableData.Sprite,
	}
	component.Sprite.SetValue(consumable, sprite)
}

func IsConsumable(entry *donburi.Entry) bool {
	return entry.HasComponent(ConsumableTag)
}

func PlaceConsumableInWorld(world *donburi.World, entry *donburi.Entry, x, y int) error {
	if !IsConsumable(entry) {
		return errors.New("entry is not an Consumable Entity")
	}

	entry.AddComponent(component.Position)
	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(entry, position)

	return nil
}
