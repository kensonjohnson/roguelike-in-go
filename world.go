package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	position    *ecs.Component
	renderable  *ecs.Component
	monster     *ecs.Component
	health      *ecs.Component
	meleeWeapon *ecs.Component
	armor       *ecs.Component
	name        *ecs.Component
	userMessage *ecs.Component
)

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	manager := ecs.NewManager()
	tags := make(map[string]ecs.Tag)

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	skellyImg, _, err := ebitenutil.NewImageFromFile("assets/skelly.png")
	if err != nil {
		log.Fatal(err)
	}

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()
	userMessage = manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImage,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Battle Axe",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		}).
		AddComponent(armor, &Armor{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		}).
		AddComponent(name, &Name{Label: "Player"}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{}).
				AddComponent(renderable, &Renderable{
					Image: skellyImg,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
				}).
				AddComponent(health, &Health{
					MaxHealth:     10,
					CurrentHealth: 10,
				}).
				AddComponent(meleeWeapon, &MeleeWeapon{
					Name:          "Short Sword",
					MinimumDamage: 2,
					MaximumDamage: 6,
					ToHitBonus:    0,
				}).
				AddComponent(armor, &Armor{
					Name:       "Bone",
					Defense:    3,
					ArmorClass: 4,
				}).
				AddComponent(name, &Name{Label: "Skeleton"}).
				AddComponent(userMessage, &UserMessage{
					AttackMessage:    "",
					DeadMessage:      "",
					GameStateMessage: "",
				})
		}
	}

	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	tags["monsters"] = monsters

	messengers := ecs.BuildTag(userMessage)
	tags["messengers"] = messengers

	return manager, tags
}
