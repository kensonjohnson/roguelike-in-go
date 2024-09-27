package items

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/internal/engine"
)

type ValuableData struct {
	ItemData
	Value int
}

type valuables struct {
	Alcohol *ValuableData
}

var Valuables valuables = valuables{
	Alcohol: &ValuableData{
		ItemData: ItemData{
			Name:        "Alcohol",
			Sprite:      assets.WorldAlcohol,
			Description: "A bottle of alcohol. Should be well aged by now.",
		},
		Value: 20,
	},
}

func (v valuables) SmallCoin() *ValuableData {
	value := engine.GetRandomBetween(3, 11)
	return &ValuableData{
		ItemData: ItemData{
			Name:   "some coins",
			Sprite: assets.WorldSmallCoin,
		},
		Value: value,
	}
}

func (v valuables) CoinStack() *ValuableData {
	value := engine.GetRandomBetween(15, 35)
	return &ValuableData{
		ItemData: ItemData{
			Name:   "a stack of coins",
			Sprite: assets.WorldCoinStack,
		},
		Value: value,
	}
}
