package archetype

import (
	"errors"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items/weapons"
	"github.com/yohamta/donburi"
)

var WeaponTag = donburi.NewTag("weapon")

func CreateNewWeapon(world donburi.World, weaponId weapons.WeaponId) *donburi.Entry {
	weapon := world.Entry(world.Create(
		WeaponTag,
		component.ItemId,
		component.Name,
		component.Sprite,
		component.Attack,
		component.ActionText,
	))

	weaponData := weapons.Data[weaponId]

	itemId := component.ItemIdData{
		Id: int(weaponId),
	}
	component.ItemId.SetValue(weapon, itemId)

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

func PlaceWeaponInWorld(world *donburi.World, entry *donburi.Entry, x, y int) error {
	if !IsWeapon(entry) {
		return errors.New("entry is not an Weapon Entity")
	}

	entry.AddComponent(component.Position)
	position := component.PositionData{
		X: x,
		Y: y,
	}
	component.Position.SetValue(entry, position)

	return nil
}
