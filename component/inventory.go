package component

import (
	"errors"

	"github.com/yohamta/donburi"
)

type InventoryData struct {
	Items    []*donburi.Entry
	capacity int
	holding  int
}

var Inventory = donburi.NewComponentType[InventoryData]()

func (i *InventoryData) New(capacity int) InventoryData {
	return InventoryData{
		Items:    make([]*donburi.Entry, capacity),
		capacity: capacity,
		holding:  0,
	}
}

func (i *InventoryData) GetCapacityInfo() (holding, capacity int) {
	return i.holding, i.capacity
}

func (i *InventoryData) IncreaseCapacityByAmount(amount int) {
	i.capacity += amount
	newStorage := make([]*donburi.Entry, i.capacity)
	copy(newStorage, i.Items)
	i.Items = newStorage
}

func (i *InventoryData) DecreaseCapacityByAmount(amount int) error {
	if i.capacity-i.holding > amount {
		return errors.New("holding too many items to reduce capacity")
	}

	return nil
}

func (i *InventoryData) AddItem(item *donburi.Entry) error {
	if i.holding >= i.capacity {
		return errors.New("inventory full")
	}

	var targetIndex = -1

	for index, element := range i.Items {
		if element == nil {
			targetIndex = index
			break
		}
	}

	if targetIndex == -1 {
		return errors.New("failed to find empty index for item")
	}

	i.Items[targetIndex] = item
	return nil
}

func (i *InventoryData) RemoveItem(index int) error {
	if index >= i.capacity {
		return errors.New("index out of range in RemoveItem")
	}
	i.Items[index] = nil
	return nil
}
