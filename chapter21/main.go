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
			if u.Alive(x+i, y+k) {
				count++
			}
		}
	}
	return count
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

	fmt.Println(world.Neighbors(3, 3))
}
