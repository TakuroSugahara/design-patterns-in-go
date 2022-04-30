package main

import (
	"fmt"

	"github.com/TakuroSugahara/design-patterns-in-go/composite"
	"github.com/TakuroSugahara/design-patterns-in-go/observer"
)

type client interface {
	GetName() string
	Execute()
}

type handler struct {
	clientList []client
}

func (h *handler) ExecuteAll() {
	for _, c := range h.clientList {
		fmt.Printf("\n\n\n-------%s----------\n", c.GetName())
		c.Execute()
		fmt.Printf("------END-----------")
	}
}

func main() {
	handler := handler{
		clientList: []client{
			&composite.Client{},
			&observer.Client{},
		},
	}
	handler.ExecuteAll()
}
