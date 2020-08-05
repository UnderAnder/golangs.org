package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func (u Universe) Show() {
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
	if y > height {
		y = y % height
	}
	if x > width {
		x = x % width
	}
	u[y][x] = b
}

func (u Universe) Alive(x, y int) bool {
	if u[y][x] == true {
		return true
	} else {
		return false
	}
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
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			a.Neighbors(i,j)
			b.Next(i, j)
		}
	}

	a, b = b, a

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			a.Neighbors(i,j)
			b.Next(i, j)
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
	world := NewUniverse()
	world.Seed()
	world.Show()

	fmt.Println(world.Neighbors(5, 5))
	fmt.Println(world.Next(5, 5))
}
