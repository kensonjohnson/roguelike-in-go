package weapons

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
)

type weaponData struct {
	Name          string
	ActionText    string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
	Sprite        *ebiten.Image
}

type WeaponId int

const (
	ShortSword WeaponId = iota
	BattleAxe
)

// MAKE SURE THAT THIS NUMBER MATCHES THE NUMBER OF WEAPONS DEFINED!
var Data = []weaponData{
	{
		Name:          "Short Sword",
		ActionText:    "swings a short sword at",
		MinimumDamage: 2,
		MaximumDamage: 6,
		ToHitBonus:    0,
		Sprite:        mustBeValidImage(assets.ShortSword, "ShortSword"),
	},
	{
		Name:          "Battle Axe",
		ActionText:    "cleaves a battle axe at",
		MinimumDamage: 10,
		MaximumDamage: 20,
		ToHitBonus:    3,
		Sprite:        mustBeValidImage(assets.BattleAxe, "BattleAxe"),
	},
}

func mustBeValidImage(image *ebiten.Image, name string) *ebiten.Image {
	if image == nil {
		logger.ErrorLogger.Panicf("%s asset not loaded!", name)
	}
	return image
}
