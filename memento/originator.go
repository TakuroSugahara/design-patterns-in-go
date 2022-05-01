package memento

type originator struct {
	state string
}

func (o *originator) createMement() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMement(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}
