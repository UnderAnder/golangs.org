package main

import (
	"fmt"
	"image"
	"time"
)

func worker() {
	pos := image.Point{X: 10, Y: 10}     // Текущая позиция (изначально [10, 10])
	direction := image.Point{X: 1, Y: 0} // Текущее направление (изначально [1, 0])
	delay := time.Second
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("текущая позиция ", pos) // Выводит текущую позицию
			delay += time.Second / 2
			next = time.After(delay)
		}
	}
}

func main() {
	go worker()
	time.Sleep(5 * time.Second)
}
