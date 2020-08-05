package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func (u Universe) Show() {
	fmt.Print("\x0c")
	for _, k := range u {
		for _, v := range k {
			if v {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

func (u Universe) Alive(x, y int) bool {
	x += width
	x %= width
	y += height
	y %= height

	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for k := -1; k < 2; k++ {
			if !(i == 0 && k == 0) && u.Alive(x+i, y+k) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	switch u.Neighbors(x, y) {
	case 0, 1:
		return false
	case 2:
		return true
	case 3:
		if u.Alive(x, y) {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func NewUniverse() Universe {

	universe := make([][]bool, height)
	for i := 0; i < height; i++ {
		universe[i] = make([]bool, width)

	}
	return universe
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a // Swap universes
	}
}
