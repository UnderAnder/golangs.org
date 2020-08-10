package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

// drive ответственен за вождение марсохода. Ожидается
// что он начнется в горутине.
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc: // Ждет новых команд на командном канале
			switch c {
			case right: // поворот направо
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left: // поворот налево
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left поворачивает марсоход налево (90° против часовой стрелки).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right поворачивает марсоход направо (90° по часовой стрелке).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {

}

func (r *RoverDriver) Stop() {

}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
