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

type consumablesList []*consumable

type ConsumablesId int

const (
	HealthPotion ConsumablesId = iota
)

// MAKE SURE THAT THIS NUMBER MATCHES THE NUMBER OF WEAPONS DEFINED!
var Data consumablesList = make(consumablesList, 1)

func init() {

	Data[HealthPotion] = &consumable{
		Name:       "Health Potion",
		AmountHeal: 10,
		Sprite:     mustBeValidImage(assets.WorldHealthPotion, "WorldHealthPotion"),
	}

}

func mustBeValidImage(image *ebiten.Image, name string) *ebiten.Image {
	if image == nil {
		logger.ErrorLogger.Panicf("%s asset not loaded!", name)
	}
	return image
}
