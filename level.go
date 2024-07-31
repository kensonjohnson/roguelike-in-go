package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/norendren/go-fov/fov"
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
	floor       *ebiten.Image
	wall        *ebiten.Image
	levelHeight int = 0
)

func loadTileImages() {
	if floor != nil && wall != nil {
		return
	}

	imgSource, err := assets.ReadFile("assets/floor.png")
	if err != nil {
		log.Fatal(err)
	}

	floor, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(imgSource))
	if err != nil {
		log.Fatal(err)
	}

	imgSource, err = assets.ReadFile("assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}

	wall, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(imgSource))
	if err != nil {
		log.Fatal(err)
	}
}

const (
	WALL TileType = iota
	FLOOR
)

type Level struct {
	Tiles         []*MapTile
	Rooms         []Rect
	PlayerVisible *fov.View
}

func NewLevel() Level {
	level := Level{}
	loadTileImages()
	rooms := make([]Rect, 0)
	level.Rooms = rooms
	level.GenerateLevelTiles()
	level.PlayerVisible = fov.New()

	return level
}

// Gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// Creates a map of all tiles as a baseline to carve out a level.
func (level *Level) createTiles() []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, gd.ScreenWidth*levelHeight)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			tile := MapTile{
				PixelX:     x * gd.TileWidth,
				PixelY:     y * gd.TileHeight,
				Blocked:    true,
				Image:      wall,
				IsRevealed: false,
				TileType:   WALL,
			}

			tiles[index] = &tile

		}
	}
	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()

	for x := 0; x < gd.ScreenWidth; x++ {
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

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = floor
		}
	}
}

// Creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	levelHeight = gd.ScreenHeight - gd.UIHeight
	tiles := level.createTiles()
	level.Tiles = tiles
	containsRooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(levelHeight - h - 1)

		newRoom := NewRect(x, y, w, h)
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

				coinflip := GetDiceRoll(2)

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
	gd := NewGameData()
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = floor
		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)

		if index > 0 && index < gd.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = floor
		}
	}
}

func (level Level) InBounds(x, y int) bool {
	gd := NewGameData()
	if x < 0 || x > gd.ScreenWidth || y < 0 || y > levelHeight {
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
