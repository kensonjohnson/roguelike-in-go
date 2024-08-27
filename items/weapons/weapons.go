package weapons

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
)

type weaponData struct {
	Name          string
	ActionText    string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
	Sprite        *ebiten.Image
}

type weaponList []weaponData

type WeaponId int

const (
	ShortSword WeaponId = iota
	BattleAxe
)

var Data weaponList = make(weaponList, 0)

func init() {
	Data[ShortSword] = weaponData{
		Name:          "Short Sword",
		ActionText:    "swings a short sword at",
		MinimumDamage: 2,
		MaximumDamage: 6,
		ToHitBonus:    0,
		Sprite:        assets.ShortSword,
	}

	Data[BattleAxe] = weaponData{
		Name:          "Battle Axe",
		ActionText:    "cleaves a battle axe at",
		MinimumDamage: 10,
		MaximumDamage: 20,
		ToHitBonus:    3,
		Sprite:        assets.BattleAxe,
	}
}
