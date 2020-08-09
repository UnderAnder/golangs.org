package main

import (
	"fmt"
	"image"
	"time"
	"log"
)

type command int
 
const (
    right = command(0)
    left  = command(1)
)
// RoverDriver ведет марсоход по поверхности Марса.
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
            case left:  // поворот налево
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

/* func worker() {
    pos := image.Point{X: 10, Y: 10} // Текущая позиция (изначально [10, 10])
    direction := image.Point{X: 1, Y: 0} // Текущее направление (изначально [1, 0])
    next := time.After(time.Second)
    for {
        select {
        case <-next:
            pos = pos.Add(direction)
            fmt.Println("текущая позиция ", pos) // Выводит текущую позицию
            next = time.After(time.Second)
        }
    } 
}*/

func main() {
    r := NewRoverDriver()
    time.Sleep(3 * time.Second)
    r.Left()
    time.Sleep(3 * time.Second)
    r.Right()
    time.Sleep(3 * time.Second)
}