package event

import (
	"github.com/yohamta/donburi/features/events"
)

type ProgressLevel struct{}

var ProgressLevelEvent = events.NewEventType[ProgressLevel]()
