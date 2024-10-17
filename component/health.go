package component

import "github.com/yohamta/donburi"

type HealthData struct {
	MaxHealth     int
	CurrentHealth int
}

var Health = donburi.NewComponentType[HealthData]()

func (h HealthData) Add(amount int) int {
	current := h.CurrentHealth
	h.CurrentHealth += amount
	if h.CurrentHealth > h.MaxHealth {
		h.CurrentHealth = h.MaxHealth
	}

	return h.CurrentHealth - current
}

func (h HealthData) Subtract(amount int) int {
	current := h.CurrentHealth
	h.CurrentHealth -= amount
	if h.CurrentHealth < 0 {
		h.CurrentHealth = 0
	}
	return current - h.CurrentHealth
}
