package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/armors"
	"github.com/yohamta/donburi"
)

var ArmorTag = donburi.NewTag("armor")

func CreateNewArmor(world donburi.World, armorId armors.ArmorId) *donburi.Entry {
	armorData := armors.Data[armorId]
	entry := CreateNewItem(world, int(armorId), armorData.Name, armorData.Sprite)

	// Mark as an armor
	entry.AddComponent(ArmorTag)

	// Add defense data
	entry.AddComponent(component.Defense)
	defense := component.DefenseData{
		Defense:    armorData.Defense,
		ArmorClass: armorData.ArmorClass,
	}
	component.Defense.SetValue(entry, defense)

	return entry
}

func IsArmor(entry *donburi.Entry) bool {
	return entry.HasComponent(ArmorTag)
}
