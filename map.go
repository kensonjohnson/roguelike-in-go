package main

type GameMap struct {
	Dungeons []Dungeon
}

// Creates a new set of maps for the entire game.
func NewGameMap() GameMap {
	level := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, level)

	dungeon := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, dungeon)

	gameMap := GameMap{Dungeons: dungeons}
	return gameMap
}
