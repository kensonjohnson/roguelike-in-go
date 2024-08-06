package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/component/gear"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
)

var MonsterTag = donburi.NewTag("monster")

func CreateMonster(world donburi.World, level *component.LevelData, room engine.Rect) {
	monster := world.Entry(world.Create(
		MonsterTag,
		component.Position,
		component.Sprite,
		component.Name,
		component.Fov,
		component.Armor,
		component.Weapon,
		component.Health,
		component.UserMessage,
		component.Discoverable,
	))

	// Set position
	startingX, startingY := room.Center()
	position := component.PositionData{
		X: startingX,
		Y: startingY,
	}
	component.Position.SetValue(monster, position)

	// Set monster's vision
	vision := component.FovData{VisibleTiles: fov.New()}
	vision.VisibleTiles.Compute(level, startingX, startingY, 8)
	component.Fov.SetValue(monster, vision)

	// Set sprite, name, and gear
	sprite := component.SpriteData{}
	name := component.NameData{}
	var armor component.ArmorData
	var weapon component.WeaponData
	coinflip := engine.GetDiceRoll(2)
	if coinflip == 2 {
		sprite.Image = assets.Orc
		name.Label = "Orc"
		armor = gear.Armor.LeatherArmor
		weapon = gear.Weapons.Machete
	} else {
		sprite.Image = assets.Skelly
		name.Label = "Skeleton"
		armor = gear.Armor.Bone
		weapon = gear.Weapons.ShortSword
	}
	component.Sprite.SetValue(monster, sprite)
	component.Name.SetValue(monster, name)
	component.Armor.SetValue(monster, armor)
	component.Weapon.SetValue(monster, weapon)
	component.Health.SetValue(
		monster,
		component.HealthData{
			MaxHealth:     30,
			CurrentHealth: 30,
		},
	)
	component.UserMessage.SetValue(
		monster,
		component.UserMessageData{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		},
	)
	component.Discoverable.SetValue(
		monster,
		component.DiscoverableData{SeenByPlayer: false},
	)
}
