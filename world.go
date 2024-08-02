package main

import (
	"github.com/kensonjohnson/roguelike-game-go/components"
	"github.com/kensonjohnson/roguelike-game-go/scenes"
	"github.com/yohamta/donburi"
)

// import (
// 	"github.com/bytearena/ecs"
// 	"github.com/kensonjohnson/roguelike-game-go/assets"
// 	"github.com/kensonjohnson/roguelike-game-go/components"
// 	"github.com/kensonjohnson/roguelike-game-go/engine"
// 	"github.com/kensonjohnson/roguelike-game-go/scenes"
// )

// var (
// 	position    *ecs.Component
// 	drawable    *ecs.Component
// 	monster     *ecs.Component
// 	health      *ecs.Component
// 	meleeWeapon *ecs.Component
// 	armor       *ecs.Component
// 	name        *ecs.Component
// 	userMessage *ecs.Component
// )

// func InitializeWorld(startingLevel scenes.Level) (*ecs.Manager, map[string]ecs.Tag) {
func InitializeWorld(startingLevel scenes.Level) donburi.World {
	world := donburi.NewWorld()
	// 	manager := ecs.NewManager()
	// 	tags := make(map[string]ecs.Tag)

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	// Create the player
	entity := world.Create(components.Player, components.Position, components.Drawable, components.Fov)
	entry := world.Entry(entity)

	components.Position.SetValue(entry, components.PositionData{X: x, Y: y})

	// 	player := manager.NewComponent()
	// 	position = manager.NewComponent()
	// 	drawable = manager.NewComponent()
	// 	movable := manager.NewComponent()
	// 	monster = manager.NewComponent()
	// 	health = manager.NewComponent()
	// 	meleeWeapon = manager.NewComponent()
	// 	armor = manager.NewComponent()
	// 	name = manager.NewComponent()
	// 	userMessage = manager.NewComponent()

	// 	manager.NewEntity().
	// 		AddComponent(player, components.Player{}).
	// 		AddComponent(drawable, &components.Drawable{
	// 			Image: assets.Player,
	// 		}).
	// 		AddComponent(movable, components.Movable{}).
	// 		AddComponent(position, &components.Position{
	// 			X: x,
	// 			Y: y,
	// 		}).
	// 		AddComponent(health, &components.Health{
	// 			MaxHealth:     30,
	// 			CurrentHealth: 30,
	// 		}).
	// 		AddComponent(meleeWeapon, &components.MeleeWeapon{
	// 			Name:          "Battle Axe",
	// 			MinimumDamage: 10,
	// 			MaximumDamage: 20,
	// 			ToHitBonus:    3,
	// 		}).
	// 		AddComponent(armor, &components.Armor{
	// 			Name:       "Plate Armor",
	// 			Defense:    15,
	// 			ArmorClass: 18,
	// 		}).
	// 		AddComponent(name, &components.Name{Label: "Player"}).
	// 		AddComponent(userMessage, &components.UserMessage{
	// 			AttackMessage:    "",
	// 			DeadMessage:      "",
	// 			GameStateMessage: "",
	// 		})

	// 	for _, room := range startingLevel.Rooms {
	// 		if room.X1 != startingRoom.X1 {
	// 			mX, mY := room.Center()

	// 			// Flip a coin to see what to add...
	// 			mobSpawn := engine.GetDiceRoll(2)

	// 			if mobSpawn == 1 {
	// 				manager.NewEntity().
	// 					AddComponent(monster, &components.Monster{}).
	// 					AddComponent(drawable, &components.Drawable{
	// 						Image: assets.Orc,
	// 					}).
	// 					AddComponent(position, &components.Position{
	// 						X: mX,
	// 						Y: mY,
	// 					}).
	// 					AddComponent(health, &components.Health{
	// 						MaxHealth:     30,
	// 						CurrentHealth: 30,
	// 					}).
	// 					AddComponent(meleeWeapon, &components.MeleeWeapon{
	// 						Name:          "Machete",
	// 						MinimumDamage: 4,
	// 						MaximumDamage: 8,
	// 						ToHitBonus:    1,
	// 					}).
	// 					AddComponent(armor, &components.Armor{
	// 						Name:       "Leather",
	// 						Defense:    5,
	// 						ArmorClass: 6,
	// 					}).
	// 					AddComponent(name, &components.Name{Label: "Orc"}).
	// 					AddComponent(userMessage, &components.UserMessage{
	// 						AttackMessage:    "",
	// 						DeadMessage:      "",
	// 						GameStateMessage: "",
	// 					})
	// 			} else {
	// 				manager.NewEntity().
	// 					AddComponent(monster, &components.Monster{}).
	// 					AddComponent(drawable, &components.Drawable{
	// 						Image: assets.Skelly,
	// 					}).
	// 					AddComponent(position, &components.Position{
	// 						X: mX,
	// 						Y: mY,
	// 					}).
	// 					AddComponent(health, &components.Health{
	// 						MaxHealth:     10,
	// 						CurrentHealth: 10,
	// 					}).
	// 					AddComponent(meleeWeapon, &components.MeleeWeapon{
	// 						Name:          "Short Sword",
	// 						MinimumDamage: 2,
	// 						MaximumDamage: 6,
	// 						ToHitBonus:    0,
	// 					}).
	// 					AddComponent(armor, &components.Armor{
	// 						Name:       "Bone",
	// 						Defense:    3,
	// 						ArmorClass: 4,
	// 					}).
	// 					AddComponent(name, &components.Name{Label: "Skeleton"}).
	// 					AddComponent(userMessage, &components.UserMessage{
	// 						AttackMessage:    "",
	// 						DeadMessage:      "",
	// 						GameStateMessage: "",
	// 					})
	// 			}

	// 		}
	// 	}

	// 	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	// 	tags["players"] = players

	// 	drawables := ecs.BuildTag(drawable, position)
	// 	tags["drawables"] = drawables

	// 	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	// 	tags["monsters"] = monsters

	// 	messengers := ecs.BuildTag(userMessage)
	// 	tags["messengers"] = messengers

	return world
}
