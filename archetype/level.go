package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/assets"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/engine"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
	"github.com/kensonjohnson/roguelike-game-go/items/armors"
	"github.com/kensonjohnson/roguelike-game-go/items/consumables"
	"github.com/kensonjohnson/roguelike-game-go/items/weapons"
	"github.com/yohamta/donburi"
)

const levelHeight = config.ScreenHeight - config.UIHeight

var LevelTag = donburi.NewTag("level")

// Creates a new Dungeon
func GenerateLevel(world donburi.World) *component.LevelData {
	entry := world.Entry(world.Create(
		LevelTag,
		component.Level,
	))

	level := GenerateLevelTiles()
	level.Redraw = true
	seedRooms(world, &level)
	component.Level.SetValue(entry, level)

	return &level
}

// Creates a new Dungeon Level Map.
func GenerateLevelTiles() component.LevelData {
	MIN_SIZE := 6
	MAX_SIZE := 13
	MAX_ROOMS := 30

	level := component.LevelData{}
	tiles := createTiles(level)
	level.Tiles = tiles
	containsRooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := engine.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := engine.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := engine.GetDiceRoll(config.ScreenWidth - w - 1)
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
			createRoom(level, newRoom)

			if containsRooms {
				newX, newY := newRoom.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()

				coinflip := engine.GetDiceRoll(2)

				if coinflip == 2 {
					createHorizontalTunnel(level, prevX, newX, prevY)
					createVerticalTunnel(level, prevY, newY, newX)
				} else {
					createHorizontalTunnel(level, prevX, newX, newY)
					createVerticalTunnel(level, prevY, newY, prevX)
				}
			}
			level.Rooms = append(level.Rooms, newRoom)
			containsRooms = true
		}
	}

	return level
}

// Creates a map of all tiles as a baseline to carve out a level.
func createTiles(level component.LevelData) []*component.MapTile {
	levelHeight := config.ScreenHeight - config.UIHeight
	tiles := make([]*component.MapTile, config.ScreenWidth*levelHeight)
	index := 0

	for x := 0; x < config.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			tile := component.MapTile{
				TileX:      x,
				TileY:      y,
				PixelX:     x * config.TileWidth,
				PixelY:     y * config.TileHeight,
				Blocked:    true,
				Image:      assets.Wall,
				IsRevealed: false,
				TileType:   component.WALL,
			}

			tiles[index] = &tile

		}
	}
	return tiles
}

// Carves out a room in the map of tiles
func createRoom(level component.LevelData, room engine.Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = component.FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
}

// Carves out a tunnel to another room horizontally
func createHorizontalTunnel(level component.LevelData, x1 int, x2 int, y int) {

	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < config.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = component.FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
}

// Carves out a tunnel to another room vertically
func createVerticalTunnel(level component.LevelData, y1 int, y2 int, x int) {

	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)

		if index > 0 && index < config.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = component.FLOOR
			level.Tiles[index].Image = assets.Floor
		}
	}
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

func seedRooms(world donburi.World, level *component.LevelData) {
	for index, room := range level.Rooms {
		if index == 0 {
			CreateNewPlayer(
				world,
				level,
				room,
				weapons.BattleAxe,
				armors.PlateArmor,
			)
		} else {
			CreateMonster(world, level, room)
			addRandomPickupsToRoom(world, level, room, 3)
		}
	}
	exitRoomIndex := engine.GetDiceRoll(len(level.Rooms) - 1)
	exitRoom := level.Rooms[exitRoomIndex]
	exitTile := level.GetFromXY(exitRoom.X1+1, exitRoom.Y1+1)
	exitTile.Image = assets.StairsDown
	exitTile.Blocked = false
	exitTile.TileType = component.STAIR_DOWN
}

func addRandomPickupsToRoom(
	world donburi.World,
	level *component.LevelData,
	room engine.Rect,
	generosity int,
) {
	// TODO: create a random distribution of items
	// TODO: add a random chance for an item to appear
	switch d10Roll := engine.GetDiceRoll(10); {
	case d10Roll <= generosity:
		width := room.X2 - room.X1 - 2
		height := room.Y2 - room.Y1 - 2
		for i := 0; i < d10Roll; i++ {
			offsetX := engine.GetRandomInt(width)
			offsetY := engine.GetRandomInt(height)
			x := room.X1 + offsetX + 1
			y := room.Y1 + offsetY + 1
			tile := level.GetFromXY(x, y)
			if tile.Blocked {
				continue
			}
			spotTaken := false
			ConsumableTag.Each(world, func(entry *donburi.Entry) {
				position := component.Position.Get(entry)
				if position.X == x && position.Y == y {
					spotTaken = true
				}
			})
			if spotTaken {
				continue
			}
			potion := CreateNewConsumable(world, consumables.HealthPotion)
			err := PlaceItemInWorld(potion, x, y, true)
			if err != nil {
				logger.ErrorLogger.Panic("Failed to place consumable in the world")
			}
		}
	}

}
