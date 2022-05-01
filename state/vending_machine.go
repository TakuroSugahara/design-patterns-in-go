package state

import "fmt"

type vendingMachine struct {
	currentState state

	itemCount int
	itemPrice int
}

type stateType string

const (
	hasItem     stateType = "hasItem"
	requestItem stateType = "requestItem"
	hasMoney    stateType = "hasMoney"
	noItem      stateType = "noItem"
)

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	v.setState(hasItem)
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s stateType) {
	switch s {
	case hasItem:
		v.currentState = &hasItemState{vendingMachine: v}
	case noItem:
		v.currentState = &noItemState{vendingMachine: v}
	case requestItem:
		v.currentState = &itemRequestedState{vendingMachine: v}
	case hasMoney:
		v.currentState = &hasMoneyState{vendingMachine: v}
	}
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
