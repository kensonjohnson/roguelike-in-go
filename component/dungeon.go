package component

import "github.com/yohamta/donburi"

type DungeonData struct {
	Name         string
	Levels       []*LevelData
	CurrentLevel *LevelData
}

var Dungeon = donburi.NewComponentType[DungeonData]()
