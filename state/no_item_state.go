package state

import "fmt"

// Concrete state
type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n *noItemState) addItem(item int) error {
	n.vendingMachine.incrementItemCount(item)
	n.vendingMachine.setState(hasItem)
	return nil
}

func (n *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
