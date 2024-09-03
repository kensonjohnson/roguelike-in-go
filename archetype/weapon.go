package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/yohamta/donburi"
)

var WeaponTag = donburi.NewTag("weapon")

func CreateNewWeapon(world donburi.World, weaponData items.WeaponData) *donburi.Entry {
	entry := CreateNewItem(world, &weaponData.ItemData)

	// Mark as a weapon
	entry.AddComponent(WeaponTag)

	// Add attack information
	entry.AddComponent(component.Attack)
	attack := component.AttackData{
		MinimumDamage: weaponData.MinimumDamage,
		MaximumDamage: weaponData.MaximumDamage,
		ToHitBonus:    weaponData.ToHitBonus,
	}
	component.Attack.SetValue(entry, attack)

	// Add action text
	entry.AddComponent(component.ActionText)
	actionText := component.ActionTextData{
		Value: weaponData.ActionText,
	}
	component.ActionText.SetValue(entry, actionText)

	return entry
}

func IsWeapon(entry *donburi.Entry) bool {
	return entry.HasComponent(WeaponTag)
}
