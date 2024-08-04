package gear

import "github.com/kensonjohnson/roguelike-game-go/component"

type armor struct {
	BurlapSack   component.ArmorData
	LeatherArmor component.ArmorData
	Bone         component.ArmorData
	PlateArmor   component.ArmorData
}

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

		Bone: component.ArmorData{
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
