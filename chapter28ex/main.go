package main

import "fmt"

type character struct {
	name     string
	lefthand item
}

type item struct {
	name string
}

func (c *character) pickup(i *item) {
	c.lefthand = *i
	fmt.Printf("%v взял в руку %v\n", c.name, i.name)
}

func (i *item) give(to *character) {
	to.pickup(i)
	fmt.Printf("")
}

func (c *character) give(to *character) {
	// if c.lefthand == nil {
	// 	fmt.Println("Нечего отдавать")
	// }
	i := c.lefthand
	i.give(to)
	fmt.Printf("%v отдал %v %v\n", c.name, i.name, to.name)
}

func main() {
	c := character{}
	arthur := character{name: "Артур"}
	knight := character{name: "Рыцарь"}
	axe := item{name: "Топор"}
	hammer := item{name: "Молот"}
	sword := item{name: "Меч"}

	fmt.Print(c)

	hammer.give(&c)
	fmt.Print(c)

	c.pickup(&axe)
	fmt.Print(c)

	arthur.pickup(&sword)
	arthur.give(&knight)

}
