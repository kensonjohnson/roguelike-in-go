package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	position   *ecs.Component
	renderable *ecs.Component
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	manager := ecs.NewManager()
	tags := make(map[string]ecs.Tag)

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImage,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return manager, tags
}
