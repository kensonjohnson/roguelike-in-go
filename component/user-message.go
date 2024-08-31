package component

import "github.com/yohamta/donburi"

type UserMessageData struct {
	AttackMessage           string
	DeadMessage             string
	GameStateMessage        string
	WorldInteractionMessage string
}

var UserMessage = donburi.NewComponentType[UserMessageData]()
