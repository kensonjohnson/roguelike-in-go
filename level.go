package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

// Gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// Creates a map of all tiles as a baseline to carve out a level.
func (level *Level) createTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenWidth*gd.ScreenHeight)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TileHeight,
				Blocked: true,
				Image:   wall,
			}

			tiles[index] = tile

		}
	}
	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
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
	tiles := level.createTiles()
	level.Tiles = tiles

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := GetDiceRoll(gd.ScreenHeight-h-1) - 1

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
			level.Rooms = append(level.Rooms, newRoom)
		}
	}

}

func NewLevel() Level {
	level := Level{}
	rooms := make([]Rect, 0)
	level.Rooms = rooms
	level.GenerateLevelTiles()

	return level
}
