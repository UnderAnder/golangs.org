package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) moveUp() {
	t.y--

}
func (t *turtle) moveDown() {
	t.y++

}
func (t *turtle) moveLeft() {
	t.x--

}
func (t *turtle) moveRight() {
	t.x++

}

func main() {
	t := turtle{x: 0, y: 0}
	t.moveRight()
	t.moveRight()
	t.moveDown()
	fmt.Print(t)
}
