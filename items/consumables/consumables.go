package consumables

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
)

type consumable struct {
	Name       string
	AmountHeal int
	Sprite     *ebiten.Image
}

type ConsumablesId int

const (
	HealthPotion ConsumablesId = iota
	GreatHealthPotion
	RoyalHealthPotion

	Apple
	Bread
	Carrot
	Cheese
	Egg
	Fish
	Ham
	Milk
	Pear
	Steak
)

var Data = []*consumable{
	{
		Name:       "Health Potion",
		AmountHeal: 10,
		Sprite:     mustBeValidImage(assets.WorldHealthPotion, "WorldHealthPotion"),
	},

	{
		Name:       "Great Heath Potion",
		AmountHeal: 20,
		Sprite:     mustBeValidImage(assets.WorldGreatHealthPotion, "WorldGreatHealthPotion"),
	},

	{
		Name:       "Royal Heath Potion",
		AmountHeal: 40,
		Sprite:     mustBeValidImage(assets.WorldRoyalHealthPotion, "WorldRoyalHealthPotion"),
	},
}

func mustBeValidImage(image *ebiten.Image, name string) *ebiten.Image {
	if image == nil {
		logger.ErrorLogger.Panicf("%s asset not loaded!", name)
	}
	return image
}
