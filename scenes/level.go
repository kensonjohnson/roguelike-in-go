package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/components"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/norendren/go-fov/fov"
	"github.com/yohamta/donburi"
)

type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	Image      *ebiten.Image
	IsRevealed bool
	TileType   TileType
}

type TileType int

var (
	levelHeight int = 0
)

const (
	WALL TileType = iota
	FLOOR
)

type Level struct {
	Tiles         []*MapTile
	Rooms         []engine.Rect
	PlayerVisible *fov.View
}

func NewLevel() Level {
	level := Level{}
	rooms := make([]engine.Rect, 0)
	level.Rooms = rooms
	level.GenerateLevelTiles()
	level.PlayerVisible = fov.New()

	return level
}

// Gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	return (y * config.Config.ScreenWidth) + x
}

// Creates a map of all tiles as a baseline to carve out a level.
func (level *Level) createTiles() []*MapTile {
	tiles := make([]*MapTile, config.Config.ScreenWidth*levelHeight)
	index := 0

	for x := 0; x < config.Config.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			tile := MapTile{
				PixelX:     x * config.Config.TileWidth,
				PixelY:     y * config.Config.TileHeight,
				Blocked:    true,
				Image:      assets.Wall,
				IsRevealed: false,
				TileType:   WALL,
			}

			tiles[index] = &tile

		}
	}
	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image, world donburi.World) {

	if entry, ok := components.Player.First(world); ok {
		pos := components.Position.Get(entry)
		level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
	}

	for x := 0; x < config.Config.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			IsVisible := level.PlayerVisible.IsVisible(x, y)
			if IsVisible {
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, options)
				level.Tiles[idx].IsRevealed = true
			} else if tile.IsRevealed {
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				options.ColorScale.ScaleAlpha(0.35)
				screen.DrawImage(tile.Image, options)
			}
		}
	}
}

func (level *Level) createRoom(room engine.Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
}

// Creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	levelHeight = config.Config.ScreenHeight - config.Config.UIHeight
	tiles := level.createTiles()
	level.Tiles = tiles
	containsRooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := engine.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := engine.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := engine.GetDiceRoll(config.Config.ScreenWidth - w - 1)
		y := engine.GetDiceRoll(levelHeight - h - 1)

		newRoom := engine.NewRect(x, y, w, h)
		okToAdd := true

		for _, otherRoom := range level.Rooms {
			if newRoom.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(newRoom)

			if containsRooms {
				newX, newY := newRoom.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()

				coinflip := engine.GetDiceRoll(2)

				if coinflip == 2 {
					level.createHorizontalTunnel(prevX, newX, prevY)
					level.createVerticalTunnel(prevY, newY, newX)
				} else {
					level.createHorizontalTunnel(prevX, newX, newY)
					level.createVerticalTunnel(prevY, newY, prevX)
				}
			}
			level.Rooms = append(level.Rooms, newRoom)
			containsRooms = true

		}
	}
}

func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {

	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < config.Config.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {

	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)

		if index > 0 && index < config.Config.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
}

func (level Level) InBounds(x, y int) bool {

	if x < 0 || x > config.Config.ScreenWidth || y < 0 || y > levelHeight {
		return false
	}
	return true
}

func (level Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == WALL
}

// Returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
