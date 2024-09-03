package items

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type ArmorData struct {
	ItemData
	Defense    int
	ArmorClass int
}

type armors struct {
	LinenShirt  ArmorData
	PaddedArmor ArmorData
	Bones       ArmorData
	PlateArmor  ArmorData
}

var Armor = armors{
	LinenShirt: ArmorData{
		ItemData: ItemData{
			Name:   "Linen Shirt",
			Sprite: assets.MustBeValidImage(assets.LinenShirt, "LinenShirt"),
		},
		Defense:    1,
		ArmorClass: 1,
	},
	PaddedArmor: ArmorData{
		ItemData: ItemData{
			Name:   "Padded Armor",
			Sprite: assets.MustBeValidImage(assets.PaddedArmor, "PaddedArmor"),
		},
		Defense:    5,
		ArmorClass: 6,
	},
	Bones: ArmorData{
		ItemData: ItemData{
			Name:   "Bone",
			Sprite: assets.MustBeValidImage(assets.Bones, "Bones"),
		},
		Defense:    3,
		ArmorClass: 4,
	},
	PlateArmor: ArmorData{
		ItemData: ItemData{
			Name:   "Plate Armor",
			Sprite: assets.MustBeValidImage(assets.PlateArmor, "PlateArmor"),
		},
		Defense:    15,
		ArmorClass: 18,
	},
}
