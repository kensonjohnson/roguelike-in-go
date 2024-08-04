package action

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/system/combat"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func TakePlayerAction(ecs *ecs.ECS) bool {
	turnTaken := false

	// Grab current level
	dungeonEntry := archetype.MustFindDungeon(ecs.World)
	level := component.Dungeon.Get(dungeonEntry).CurrentLevel

	// Grab player data
	playerEntry := archetype.PlayerTag.MustFirst(ecs.World)
	position := component.Position.Get(playerEntry)
	vision := component.Fov.Get(playerEntry)

	// Capture any keypresses we care about
	moveX := 0
	moveY := 0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveX = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveX = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveY = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveY = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}

	if moveX == 0 && moveY == 0 && !turnTaken {
		return false
	}

	// Attempt to move
	tile := level.GetFromXY(position.X+moveX, position.Y+moveY)
	if !tile.Blocked {
		// We can move the player
		// Free up the tile we are currently on and block the one we're going to
		level.GetFromXY(position.X, position.Y).Blocked = false
		tile.Blocked = true
		// Update the player's position
		position.X += moveX
		position.Y += moveY
		// Update the player's field of view
		vision.VisibleTiles.Compute(level, position.X, position.Y, 8)
	} else if tile.TileType != component.WALL {
		// Not a wall, so it must be an enemy. Attack!
		// Find the monster in the direction we're pointing
		enemyPosition := component.PositionData{
			X: position.X + moveX,
			Y: position.Y + moveY,
		}
		var monsterEntry *donburi.Entry
		archetype.MonsterTag.Each(ecs.World, func(entry *donburi.Entry) {
			position := component.Position.Get(entry)
			if position.IsEqual(&enemyPosition) {
				monsterEntry = entry
			}
		})
		combat.AttackSystem(playerEntry, monsterEntry)
	}

	return true
}
