package component

import "github.com/yohamta/donburi"

type WalletData struct {
	Amount int
}

var Wallet = donburi.NewComponentType[WalletData]()

func (w *WalletData) AddAmount(amount int) {
	w.Amount += amount
	if w.Amount > 99999 {
		w.Amount = 99999
	}
}

func (w *WalletData) SubtractAmount(amount int) {
	w.Amount -= amount
	if w.Amount < 0 {
		w.Amount = 0
	}
}
