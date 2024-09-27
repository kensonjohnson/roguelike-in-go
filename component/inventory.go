package component

import (
	"errors"
	"iter"
	"log"

	"github.com/kensonjohnson/roguelike-game-go/archetype/tags"
	"github.com/yohamta/donburi"
)

type InventoryData struct {
	items    []*donburi.Entry
	capacity int
	holding  int
}

var Inventory = donburi.NewComponentType[InventoryData]()

func NewInventory(capacity int) InventoryData {
	return InventoryData{
		items:    make([]*donburi.Entry, capacity),
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
	copy(newStorage, i.items)
	i.items = newStorage
}

func (i *InventoryData) DecreaseCapacityByAmount(amount int) error {
	if i.capacity-i.holding > amount {
		return errors.New("holding too many items to reduce capacity")
	}

	return nil
}

func (i *InventoryData) GetItem(index int) (*donburi.Entry, error) {
	if index < 0 || index >= i.capacity {
		return nil, errors.New("Index out of range")
	}
	return i.items[index], nil
}

func (i *InventoryData) AddItem(item *donburi.Entry) error {
	if i.holding >= i.capacity {
		return errors.New("inventory full")
	}

	if !item.HasComponent(tags.ItemTag) {
		log.Panic("Entry is not an item: ", item)
	}

	var targetIndex = -1

	for index, element := range i.items {
		if element == nil {
			targetIndex = index
			break
		}
	}

	if targetIndex == -1 {
		return errors.New("failed to find empty index for item")
	}

	i.items[targetIndex] = item
	return nil
}

func (i *InventoryData) RemoveItem(index int) error {
	if index < 0 || index >= i.capacity {
		log.Panic("index out of range in RemoveItem. Recieved: ", index)
	}
	i.items[index] = nil
	return nil
}

func (i *InventoryData) Iter() iter.Seq2[int, *donburi.Entry] {
	return func(yield func(int, *donburi.Entry) bool) {
		for index, entry := range i.items {
			if !yield(index, entry) {
				return
			}
		}
	}
}
