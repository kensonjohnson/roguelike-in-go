package items

import "github.com/kensonjohnson/roguelike-game-go/component"

type armor struct {
	BurlapSack   component.ArmorData
	LeatherArmor component.ArmorData
	Bones        component.ArmorData
	PlateArmor   component.ArmorData
}

type armorId int

const (
	BurlapSack armorId = iota
	LeatherArmor
	Bones
	PlateArmor
)

var Armor armor

func init() {
	Armor = armor{
		BurlapSack: component.ArmorData{
			Name:       "Burlap Sack",
			Defense:    1,
			ArmorClass: 1,
		},

		LeatherArmor: component.ArmorData{
			Name:       "Leather",
			Defense:    5,
			ArmorClass: 6,
		},

		Bones: component.ArmorData{
			Name:       "Bone",
			Defense:    3,
			ArmorClass: 4,
		},

		PlateArmor: component.ArmorData{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		},
	}
}
