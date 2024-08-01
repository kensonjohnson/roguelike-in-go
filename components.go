package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

type Monster struct{}

type Name struct {
	Label string
}

type Health struct {
	MaxHealth     int
	CurrentHealth int
}

type MeleeWeapon struct {
	Name          string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

type Armor struct {
	Name       string
	Defense    int
	ArmorClass int
}

type UserMessage struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}
