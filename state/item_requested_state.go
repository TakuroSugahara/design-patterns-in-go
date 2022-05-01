package state

import "fmt"

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("Item already reqested\n")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress\n")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please isnert %d", i.vendingMachine.itemPrice)
	}
	fmt.Printf("Money enterd is ok\n")
	i.vendingMachine.setState(hasMoney)
	return nil
}

func (n *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
