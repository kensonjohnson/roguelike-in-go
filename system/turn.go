package system

import (
	"fmt"

	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/system/action"
	"github.com/yohamta/donburi/ecs"
)

type turnState int

type TurnData struct {
	TurnState   turnState
	TurnCounter int
}

var Turn = TurnData{
	TurnState:   PlayerTurn,
	TurnCounter: 0,
}

const (
	BeforePlayerAction = iota
	UIOpen
	PlayerTurn
	MonsterTurn
	GameOver
)

func (td *TurnData) Update(ecs *ecs.ECS) {

	if td.TurnState == BeforePlayerAction {
		td.TurnCounter++
		// Check if player survived the last cycle of monster turns
		entry := tags.PlayerTag.MustFirst(ecs.World)
		playerHealth := component.Health.Get(entry)
		if playerHealth.CurrentHealth <= 0 {
			td.gameOver()
			playerMessages := component.UserMessage.Get(entry)
			playerMessages.GameStateMessage = "Game over!"
		}

		level := component.Level.Get(tags.LevelTag.MustFirst(ecs.World))
		// Remove any enemies that died during the last turn
		for entry = range tags.MonsterTag.Iter(ecs.World) {
			health := component.Health.Get(entry)
			if health.CurrentHealth <= 0 {
				position := component.Position.Get(entry)
				tile := level.GetFromXY(position.X, position.Y)
				tile.Blocked = false
				archetype.RemoveMonster(entry, ecs.World)
			}

		}

		for spriteEntry := range component.Sprite.Iter(ecs.World) {
			sprite := component.Sprite.Get(spriteEntry)
			sprite.SetProgress(float64(td.TurnCounter) / 12)

		}

		if td.TurnCounter < 13 {
			return
		}
		// Reset the progress of all sprites
		for entry = range component.Sprite.Iter(ecs.World) {
			sprite := component.Sprite.Get(entry)
			sprite.SetProgress(0)
			sprite.Animating = false
			sprite.OffestX = 0
			sprite.OffestY = 0
		}

		playerEntry := tags.PlayerTag.MustFirst(ecs.World)
		playerPosition := component.Position.Get(playerEntry)
		playerMessages := component.UserMessage.Get(playerEntry)

		for entry = range tags.PickupTag.Iter(ecs.World) {
			if !entry.HasComponent(component.Position) {
				continue
			}
			pickupPosition := component.Position.Get(entry)

			if pickupPosition.X == playerPosition.X && pickupPosition.Y == playerPosition.Y {

				// If pickup is coinage, add to wallet
				if entry.HasComponent(tags.CoinTag) {
					component.Wallet.Get(playerEntry).AddAmount(
						component.Value.Get(entry).Amount,
					)
					itemName := component.Name.Get(entry)
					playerMessages.WorldInteractionMessage = fmt.Sprintf("Picked up %s!", itemName.Value)
					ecs.World.Remove(entry.Entity())
					break
				}

				// Otherwise, must be item, place in inventory
				err := component.Inventory.Get(playerEntry).AddItem(entry)
				if err != nil {
					playerMessages.WorldInteractionMessage = "Inventory full! Can't pick up anymore items!"
				} else {
					archetype.RemoveItemFromWorld(entry)
					itemName := component.Name.Get(entry)
					playerMessages.WorldInteractionMessage = fmt.Sprintf("Picked up %s!", itemName.Value)
				}

				// Only one pickup can fill a tile
				break
			}
		}

		level.Redraw = true
		td.progressTurnState()
		td.resetCounter()
	}

	if td.TurnState == UIOpen {
		// Do some input detection for inventory
	}

	if td.TurnState == PlayerTurn {
		if turnTaken := action.TakePlayerAction(ecs); turnTaken {
			td.progressTurnState()
		}
	}

	if td.TurnState == MonsterTurn {
		action.TakeMonsterAction(ecs)
		td.progressTurnState()
	}

}

func (td *TurnData) progressTurnState() {
	switch td.TurnState {
	case BeforePlayerAction:
		td.TurnState = PlayerTurn
	case PlayerTurn:
		td.TurnState = MonsterTurn
	case MonsterTurn:
		td.TurnState = BeforePlayerAction
	case GameOver:
		td.TurnState = GameOver
	default:
		td.TurnState = PlayerTurn
	}
}

func (td *TurnData) resetCounter() {
	td.TurnCounter = 0
}

func (td *TurnData) gameOver() {
	td.TurnState = GameOver
}
