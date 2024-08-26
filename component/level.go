package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
	"github.com/yohamta/donburi"
)

type MapTile struct {
	TileX      int
	TileY      int
	PixelX     int
	PixelY     int
	Blocked    bool
	Image      *ebiten.Image
	IsRevealed bool
	TileType   tileType
}

type tileType int

const levelHeight = config.ScreenHeight - config.UIHeight

const (
	WALL tileType = iota
	FLOOR
	STAIR_UP
	STAIR_DOWN
)

type LevelData struct {
	Tiles  []*MapTile
	Rooms  []engine.Rect
	Redraw bool
}

// Gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (l *LevelData) GetIndexFromXY(x int, y int) int {
	return (y * config.ScreenWidth) + x
}

// Gets the pointer to the MapTile at the given (x, y) coordinate.
// This coordinate is logical tiles, not pixels.
func (l *LevelData) GetFromXY(x, y int) *MapTile {
	if len(l.Tiles) == 0 {
		logger.Fatal("levelData.Tiles has no length")
	}
	return l.Tiles[(y*config.ScreenWidth)+x]
}

// Needed for fov package
func (level LevelData) InBounds(x, y int) bool {
	if x < 0 || x > config.ScreenWidth || y < 0 || y > levelHeight {
		return false
	}
	return true
}

// Needed for fov package
func (level LevelData) IsOpaque(x, y int) bool {
	return level.GetFromXY(x, y).TileType == WALL
}

var Level = donburi.NewComponentType[LevelData]()
