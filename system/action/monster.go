package action

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/engine/pathing"
	"github.com/kensonjohnson/roguelike-game-go/system/combat"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func TakeMonsterAction(ecs *ecs.ECS) {
	// Grab level data
	entry := archetype.LevelTag.MustFirst(ecs.World)
	level := component.Level.Get(entry)

	// Grab player data
	playerEntry := archetype.PlayerTag.MustFirst(ecs.World)
	playerPos := component.Position.Get(playerEntry)

	archetype.MonsterTag.Each(ecs.World, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		sprite := component.Sprite.Get(entry)
		monsterVision := component.Fov.Get(entry).VisibleTiles

		// Check if the monster can see the player
		if monsterVision.IsVisible(playerPos.X, playerPos.Y) {
			if position.GetManhattanDistance(playerPos) == 1 {
				// The monster is directly next to the player. Smack him!
				combat.AttackSystem(ecs.World, entry, playerEntry)
			} else {
				astar := pathing.AStar{}
				path := astar.GetPath(level, position, playerPos)
				if len(path) > 1 {
					nextTile := level.GetFromXY(path[1].X, path[1].Y)
					if !nextTile.Blocked {
						// Update the tile this monster is blocking
						level.GetFromXY(position.X, position.Y).Blocked = false
						nextTile.Blocked = true
						sprite.OffestX = position.X - path[1].X
						sprite.OffestY = position.Y - path[1].Y
						position.X = path[1].X
						position.Y = path[1].Y
						sprite.Animating = true
						sprite.SetProgress(0)
						// Since the monster moved, update the field of view
						monsterVision.Compute(level, position.X, position.Y, 8)
					}
				}
			}
		}
	})
}
