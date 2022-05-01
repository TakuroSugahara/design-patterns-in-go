package observer

import "fmt"

// Concrete subject
type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailavility() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notify()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notify() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

func removeFromSlice(observerList []observer, o observer) []observer {
	listLength := len(observerList)
	for i, o := range observerList {
		if o.getID() == o.getID() {
			observerList[i] = observerList[listLength-1]
			return observerList[:listLength-1]
		}
	}
	return observerList
}
