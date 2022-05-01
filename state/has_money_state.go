package state

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progrss")
}

func (h *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progrss")
}

func (h *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (h *hasMoneyState) dispenseItem() error {
	fmt.Printf("Dispensing Item\n")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(noItem)
	} else {
		h.vendingMachine.setState(hasItem)
	}
	return nil
}
