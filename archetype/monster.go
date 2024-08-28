package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/kensonjohnson/roguelike-game-go/items/armors"
	"github.com/kensonjohnson/roguelike-game-go/items/weapons"
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
		component.Equipment,
		component.Health,
		component.UserMessage,
		component.Discoverable,
		component.Attack,
		component.ActionText,
		component.Defense,
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
	equipment := component.EquipmentData{}
	coinflip := engine.GetDiceRoll(2)
	if coinflip == 2 {
		sprite.Image = assets.Orc
		name.Value = "Orc"
		equipment.Armor = CreateNewArmor(world, armors.PaddedArmor)
		equipment.Weapon = CreateNewWeapon(world, weapons.ShortSword)
	} else {
		sprite.Image = assets.Skelly
		name.Value = "Skeleton"
		equipment.Armor = CreateNewArmor(world, armors.Bones)
		equipment.Weapon = CreateNewWeapon(world, weapons.ShortSword)
	}
	component.Sprite.SetValue(monster, sprite)
	component.Name.SetValue(monster, name)
	component.Equipment.SetValue(monster, equipment)

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

	// Total up all of the attack values
	// Right now, only the weapon contributes to attack.
	// TODO: Add up all attack values from all equipment
	attack := component.Attack.Get(equipment.Weapon)
	component.Attack.SetValue(monster, *attack)

	// Set action text for equiped weapon
	actionText := component.ActionText.Get(equipment.Weapon)
	component.ActionText.SetValue(monster, *actionText)

	// Total all of the defense values
	// Right now, only the armor contributes to defense.
	// TODO: Add up all defense values from all equipment
	defense := component.Defense.Get(equipment.Armor)
	component.Defense.SetValue(monster, *defense)
}
