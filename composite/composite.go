package composite

import "fmt"

// Composite
type folder struct {
	components []Component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, c := range f.components {
		c.search(keyword)
	}
}

func (f *folder) add(c Component) {
	f.components = append(f.components, c)
}
