package memento

import "fmt"

type Client struct{}

func (c *Client) GetName() string {
	return "memento"
}

func (c *Client) Execute() {
	fmt.Printf("run memento pattern")

	taker := &caretaker{mementoList: make([]*memento, 0)}
	o := &originator{
		state: "A",
	}

	fmt.Printf("Originator current state: %s\n", o.getState())
	taker.addMemento(o.createMement())

	o.setState("B")
	fmt.Printf("Originator current state: %s\n", o.getState())
	taker.addMemento(o.createMement())

	o.setState("C")
	fmt.Printf("Originator current state: %s\n", o.getState())
	taker.addMemento(o.createMement())

	o.restoreMement(taker.getMemento(1))
	fmt.Printf("Restored to state: %s\n", o.getState())

	o.restoreMement(taker.getMemento(0))
	fmt.Printf("Restored to state: %s\n", o.getState())
}
