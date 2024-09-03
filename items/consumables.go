package items

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type ConsumableData struct {
	ItemData
	AmountHeal int
}

type ConsumablesId int

type consumables struct {
	HealthPotion      ConsumableData
	GreatHealthPotion ConsumableData
	RoyalHealthPotion ConsumableData

	Apple  ConsumableData
	Bread  ConsumableData
	Carrot ConsumableData
	Cheese ConsumableData
	Egg    ConsumableData
	Fish   ConsumableData
	Ham    ConsumableData
	Milk   ConsumableData
	Pear   ConsumableData
	Steak  ConsumableData
}

var Consumables = consumables{
	HealthPotion: ConsumableData{
		ItemData: ItemData{
			Name:   "Health Potion",
			Sprite: assets.MustBeValidImage(assets.WorldHealthPotion, "WorldHealthPotion"),
		},
		AmountHeal: 10,
	},

	GreatHealthPotion: ConsumableData{
		ItemData: ItemData{
			Name:   "Great Heath Potion",
			Sprite: assets.MustBeValidImage(assets.WorldGreatHealthPotion, "WorldGreatHealthPotion"),
		},
		AmountHeal: 20,
	},

	RoyalHealthPotion: ConsumableData{
		ItemData: ItemData{
			Name:   "Royal Heath Potion",
			Sprite: assets.MustBeValidImage(assets.WorldRoyalHealthPotion, "WorldRoyalHealthPotion"),
		},
		AmountHeal: 40,
	},

	Apple: ConsumableData{
		ItemData: ItemData{
			Name:   "Apple",
			Sprite: assets.MustBeValidImage(assets.Apple, "Apple"),
		},
		AmountHeal: 3,
	},
	Bread: ConsumableData{
		ItemData: ItemData{
			Name:   "Bread",
			Sprite: assets.MustBeValidImage(assets.Bread, "Bread"),
		},
		AmountHeal: 5,
	},
	Carrot: ConsumableData{
		ItemData: ItemData{
			Name:   "Carrot",
			Sprite: assets.MustBeValidImage(assets.Carrot, "Carrot"),
		},
		AmountHeal: 2,
	},
	Cheese: ConsumableData{
		ItemData: ItemData{
			Name:   "Cheese",
			Sprite: assets.MustBeValidImage(assets.Cheese, "Cheese"),
		},
		AmountHeal: 6,
	},
	Egg: ConsumableData{
		ItemData: ItemData{
			Name:   "Egg",
			Sprite: assets.MustBeValidImage(assets.Egg, "Egg"),
		},
		AmountHeal: 6,
	},
	Fish: ConsumableData{
		ItemData: ItemData{
			Name:   "Fish",
			Sprite: assets.MustBeValidImage(assets.Fish, "Fish"),
		},
		AmountHeal: 9,
	},
	Ham: ConsumableData{
		ItemData: ItemData{
			Name:   "Ham",
			Sprite: assets.MustBeValidImage(assets.Ham, "Ham"),
		},
		AmountHeal: 12,
	},
	Milk: ConsumableData{
		ItemData: ItemData{
			Name:   "Milk",
			Sprite: assets.MustBeValidImage(assets.Milk, "Milk"),
		},
		AmountHeal: 6,
	},
	Pear: ConsumableData{
		ItemData: ItemData{
			Name:   "Pear",
			Sprite: assets.MustBeValidImage(assets.Pear, "Pear"),
		},
		AmountHeal: 3,
	},
	Steak: ConsumableData{
		ItemData: ItemData{
			Name:   "Steak",
			Sprite: assets.MustBeValidImage(assets.Steak, "Steak"),
		},
		AmountHeal: 15,
	},
}
