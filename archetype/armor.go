package archetype

import (
	"errors"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/armors"
	"github.com/yohamta/donburi"
)

var ArmorTag = donburi.NewTag("armor")

func CreateNewArmor(world donburi.World, armorId armors.ArmorId) *donburi.Entry {
	armor := world.Entry(world.Create(
		ArmorTag,
		component.ItemId,
		component.Name,
		component.Sprite,
		component.Defense,
	))

	armorData := armors.Data[armorId]

	itemId := component.ItemIdData{
		Id: int(armorId),
	}
	component.ItemId.SetValue(armor, itemId)

	name := component.NameData{
		Value: armorData.Name,
	}
	component.Name.SetValue(armor, name)

	sprite := component.SpriteData{
		Image: armorData.Sprite,
	}
	component.Sprite.SetValue(armor, sprite)

	defense := component.DefenseData{
		Defense:    armorData.Defense,
		ArmorClass: armorData.ArmorClass,
	}
	component.Defense.SetValue(armor, defense)

	return armor
}

func IsArmor(entry *donburi.Entry) bool {
	return entry.HasComponent(ArmorTag)
}

func PlaceArmorInWorld(world *donburi.World, entry *donburi.Entry, x, y int) error {
	if !IsArmor(entry) {
		return errors.New("entry is not an Armor Entity")
	}

	entry.AddComponent(component.Position)
	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(entry, position)

	return nil
}
