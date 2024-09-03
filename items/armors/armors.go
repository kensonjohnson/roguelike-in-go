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

var Data = []armorData{
	{
		Name:       "Linen Shirt",
		Defense:    1,
		ArmorClass: 1,
		Sprite:     mustBeValidImage(assets.LinenShirt, "LinenShirt"),
	},
	{
		Name:       "Padded Armor",
		Defense:    5,
		ArmorClass: 6,
		Sprite:     mustBeValidImage(assets.PaddedArmor, "PaddedArmor"),
	},
	{
		Name:       "Bone",
		Defense:    3,
		ArmorClass: 4,
		Sprite:     mustBeValidImage(assets.Bones, "Bones"),
	},
	{
		Name:       "Plate Armor",
		Defense:    15,
		ArmorClass: 18,
		Sprite:     mustBeValidImage(assets.PlateArmor, "PlateArmor"),
	},
}

func mustBeValidImage(image *ebiten.Image, name string) *ebiten.Image {
	if image == nil {
		logger.ErrorLogger.Panicf("%s asset not loaded!", name)
	}
	return image
}
