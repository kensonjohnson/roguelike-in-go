package component

import "github.com/yohamta/donburi"

type UserMessageData struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}

var UserMessage = donburi.NewComponentType[UserMessageData]()
