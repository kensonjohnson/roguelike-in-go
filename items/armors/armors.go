package armors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type armorData struct {
	Name       string
	Defense    int
	ArmorClass int
	Sprite     *ebiten.Image
}

type ArmorId int

const (
	LinenShirt ArmorId = iota
	PaddedArmor
	Bones
	PlateArmor
)

type armorList []armorData

// MAKE SURE THAT THIS NUMBER MATCHES THE NUMBER OF ARMORS DEFINED!
var Data = make(armorList, 4)

func init() {
	Data[LinenShirt] = armorData{
		Name:       "Linen Shirt",
		Defense:    1,
		ArmorClass: 1,
		Sprite:     assets.LinenShirt,
	}

	Data[PaddedArmor] = armorData{
		Name:       "Padded Armor",
		Defense:    5,
		ArmorClass: 6,
		Sprite:     assets.PaddedArmor,
	}

	Data[Bones] = armorData{
		Name:       "Bone",
		Defense:    3,
		ArmorClass: 4,
		Sprite:     assets.Bones,
	}

	Data[PlateArmor] = armorData{
		Name:       "Plate Armor",
		Defense:    15,
		ArmorClass: 18,
		Sprite:     assets.PlateArmor,
	}
}
