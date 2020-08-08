package main

import "fmt"

type character struct {
	name     string
	leftHand *item
}

type item struct {
	name string
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	c.leftHand = i
	fmt.Printf("%v взял в руку %v\n", c.name, i.name)
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v ничего не может дать\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v с занятыми руками\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil

	fmt.Printf("%v отдал %v %v\n", c.name, to.leftHand.name, to.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v ничего не несет", c.name)
	}
	return fmt.Sprintf("%v несет %v", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{name: "Артур"}
	knight := &character{name: "Рыцарь"}
	sword := &item{name: "Меч"}

	fmt.Println(arthur) // Выводит: Артур ничего не несет
	fmt.Println(knight)
	arthur.pickup(sword)
	arthur.give(knight)

	fmt.Println(arthur)
	fmt.Println(knight)
}
