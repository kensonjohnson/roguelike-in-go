package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
)

func CreateNewPlayer(
	world donburi.World,
	weapon items.WeaponData,
	armorId items.ArmorData,
) *donburi.Entry {
	player := world.Entry(world.Create(
		tags.PlayerTag,
		component.Position,
		component.Sprite,
		component.Name,
		component.Fov,
		component.Equipment,
		component.Inventory,
		component.Wallet,
		component.Health,
		component.UserMessage,
		component.Attack,
		component.ActionText,
		component.Defense,
	))

	position := component.PositionData{}
	component.Position.SetValue(player, position)

	// Update player's field of view
	vision := component.FovData{VisibleTiles: fov.New()}
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
		Weapon: CreateNewWeapon(world, items.Weapons.BattleAxe),
		Armor:  CreateNewArmor(world, items.Armor.PlateArmor),
	}
	component.Equipment.SetValue(player, equipment)

	// Setup inventory
	inventory := component.NewInventory(30)
	component.Inventory.SetValue(player, inventory)

	wallet := component.WalletData{}
	component.Wallet.SetValue(player, wallet)

	// Set default messages
	component.UserMessage.SetValue(
		player,
		component.UserMessageData{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		},
	)

	// Total up all of the attack values
	// Right now, only the weapon contributes to attack.
	// TODO: Add up all attack values from all equipment
	attack := component.Attack.Get(equipment.Weapon)
	component.Attack.SetValue(player, *attack)

	// Set action text for equiped weapon
	actionText := component.ActionText.Get(equipment.Weapon)
	component.ActionText.SetValue(player, *actionText)

	// Total all of the defense values
	// Right now, only the armor contributes to defense.
	// TODO: Add up all defense values from all equipment
	defense := component.Defense.Get(equipment.Armor)
	component.Defense.SetValue(player, *defense)

	return player
}
