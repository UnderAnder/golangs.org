package main

import (
	"image"
	"log"
	"math/rand"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(3)
	stop  = command(4)
)

// MarsGrid представляет сетку какой-то поверхности
// Марса. Может использовать конкурентно другой
// горутиной.
type MarsGrid struct {
	bounds image.Rectangle
	cells  [][]cell
}

type cell struct {
	occupier *Occupier
}

// Occupier представляет занятую клетку сетки
// Может использоваться конкурентно другой сеткой
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

func (o *Occupier) Pos() image.Point {
	return o.pos
}

// Occupy занимает ячейку в данной точке сетки. Он
// возвращает nil, если точка уже занята или точка
// за пределами сетки. В противном случае возвращается значение, что
// можно переместить в другое место сетки.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

// возвращает клетку по ее координатам
func (g *MarsGrid) cell(p image.Point) *cell {
	return &g.cells[p.Y][p.X]
}

func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
	}
	return grid
}

// Move перемещает occupier на другую клетку сетки
// Сообщает, было ли перемещение успешным
// Может не получиться, если пытается выйти за пределы
// сетки или потому что пытается перместиться в клетку, что
// уже занята. Если проваливается, occupier остается на прежнем месте
//func (g *Occupier) Move(p image.Point) bool

type RoverDriver struct {
	commandc chan command
	occupier *Occupier
}

func NewRoverDriver(occupier *Occupier) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
	}
	go r.drive()
	return r
}

// drive ответственен за вождение марсохода. Ожидается
// что он начнется в горутине.
func (r *RoverDriver) drive() {
	pos := r.occupier.Pos()
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
			case start: // начало движения
				direction = image.Point{X: 1, Y: 0}
				log.Printf("начало движения %v", direction)
			case stop:
				direction = image.Point{X: 0, Y: 0}
				log.Printf("остановка %v", direction)
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
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func main() {
	size := image.Point{X: 20, Y: 10}
	grid := NewMarsGrid(size)

	var o *Occupier
	// Попытка получить случайную точку продолжается до тех пор, пока не будет найдена та,
	// что сейчас не занята
	for o == nil {
		startPoint := image.Point{
			X: rand.Intn(grid.Size().X),
			Y: rand.Intn(grid.Size().Y),
		}
		o = grid.Occupy(startPoint)
	}
	r := NewRoverDriver(o)
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)

}
