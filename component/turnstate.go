package component

import "github.com/yohamta/donburi"

type turnState int

type TurnData struct {
	TurnState   turnState
	TurnCounter int
}

const (
	BeforePlayerAction = iota
	PlayerTurn
	MonsterTurn
	GameOver
)

func (td *TurnData) GetNextState(state turnState) turnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return BeforePlayerAction
	case GameOver:
		return GameOver
	default:
		return PlayerTurn
	}
}

var Turn = donburi.NewComponentType[TurnData]()
