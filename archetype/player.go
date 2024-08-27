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

var PlayerTag = donburi.NewTag("player")

func CreateNewPlayer(
	world donburi.World,
	level *component.LevelData,
	startingRoom engine.Rect,
	weaponId weapons.WeaponId,
	armorId armors.ArmorId,
) {
	player := world.Entry(world.Create(
		PlayerTag,
		component.Position,
		component.Sprite,
		component.Name,
		component.Fov,
		component.Equipment,
		component.Health,
		component.UserMessage,
	))

	// Set starting position
	startingX, startingY := startingRoom.Center()
	position := component.PositionData{
		X: startingX,
		Y: startingY,
	}
	component.Position.SetValue(player, position)

	// Update player's field of view
	vision := component.FovData{VisibleTiles: fov.New()}
	vision.VisibleTiles.Compute(level, startingX, startingY, 8)
	component.Fov.SetValue(player, vision)

	// Set sprite
	sprite := component.SpriteData{
		Image: assets.Player,
	}
	component.Sprite.SetValue(player, sprite)

	// Set name
	name := component.NameData{Value: "Player"}
	component.Name.SetValue(player, name)

	// Set health
	health := component.HealthData{
		MaxHealth:     30,
		CurrentHealth: 30,
	}
	component.Health.SetValue(player, health)

	// Add gear
	equipment := component.EquipmentData{
		Weapon: CreateNewWeapon(world, weaponId),
		Armor:  CreateNewArmor(world, armorId),
	}
	component.Equipment.SetValue(player, equipment)

	// Set default messages
	component.UserMessage.SetValue(
		player,
		component.UserMessageData{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		},
	)
}
