package observer

import (
	"fmt"
)

type Client struct{}

func (c *Client) GetName() string {
	return "observer"
}

// itemを検知するobserverを登録し
// 登録後にpublishする
func (c *Client) Execute() {
	fmt.Printf("Execute observer pattern.\n")

	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	// regsiter observer
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	// publish
	shirtItem.notify()

	fmt.Printf("#############\n")

	shirtItem.deregister(observerFirst)

	shirtItem.notify()
	fmt.Printf("#############\n")

	fmt.Printf("Finished observer pattern.\n")
}
