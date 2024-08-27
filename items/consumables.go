package items

type consumable struct {
	Name       string
	AmountHeal int
}

type consumables struct {
	HealthPotion consumable
}

var Consumables consumables

func init() {
	Consumables = consumables{
		HealthPotion: consumable{
			Name:       "Health Potion",
			AmountHeal: 10,
		},
	}
}
