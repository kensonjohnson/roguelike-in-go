package armors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
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
		Sprite:     mustBeValidImage(assets.LinenShirt, "LinenShirt"),
	}

	Data[PaddedArmor] = armorData{
		Name:       "Padded Armor",
		Defense:    5,
		ArmorClass: 6,
		Sprite:     mustBeValidImage(assets.PaddedArmor, "PaddedArmor"),
	}

	Data[Bones] = armorData{
		Name:       "Bone",
		Defense:    3,
		ArmorClass: 4,
		Sprite:     mustBeValidImage(assets.Bones, "Bones"),
	}

	Data[PlateArmor] = armorData{
		Name:       "Plate Armor",
		Defense:    15,
		ArmorClass: 18,
		Sprite:     mustBeValidImage(assets.PlateArmor, "PlateArmor"),
	}
}

func mustBeValidImage(image *ebiten.Image, name string) *ebiten.Image {
	if image == nil {
		logger.ErrorLogger.Panicf("%s asset not loaded!", name)
	}
	return image
}
