package memento

type caretaker struct {
	mementoList []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoList[index]
}
