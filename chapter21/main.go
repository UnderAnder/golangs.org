package main

import "fmt"

type Universe [][]bool

func (r Universe) show() {
	for _, k := range r {
		for _, v := range k {
			fmt.Printf("%v", v)
		}
		fmt.Println()
	}
}

func NewUniverse() Universe {
	const (
		width  = 80
		height = 15
	)
	universe := make([][]bool, height)
	for i := 0; i < height; i++ {
		universe[i] = make([]bool, width)

	}
	return universe
}

func main() {
	NewUniverse().show()

}
