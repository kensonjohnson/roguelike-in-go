package consumables

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type consumable struct {
	Name       string
	AmountHeal int
	Sprite     *ebiten.Image
}

type consumablesList []consumable

type ConsumablesId int

const (
	HealthPotion ConsumablesId = iota
)

// MAKE SURE THAT THIS NUMBER MATCHES THE NUMBER OF WEAPONS DEFINED!
var Data consumablesList = make(consumablesList, 1)

func init() {

	Data[HealthPotion] = consumable{
		Name:       "Health Potion",
		AmountHeal: 10,
		Sprite:     assets.HealthPotion,
	}

}
