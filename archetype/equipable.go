package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/yohamta/donburi"
)

func IsEquipable(entry *donburi.Entry) bool {
	return entry.HasComponent(tags.EquipableTag)
}
