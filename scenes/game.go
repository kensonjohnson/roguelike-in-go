package scenes

import (
	"github.com/bytearena/ecs"
)

type Dungeon struct {
	Name   string
	Levels []Level
}

type Game struct {
	Map         GameMap
	World       ecs.Manager
	Turn        TurnState
	TurnCounter int
}

type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel Level
}

// Creates a new set of maps for the entire game.
func NewGameMap() GameMap {
	level := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, level)

	dungeon := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, dungeon)

	gameMap := GameMap{Dungeons: dungeons, CurrentLevel: level}
	return gameMap
}
