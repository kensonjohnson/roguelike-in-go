package system

import (
	"github.com/kensonjohnson/roguelike-game-go/archetype"
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
	PlayerTurn
	MonsterTurn
	GameOver
)

func (td *TurnData) Update(ecs *ecs.ECS) {
	td.TurnCounter++

	if td.TurnState == BeforePlayerAction {
		// Check if player survived the last cycle of monster turns
		entry := archetype.PlayerTag.MustFirst(ecs.World)
		playerHealth := component.Health.Get(entry)
		if playerHealth.CurrentHealth <= 0 {
			td.gameOver()
			playerMessages := component.UserMessage.Get(entry)
			playerMessages.GameStateMessage = "Game over!"
		}
		td.progressTurnState()
	}

	if td.TurnState == PlayerTurn && td.TurnCounter > 8 {
		turnTaken := action.TakePlayerAction(ecs)
		if turnTaken {
			td.progressTurnState()
			td.resetCounter()
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
