package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

var ArmorTag = donburi.NewTag("armor")

func CreateNewArmor(world donburi.World, armorData items.ArmorData) *donburi.Entry {
	entry := CreateNewItem(world, &armorData.ItemData)

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
