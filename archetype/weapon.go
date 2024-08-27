package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/weapons"
	"github.com/yohamta/donburi"
)

var WeaponTag = donburi.NewTag("weapon")

func CreateNewWeapon(world donburi.World, weaponId weapons.WeaponId) *donburi.Entry {
	weapon := world.Entry(world.Create(
		WeaponTag,
		component.Name,
		component.Sprite,
		component.Attack,
		component.ActionText,
	))

	weaponData := weapons.Data[weaponId]

	name := component.NameData{
		Value: weaponData.Name,
	}
	component.Name.SetValue(weapon, name)

	sprite := component.SpriteData{
		Image: weaponData.Sprite,
	}
	component.Sprite.SetValue(weapon, sprite)

	attack := component.AttackData{
		MinimumDamage: weaponData.MinimumDamage,
		MaximumDamage: weaponData.MaximumDamage,
		ToHitBonus:    weaponData.ToHitBonus,
	}
	component.Attack.SetValue(weapon, attack)

	actionText := component.ActionTextData{
		Value: weaponData.ActionText,
	}
	component.ActionText.SetValue(weapon, actionText)

	return weapon
}

func IsWeapon(entry *donburi.Entry) bool {
	return entry.HasComponent(WeaponTag)
}
