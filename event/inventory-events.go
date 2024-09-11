package event

import "github.com/yohamta/donburi/features/events"

type OpenInventory struct{}

var OpenInventoryEvent = events.NewEventType[OpenInventory]()

type CloseInventory struct{}

var CloseInventoryEvent = events.NewEventType[CloseInventory]()
